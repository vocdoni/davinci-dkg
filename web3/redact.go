package web3

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Providers put API keys in an endpoint's path (…/v2/<key>), its query
// (?apikey=) or its userinfo (user:pass@), and the node's logs are shipped
// to third parties. Only the scheme and host of an endpoint are ever logged,
// and http(s) endpoints are dialed so that the full URL never reaches an
// error message either (see dialEndpoint).

// unparseableURL stands in for an endpoint URL that does not parse.
const unparseableURL = "<unparseable rpc url>"

// The schemes dialEndpoint dials through endpointTransport.
const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
)

// RedactURL returns the part of an endpoint URL that may be logged:
// scheme://host[:port], without userinfo, path, query or fragment.
func RedactURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return unparseableURL
	}
	return u.Scheme + "://" + u.Host
}

// redactor holds the secret parts of the configured endpoint URLs, longest
// first, and replaces them in error text. It is the belt over the
// transport's rewrite of error bodies: whatever spelling or parser
// difference lets a provider's echo through, the secrets are known in
// advance.
type redactor [][]byte

// newRedactor collects the secrets of every endpoint URL that parses.
func newRedactor(urls ...string) redactor {
	var secrets [][]byte
	for _, raw := range urls {
		if u, err := url.Parse(raw); err == nil {
			secrets = append(secrets, endpointSecrets(u)...)
		}
	}
	slices.SortFunc(secrets, func(a, b []byte) int { return cmp.Or(cmp.Compare(len(b), len(a)), bytes.Compare(a, b)) })
	return slices.CompactFunc(secrets, bytes.Equal)
}

// error returns err with every secret replaced in its text. The result
// keeps err's behaviour: errors.Is and errors.As reach err, and for a
// JSON-RPC error the code and the revert data stay available
// (redactedRPCError) without the raw text.
func (r redactor) error(err error) error {
	if err == nil || len(r) == 0 {
		return err
	}
	text := err.Error()
	clean := text
	for _, secret := range r {
		clean = strings.ReplaceAll(clean, string(secret), string(redactedMark))
	}
	if clean == text {
		return err
	}
	var rpcErr rpc.Error
	if errors.As(err, &rpcErr) {
		return &redactedRPCError{redactedError{err: err, text: clean}}
	}
	return &redactedError{err: err, text: clean}
}

// redactedError is an error whose text has the endpoints' secrets replaced.
type redactedError struct {
	err  error
	text string
}

func (e *redactedError) Error() string { return e.text }
func (e *redactedError) Unwrap() error { return e.err }

// redactedRPCError is a redacted JSON-RPC error. errors.As for rpc.Error and
// rpc.DataError stops here, so callers get the code and the revert data
// (ethclient.RevertErrorData) and never the raw message.
type redactedRPCError struct{ redactedError }

func (e *redactedRPCError) ErrorCode() int {
	var rpcErr rpc.Error
	if errors.As(e.err, &rpcErr) {
		return rpcErr.ErrorCode()
	}
	return 0
}

func (e *redactedRPCError) ErrorData() any {
	var dataErr rpc.DataError
	if errors.As(e.err, &dataErr) {
		return dataErr.ErrorData()
	}
	return nil
}

// RedactEndpointError is redactor.error for a client from DialRPC: err with
// the secret parts of endpoint (userinfo, query values, path) replaced.
func RedactEndpointError(endpoint string, err error) error {
	return newRedactor(endpoint).error(err)
}

// redactError rewrites the URL net/url and net/http embed in their errors
// (*url.Error, e.g. `parse "<url>": …` or `Post "<url>": …`) to RedactURL's
// form, in place, and returns err.
func redactError(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		urlErr.URL = RedactURL(urlErr.URL)
	}
	return err
}

// dialEndpoint dials one pool endpoint. An http(s) endpoint is dialed at
// scheme://host only, and its transport puts the path, query and
// credentials back on every request: go-ethereum and net/http only ever see
// the redacted URL, so none of the errors they build (`Post "<url>": read
// tcp …: connection reset by peer`) can carry the key into a log line, on
// any path that uses the client, the txmanager's included. Redirects are
// never followed, same origin or not, so the credentials cannot go to
// another URL: the 3xx comes back as an HTTP error, which the pool treats
// as the endpoint's fault (isRedirect).
// Websocket and ipc endpoints are dialed as they are.
func dialEndpoint(raw string) (*ethclient.Client, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, redactError(err)
	}
	endpoint := raw
	// The timeout bounds every round trip, so a hung endpoint fails over
	// instead of blocking a tick forever. It only applies to http(s).
	httpClient := &http.Client{Timeout: rpcHTTPTimeout}
	if u.Scheme == schemeHTTP || u.Scheme == schemeHTTPS {
		endpoint = RedactURL(raw)
		httpClient.Transport = &endpointTransport{base: http.DefaultTransport, target: u, secrets: endpointSecrets(u)}
		httpClient.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	}
	rpcClient, err := rpc.DialOptions(context.Background(), endpoint, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, redactError(err)
	}
	return ethclient.NewClient(rpcClient), nil
}

// DialRPC dials one endpoint the way the pool does (see dialEndpoint), for
// callers that need a plain *rpc.Client: none of the errors it returns
// carries the endpoint's path, query or credentials.
func DialRPC(raw string) (*rpc.Client, error) {
	client, err := dialEndpoint(raw)
	if err != nil {
		return nil, err
	}
	return client.Client(), nil
}

// endpointTransport sends each request, addressed to the endpoint's
// scheme://host, to its full URL with its credentials, and keeps them out
// of what comes back: net/http builds its errors from the response's
// request, so that is the redacted one; a redirect's Location is dropped
// (redirects are never followed, and without one net/http returns the 3xx
// as is rather than parse and echo it); and the endpoint's secrets are
// replaced where go-ethereum takes error text from, a non-2xx body or a
// JSON-RPC error's message (see scrub). Status codes, JSON-RPC ids and
// codes, revert data and results keep their exact bytes.
type endpointTransport struct {
	base    http.RoundTripper
	target  *url.URL // the endpoint's full URL
	secrets [][]byte // see endpointSecrets
}

func (t *endpointTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	out := req.Clone(req.Context())
	target := *t.target
	target.User = nil
	target.Fragment, target.RawFragment = "", ""
	out.URL = &target
	if user := t.target.User; user != nil && out.Header.Get("Authorization") == "" {
		pass, _ := user.Password()
		out.SetBasicAuth(user.Username(), pass)
	}
	resp, err := t.base.RoundTrip(out)
	if err != nil {
		return nil, err
	}
	resp.Request = req
	resp.Header.Del("Location")
	if err := t.scrub(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

const (
	// maxScannedBody bounds the non-2xx body searched for secrets: a
	// larger one is replaced by its size (it is never parsed, only
	// quoted). A 2xx body is scanned whole, whatever its size.
	maxScannedBody = 1 << 20
	// maxCountedBody bounds how much of an oversized error body is read to
	// report its size.
	maxCountedBody = 64 << 20
	// minSegmentLen is the shortest path segment treated as a possible key:
	// "v2" or "rpc" are none. Userinfo and query values are secrets at any
	// length.
	minSegmentLen = 8
)

// redactedMark replaces a secret in a response body.
var redactedMark = []byte("<redacted>")

// endpointSecrets lists the parts of an endpoint URL that may carry a key,
// longest first so a longer match is replaced whole: the userinfo, the user
// name and password, the query and each of its values, at any length; the
// path (unless just "/"), and each of its segments at least minSegmentLen
// long. They are only ever replaced in error text (see scrub).
func endpointSecrets(u *url.URL) [][]byte {
	var parts []string
	if u.User != nil {
		parts = append(parts, u.User.String(), u.User.Username())
		if pass, ok := u.User.Password(); ok {
			parts = append(parts, pass)
		}
	}
	parts = append(parts, u.RawQuery)
	for _, values := range u.Query() {
		parts = append(parts, values...)
	}
	if u.Path != "/" {
		parts = append(parts, u.Path, u.EscapedPath())
	}
	for _, segment := range strings.Split(u.Path, "/") {
		if len(segment) >= minSegmentLen {
			parts = append(parts, segment)
		}
	}
	parts = slices.DeleteFunc(parts, func(p string) bool { return p == "" })
	slices.SortFunc(parts, func(a, b string) int { return cmp.Or(cmp.Compare(len(b), len(a)), strings.Compare(a, b)) })
	parts = slices.Compact(parts)
	secrets := make([][]byte, len(parts))
	for i, p := range parts {
		secrets[i] = []byte(p)
	}
	return secrets
}

// scrub replaces the endpoint's secrets where go-ethereum takes error text
// from: the whole body of a non-2xx answer, and the message of every
// JSON-RPC error in a 2xx one, of any size. A non-2xx body too large to
// scan fails closed: it is replaced by its size.
func (t *endpointTransport) scrub(resp *http.Response) error {
	if len(t.secrets) == 0 || resp.Body == nil {
		return nil
	}
	limit := int64(maxScannedBody + 1)
	if resp.StatusCode/100 == 2 {
		// A JSON-RPC error comes over HTTP 200 and can be any size: the
		// whole body is scanned, as go-ethereum reads it whole anyway.
		limit = math.MaxInt64
	}
	head, err := io.ReadAll(io.LimitReader(resp.Body, limit))
	if err != nil {
		_ = resp.Body.Close()
		return err
	}
	if resp.StatusCode/100 != 2 && len(head) > maxScannedBody {
		rest, _ := io.Copy(io.Discard, io.LimitReader(resp.Body, maxCountedBody))
		_ = resp.Body.Close()
		head = fmt.Appendf(nil, "<provider error body redacted: %d bytes>", int64(len(head))+rest)
		resp.Body = io.NopCloser(bytes.NewReader(head))
		resp.ContentLength = int64(len(head))
		return nil
	}
	_ = resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		// go-ethereum never parses a non-2xx body: it only quotes it as
		// the error text (rpc.HTTPError), so all of it is text.
		head = t.redactText(head)
	} else {
		head = t.redactMessages(head)
	}
	resp.Body = io.NopCloser(bytes.NewReader(head))
	resp.ContentLength = int64(len(head))
	return nil
}

// redactText replaces every one of the endpoint's secrets in text.
func (t *endpointTransport) redactText(text []byte) []byte {
	for _, secret := range t.secrets {
		text = bytes.ReplaceAll(text, secret, redactedMark)
	}
	return text
}

// redactMessages redacts the "message" string of every JSON-RPC error in a
// response or a batch of them, re-encoded in place: ids, codes, error data
// and results keep their exact bytes. A body that is not a JSON-RPC
// response is plain text and is redacted whole.
func (t *endpointTransport) redactMessages(body []byte) []byte {
	spans, ok := errorMessageSpans(body)
	if !ok {
		return t.redactText(body)
	}
	var out []byte
	last := int64(0)
	for _, sp := range spans {
		var msg string
		if err := json.Unmarshal(body[sp.start:sp.end], &msg); err != nil {
			continue
		}
		redacted := string(t.redactText([]byte(msg)))
		if redacted == msg {
			continue
		}
		literal, err := jsonString(redacted)
		if err != nil {
			continue
		}
		out = append(append(out, body[last:sp.start]...), literal...)
		last = sp.end
	}
	if out == nil {
		return body
	}
	return append(out, body[last:]...)
}

// span is a byte range [start, end) of a body.
type span struct{ start, end int64 }

// errorMessageSpans locates the string literal of every error's "message"
// in a JSON-RPC response object or batch array; ok is false for anything
// else.
func errorMessageSpans(body []byte) (spans []span, ok bool) {
	dec := json.NewDecoder(bytes.NewReader(body))
	tok, err := dec.Token()
	if err != nil {
		return nil, false
	}
	switch tok {
	case json.Delim('{'):
		if !responseMessageSpans(dec, &spans) {
			return nil, false
		}
	case json.Delim('['):
		for dec.More() {
			if tok, err := dec.Token(); err != nil || tok != json.Delim('{') {
				return nil, false
			}
			if !responseMessageSpans(dec, &spans) {
				return nil, false
			}
		}
		if _, err := dec.Token(); err != nil {
			return nil, false
		}
	default:
		return nil, false
	}
	return spans, true
}

// responseMessageSpans reads one response's members (its '{' already read)
// through its '}' and collects the spans of its errors' messages: every
// "message" of every "error" member, duplicates included.
func responseMessageSpans(dec *json.Decoder, spans *[]span) bool {
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return false
		}
		var value json.RawMessage
		if err := dec.Decode(&value); err != nil {
			return false
		}
		// go-ethereum matches keys case-insensitively (the decoder has
		// already unescaped them).
		if k, _ := key.(string); !strings.EqualFold(k, "error") {
			continue
		}
		start := dec.InputOffset() - int64(len(value))
		for _, msg := range messageSpans(value) {
			*spans = append(*spans, span{start + msg.start, start + msg.end})
		}
	}
	_, err := dec.Token()
	return err == nil
}

// messageSpans locates every "message" string literal of one error object,
// whatever its case: Go's decoder matches keys case-insensitively and keeps
// the last of duplicates, so all of them are text.
func messageSpans(obj []byte) []span {
	dec := json.NewDecoder(bytes.NewReader(obj))
	if tok, err := dec.Token(); err != nil || tok != json.Delim('{') {
		return nil
	}
	var spans []span
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return spans
		}
		var value json.RawMessage
		if err := dec.Decode(&value); err != nil {
			return spans
		}
		if k, _ := key.(string); strings.EqualFold(k, "message") && len(value) > 0 && value[0] == '"' {
			end := dec.InputOffset()
			spans = append(spans, span{end - int64(len(value)), end})
		}
	}
	return spans
}

// jsonString encodes s as a JSON string literal, without HTML escaping.
func jsonString(s string) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(buf.Bytes(), []byte("\n")), nil
}
