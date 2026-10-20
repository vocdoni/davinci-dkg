package web3

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/log"
)

// lockedBuffer is a bytes.Buffer safe for the logger and the test to share.
type lockedBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *lockedBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *lockedBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// captureLogs sends every log line, at debug level and up, to a file (the
// console format operators see) and to a buffer (JSON), and returns both
// read back.
func captureLogs(t *testing.T) func() string {
	t.Helper()
	console := filepath.Join(t.TempDir(), "node.log")
	var structured lockedBuffer
	log.Init(log.LogLevelDebug, console, &structured)
	t.Cleanup(func() { log.Init(cmp.Or(os.Getenv("LOG_LEVEL"), log.LogLevelError), "stderr", nil) })
	return func() string {
		raw, err := os.ReadFile(console)
		qt.Assert(t, err, qt.IsNil)
		return string(raw) + structured.String()
	}
}

// Operators put provider API keys in the endpoint's path (…/v2/<key>), its
// query (?apikey=) or its userinfo (user:pass@), and the node's logs end up
// on Railway. An endpoint that fails to dial, refuses (403), rate-limits
// (429), is unreachable, rotates away and comes back must never print any of
// those parts: not in the pool's own lines and not in the errors callers log.
// The endpoint that works still gets its full URL and credentials.
func TestPoolNeverLogsEndpointSecrets(t *testing.T) {
	c := qt.New(t)
	logs := captureLogs(t)

	refusing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(archiveRefusal))
	}))
	defer refusing.Close()
	throttling := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"error":{"code":-32005,"message":"rate limit exceeded"}}`))
	}))
	defer throttling.Close()
	// Nothing listens on port 1; a closed httptest server's port could be
	// handed to the servers below and stop being unreachable.
	const down = "http://127.0.0.1:1"
	log0, _ := fakeLogAndReceipt()
	healthy := &fakeEndpoint{head: 100, log: log0, calls: map[string]int{}}
	// The healthy endpoint only answers with its own path, query and credentials.
	authed := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if r.URL.Path != "/v2/GOODKEY" || r.URL.Query().Get("apikey") != "GOODQUERY" ||
			!ok || user != "gooduser" || pass != "GOODPASS" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		healthy.serve(w, r)
	}))
	defer authed.Close()

	withSecrets := func(base, userinfo, path string) string {
		return strings.Replace(base, "http://", "http://"+userinfo, 1) + path
	}
	secrets := []string{
		"DIALUSER", "DIALPASS", "DIALKEY",
		"USER403", "PASS403", "KEY403", "QUERY403",
		"KEY429", "QUERY429",
		"USERDOWN", "PASSDOWN", "KEYDOWN", "QUERYDOWN",
		"gooduser", "GOODPASS", "GOODKEY", "GOODQUERY",
	}
	urls := []string{
		"https://DIALUSER:DIALPASS@bad.example/%zz/DIALKEY",
		withSecrets(refusing.URL, "USER403:PASS403@", "/v3/KEY403?apikey=QUERY403"),
		withSecrets(throttling.URL, "", "/rpc/KEY429?token=QUERY429"),
		withSecrets(down, "USERDOWN:PASSDOWN@", "/v2/KEYDOWN?apikey=QUERYDOWN"),
		withSecrets(authed.URL, "gooduser:GOODPASS@", "/v2/GOODKEY?apikey=GOODQUERY"),
	}
	var errs []string
	ctx := context.Background()
	query := ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)}

	pool, err := NewRPCPool(urls) // the first endpoint does not dial
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := NewPooledBackend(pool)
	// 403, 429, then the unreachable endpoint until it is disabled; the last
	// endpoint answers every call.
	for range rpcMaxFailures {
		got, err := backend.FilterLogs(ctx, query)
		c.Assert(err, qt.IsNil, qt.Commentf("the authenticated endpoint must answer"))
		c.Assert(got, qt.HasLen, 1)
	}
	// The cooled-down endpoints come back and refuse again.
	for _, e := range pool.entries {
		if !e.disabledAt.IsZero() {
			e.disabledAt = time.Now().Add(-rpcCooldownDuration)
		}
	}
	_, err = backend.FilterLogs(ctx, query)
	c.Assert(err, qt.IsNil)

	// Errors handed back to callers, which log them: from the pooled backend
	// and from a raw client (the txmanager's), for every failing endpoint.
	for _, raw := range urls[1:4] {
		alone, err := NewRPCPool([]string{raw})
		c.Assert(err, qt.IsNil)
		_, err = NewPooledBackend(alone).FilterLogs(ctx, query)
		c.Assert(err, qt.IsNotNil)
		errs = append(errs, err.Error())
		_, err = alone.Current().BlockNumber(ctx)
		c.Assert(err, qt.IsNotNil)
		errs = append(errs, err.Error())
		alone.Close()
	}
	_, err = NewRPCPool(urls[:1])
	c.Assert(err, qt.IsNotNil)
	errs = append(errs, err.Error())

	out := logs()
	for _, want := range []string{"failed to dial endpoint", "refused the request", "rate-limited", "unreachable", "re-enabled endpoint"} {
		c.Assert(out, qt.Contains, want)
	}
	for _, secret := range secrets {
		c.Assert(strings.Contains(out, secret), qt.IsFalse, qt.Commentf("%s leaked into the logs:\n%s", secret, out))
		for _, e := range errs {
			c.Assert(strings.Contains(e, secret), qt.IsFalse, qt.Commentf("%s leaked into a returned error: %s", secret, e))
		}
	}
}

// RedactURL keeps the scheme, host and port of an endpoint and nothing a
// provider hides a key in.
func TestRedactURLKeepsOnlySchemeAndHost(t *testing.T) {
	for _, tc := range []struct{ raw, want string }{
		{"https://user:secret@eth-sepolia.example.com/v3/apikey123", "https://eth-sepolia.example.com"},
		{"https://opt-mainnet.g.example.com/v2/KEY?apikey=Q#frag", "https://opt-mainnet.g.example.com"},
		{"http://127.0.0.1:8545", "http://127.0.0.1:8545"},
		{"ws://[::1]:8546/ws?key=1", "ws://[::1]:8546"},
		{"https://KEY@rpc.example", "https://rpc.example"},
		{"https://rpc.example/%zz/KEY", unparseableURL},
		{"not a url", unparseableURL},
		{"", unparseableURL},
	} {
		qt.Check(t, RedactURL(tc.raw), qt.Equals, tc.want, qt.Commentf("%q", tc.raw))
	}
}

// Endpoint URLs must be canonical: the node never follows a redirect, so the
// credentials in an endpoint's path, query or userinfo never reach another
// URL, same origin or not. A redirecting endpoint is the endpoint's fault:
// it is disabled and the call is served by the next one, and the log says
// so with the endpoint redacted.
func TestRedirectingEndpointIsSkippedAndCredentialsStay(t *testing.T) {
	logs := captureLogs(t)
	var foreignHits atomic.Int32
	foreign := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { foreignHits.Add(1) }))
	defer foreign.Close()
	log0, _ := fakeLogAndReceipt()
	healthyURL := newFakeEndpoint(t, &fakeEndpoint{head: 100, log: log0})
	ctx := context.Background()

	for _, tc := range []struct {
		name     string
		status   int
		location string
	}{
		{"same-origin 307", http.StatusTemporaryRedirect, "/v2/OTHERKEY-0123456789"},
		{"same-origin 308", http.StatusPermanentRedirect, "/v2/OTHERKEY-0123456789"},
		{"cross-origin 302", http.StatusFound, foreign.URL + "/collect"},
		{"cross-origin 307", http.StatusTemporaryRedirect, foreign.URL + "/collect"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := qt.New(t)
			var hits atomic.Int32
			redirecting := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				hits.Add(1)
				w.Header().Set("Location", tc.location)
				w.WriteHeader(tc.status)
			}))
			defer redirecting.Close()
			primary := strings.Replace(redirecting.URL, "http://", "http://REDIRUSER:REDIRPASS@", 1) + "/v2/REDIRKEY?apikey=REDIRQUERY"
			pool, err := NewRPCPool([]string{primary, healthyURL})
			c.Assert(err, qt.IsNil)
			defer pool.Close()

			got, err := NewPooledBackend(pool).FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
			c.Assert(err, qt.IsNil, qt.Commentf("the next endpoint must serve the call"))
			c.Assert(got, qt.HasLen, 1)
			c.Assert(hits.Load(), qt.Equals, int32(1), qt.Commentf("the redirect must not be followed"))
			c.Assert(pool.entries[0].disabledAt.IsZero(), qt.IsFalse, qt.Commentf("a redirecting endpoint is disabled"))
		})
	}
	qt.Assert(t, foreignHits.Load(), qt.Equals, int32(0), qt.Commentf("no request may reach a redirect target"))
	out := logs()
	qt.Assert(t, out, qt.Contains, "redirect")
	for _, secret := range []string{"REDIRUSER", "REDIRPASS", "REDIRKEY", "REDIRQUERY", "OTHERKEY"} {
		qt.Assert(t, strings.Contains(out, secret), qt.IsFalse, qt.Commentf("%s leaked into the logs:\n%s", secret, out))
	}
}

// The answers a test endpoint gives: an HTTP error status, a JSON-RPC error,
// a redirect.
const (
	modeStatus   = "status"
	modeJSONRPC  = "jsonrpc"
	modeRedirect = "redirect"
)

// reflectingEndpoint answers like a provider or gateway whose errors echo
// the request: its URI (path and query) and the basic-auth credentials.
// mode picks the answer: a malformed redirect whose Location carries the
// path, an HTTP 403 whose body quotes the request, or a JSON-RPC error (HTTP
// 200) whose message quotes it, with a revert's code and data.
func reflectingEndpoint(t *testing.T, mode string) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		echo := r.RequestURI + " as " + user + ":" + pass
		switch mode {
		case modeRedirect:
			w.Header().Set("Location", "http://[::1"+r.RequestURI+"/%zz")
			w.WriteHeader(http.StatusTemporaryRedirect)
		case modeStatus:
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":"no access for ` + echo + `"}`))
		case modeJSONRPC:
			var req testRPCRequest
			_ = json.NewDecoder(r.Body).Decode(&req)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"jsonrpc": "2.0", "id": req.ID,
				"error": map[string]any{"code": 3, "message": "execution reverted: bad key " + echo, "data": "0x08c379a0"},
			})
		}
	}))
	t.Cleanup(srv.Close)
	return strings.Replace(srv.URL, "http://", "http://REFLUSER:REFLPASS@", 1) +
		"/v2/REFLKEY-0123456789abcdef?apikey=REFLQUERY-0123456789"
}

// Errors built at run time must not carry the endpoint's secrets either: not
// the URL net/http puts in a redirect error, not a provider body or JSON-RPC
// message that echoes the request, on the pooled backend or a raw client
// (the txmanager's). The structured parts callers use survive: the HTTP
// status, the JSON-RPC code and the revert data.
func TestRuntimeErrorsCarryNoEndpointSecrets(t *testing.T) {
	ctx := context.Background()
	secrets := []string{"REFLUSER", "REFLPASS", "REFLKEY", "REFLQUERY", "0123456789"}
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	for _, mode := range []string{modeRedirect, modeStatus, modeJSONRPC} {
		t.Run(mode, func(t *testing.T) {
			c := qt.New(t)
			pool, err := NewRPCPool([]string{reflectingEndpoint(t, mode)})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			_, pooledErr := NewPooledBackend(pool).CallContract(ctx, ethereum.CallMsg{To: &to}, nil)
			_, rawErr := pool.Current().CallContract(ctx, ethereum.CallMsg{To: &to}, nil)
			for _, err := range []error{pooledErr, rawErr} {
				c.Assert(err, qt.IsNotNil)
				for _, secret := range secrets {
					c.Assert(strings.Contains(err.Error(), secret), qt.IsFalse, qt.Commentf("%s leaked: %v", secret, err))
				}
				switch mode {
				case modeStatus:
					var httpErr rpc.HTTPError
					c.Assert(errors.As(err, &httpErr), qt.IsTrue)
					c.Assert(httpErr.StatusCode, qt.Equals, http.StatusForbidden)
				case modeJSONRPC:
					var rpcErr rpc.Error
					c.Assert(errors.As(err, &rpcErr), qt.IsTrue)
					c.Assert(rpcErr.ErrorCode(), qt.Equals, 3)
					data, ok := ethclient.RevertErrorData(err)
					c.Assert(ok, qt.IsTrue)
					c.Assert(data, qt.DeepEquals, []byte{0x08, 0xc3, 0x79, 0xa0})
				}
			}
		})
	}
}

// jsonRPCEndpoint answers every request of a call or a batch with answer(i)
// for the i-th request, keeping each request's id.
func jsonRPCEndpoint(t *testing.T, userinfo, pathQuery string, answer func(i int) map[string]any) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body json.RawMessage
		_ = json.NewDecoder(r.Body).Decode(&body)
		batch := len(body) > 0 && body[0] == '['
		var reqs []testRPCRequest
		if batch {
			_ = json.Unmarshal(body, &reqs)
		} else {
			var one testRPCRequest
			_ = json.Unmarshal(body, &one)
			reqs = []testRPCRequest{one}
		}
		resps := make([]map[string]any, len(reqs))
		for i, req := range reqs {
			resps[i] = answer(i)
			resps[i]["jsonrpc"], resps[i]["id"] = "2.0", req.ID
		}
		w.Header().Set("Content-Type", "application/json")
		if batch {
			_ = json.NewEncoder(w).Encode(resps)
			return
		}
		_ = json.NewEncoder(w).Encode(resps[0])
	}))
	t.Cleanup(srv.Close)
	return strings.Replace(srv.URL, "http://", "http://"+userinfo, 1) + pathQuery
}

// Redaction touches the text of JSON-RPC error messages only. A credential
// that happens to equal a revert's selector is redacted from the message,
// while the revert data, which the node decodes into a contract error, keeps
// its exact bytes.
func TestRedactionKeepsRevertDataIntact(t *testing.T) {
	c := qt.New(t)
	url := jsonRPCEndpoint(t, "", "/rpc?apikey=08c379a0", func(int) map[string]any {
		return map[string]any{"error": map[string]any{
			"code": 3, "message": "execution reverted: key 08c379a0 refused", "data": "0x08c379a0",
		}}
	})
	pool, err := NewRPCPool([]string{url})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	_, err = NewPooledBackend(pool).CallContract(context.Background(), ethereum.CallMsg{To: &to}, nil)
	c.Assert(err, qt.IsNotNil)
	c.Assert(strings.Contains(err.Error(), "08c379a0"), qt.IsFalse, qt.Commentf("the key leaked: %v", err))
	data, ok := ethclient.RevertErrorData(err)
	c.Assert(ok, qt.IsTrue, qt.Commentf("revert data lost: %v", err))
	c.Assert(data, qt.DeepEquals, []byte{0x08, 0xc3, 0x79, 0xa0})
}

// In a batch that mixes a result and an error, the error's message is
// redacted and the result keeps its exact bytes, even when it contains the
// credential's characters.
func TestRedactionKeepsBatchResultsByteIdentical(t *testing.T) {
	c := qt.New(t)
	const key = "deadbeefcafe0123"
	url := jsonRPCEndpoint(t, "", "/v2/"+key, func(i int) map[string]any {
		if i == 0 {
			return map[string]any{"result": "0x00" + key + "00"}
		}
		return map[string]any{"error": map[string]any{"code": -32000, "message": "bad key " + key}}
	})
	pool, err := NewRPCPool([]string{url})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	var ok, failed hexutil.Bytes
	batch := []rpc.BatchElem{
		{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &ok},
		{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &failed},
	}
	c.Assert(NewPooledBackend(pool).BatchCallContext(context.Background(), batch), qt.IsNil)
	c.Assert(batch[0].Error, qt.IsNil)
	c.Assert(ok.String(), qt.Equals, "0x00"+key+"00")
	c.Assert(batch[1].Error, qt.IsNotNil)
	c.Assert(strings.Contains(batch[1].Error.Error(), key), qt.IsFalse, qt.Commentf("the key leaked: %v", batch[1].Error))
}

// An error body too large to scan is not quoted: go-ethereum would put all
// of it in the error text, credentials at its start or end included, so it
// is replaced whole by its size.
func TestOversizedErrorBodyFailsClosed(t *testing.T) {
	c := qt.New(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("denied " + r.RequestURI + " "))
		_, _ = w.Write(bytes.Repeat([]byte("x"), 2<<20))
		_, _ = w.Write([]byte(" denied " + r.RequestURI))
	}))
	defer srv.Close()
	pool, err := NewRPCPool([]string{srv.URL + "/v2/BIGBODYKEY-0123456789?apikey=BIGBODYQUERY"})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	_, pooledErr := NewPooledBackend(pool).CallContract(context.Background(), ethereum.CallMsg{To: &to}, nil)
	_, rawErr := pool.Current().CallContract(context.Background(), ethereum.CallMsg{To: &to}, nil)
	for _, err := range []error{pooledErr, rawErr} {
		c.Assert(err, qt.IsNotNil)
		c.Assert(err.Error(), qt.Matches, `403 Forbidden: <provider error body redacted: \d+ bytes>`)
		var httpErr rpc.HTTPError
		c.Assert(errors.As(err, &httpErr), qt.IsTrue)
		c.Assert(httpErr.StatusCode, qt.Equals, http.StatusForbidden)
	}
}

// Query values and userinfo are redacted whatever their length: a short
// key ("k9x") echoed on its own, without its parameter name, must not reach
// the error text. Only path segments keep a length floor.
func TestShortCredentialsAreRedacted(t *testing.T) {
	for _, mode := range []string{modeStatus, modeJSONRPC} {
		t.Run(mode, func(t *testing.T) {
			c := qt.New(t)
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				user, pass, _ := r.BasicAuth()
				msg := "invalid key " + r.URL.Query().Get("apikey") + " for " + user + ":" + pass
				if mode == modeStatus {
					w.WriteHeader(http.StatusUnauthorized)
					_, _ = w.Write([]byte(msg))
					return
				}
				var req testRPCRequest
				_ = json.NewDecoder(r.Body).Decode(&req)
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"jsonrpc": "2.0", "id": req.ID, "error": map[string]any{"code": -32000, "message": msg},
				})
			}))
			defer srv.Close()
			url := strings.Replace(srv.URL, "http://", "http://zq:jw7@", 1) + "/rpc?apikey=k9x"
			pool, err := NewRPCPool([]string{url})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			to := common.HexToAddress("0x2000000000000000000000000000000000000002")
			_, err = NewPooledBackend(pool).CallContract(context.Background(), ethereum.CallMsg{To: &to}, nil)
			c.Assert(err, qt.IsNotNil)
			for _, secret := range []string{"k9x", "zq", "jw7"} {
				c.Assert(strings.Contains(err.Error(), secret), qt.IsFalse, qt.Commentf("%s leaked: %v", secret, err))
			}
		})
	}
}

// rawEndpoint answers every request with body(id), raw JSON bytes built
// around the request's id (or ids, for a batch), status 200.
func rawEndpoint(t *testing.T, pathQuery string, body func(ids []json.RawMessage) []byte) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var raw json.RawMessage
		_ = json.NewDecoder(r.Body).Decode(&raw)
		var reqs []testRPCRequest
		if len(raw) > 0 && raw[0] == '[' {
			_ = json.Unmarshal(raw, &reqs)
		} else {
			var one testRPCRequest
			_ = json.Unmarshal(raw, &one)
			reqs = []testRPCRequest{one}
		}
		ids := make([]json.RawMessage, len(reqs))
		for i, req := range reqs {
			ids[i] = req.ID
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body(ids))
	}))
	t.Cleanup(srv.Close)
	return srv.URL + pathQuery
}

// JSON-RPC errors come over HTTP 200, and a large one is still error text:
// its message is redacted whatever the body's size, in a single response
// and in a batch whose large result keeps its exact bytes.
func TestOversizedJSONRPCErrorsAreRedacted(t *testing.T) {
	const key = "LARGEKEY-0123456789"
	padding := strings.Repeat("x", 3<<19) // 1.5 MiB
	ctx := context.Background()
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")

	t.Run("single", func(t *testing.T) {
		c := qt.New(t)
		url := rawEndpoint(t, "/v2/"+key, func(ids []json.RawMessage) []byte {
			return []byte(`{"jsonrpc":"2.0","id":` + string(ids[0]) + `,"error":{"code":-32000,"message":"bad key ` +
				key + " " + padding + " bad key " + key + `"}}`)
		})
		pool, err := NewRPCPool([]string{url})
		c.Assert(err, qt.IsNil)
		defer pool.Close()
		_, err = NewPooledBackend(pool).CallContract(ctx, ethereum.CallMsg{To: &to}, nil)
		c.Assert(err, qt.IsNotNil)
		c.Assert(strings.Contains(err.Error(), key), qt.IsFalse, qt.Commentf("the key leaked from a %d-byte error", len(err.Error())))
	})
	t.Run("batch", func(t *testing.T) {
		c := qt.New(t)
		result := "0x" + strings.Repeat("ab", 3<<19)
		url := rawEndpoint(t, "/v2/"+key, func(ids []json.RawMessage) []byte {
			return []byte(`[{"jsonrpc":"2.0","id":` + string(ids[0]) + `,"error":{"code":-32000,"message":"bad key ` + key + `"}},` +
				`{"jsonrpc":"2.0","id":` + string(ids[1]) + `,"result":"` + result + `"},` +
				`{"jsonrpc":"2.0","id":` + string(ids[2]) + `,"error":{"code":-32000,"message":"bad key ` + key + `"}}]`)
		})
		pool, err := NewRPCPool([]string{url})
		c.Assert(err, qt.IsNil)
		defer pool.Close()
		var first, middle, last hexutil.Bytes
		batch := []rpc.BatchElem{
			{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &first},
			{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &middle},
			{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &last},
		}
		c.Assert(NewPooledBackend(pool).BatchCallContext(ctx, batch), qt.IsNil)
		c.Assert(middle.String(), qt.Equals, result)
		for _, i := range []int{0, 2} {
			c.Assert(batch[i].Error, qt.IsNotNil)
			c.Assert(strings.Contains(batch[i].Error.Error(), key), qt.IsFalse, qt.Commentf("element %d: %v", i, batch[i].Error))
		}
	})
}

// Go's JSON decoder keeps the last of duplicate properties, so every
// "message" of an error object is redacted, not just the first, in a single
// response and in a batch; the bytes around them are left alone.
func TestDuplicateErrorMessagesAreAllRedacted(t *testing.T) {
	const key = "DUPKEY-0123456789"
	ctx := context.Background()
	errorObject := `{"code":-32000,"message":"denied","message":"bad key ` + key + `"}`
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")

	t.Run("single", func(t *testing.T) {
		c := qt.New(t)
		url := rawEndpoint(t, "/rpc?apikey="+key, func(ids []json.RawMessage) []byte {
			return []byte(`{"jsonrpc":"2.0","id":` + string(ids[0]) + `,"error":` + errorObject + `}`)
		})
		pool, err := NewRPCPool([]string{url})
		c.Assert(err, qt.IsNil)
		defer pool.Close()
		_, err = NewPooledBackend(pool).CallContract(ctx, ethereum.CallMsg{To: &to}, nil)
		c.Assert(err, qt.IsNotNil)
		c.Assert(strings.Contains(err.Error(), key), qt.IsFalse, qt.Commentf("leaked: %v", err))
	})
	t.Run("batch", func(t *testing.T) {
		c := qt.New(t)
		url := rawEndpoint(t, "/rpc?apikey="+key, func(ids []json.RawMessage) []byte {
			return []byte(`[{"jsonrpc":"2.0","id":` + string(ids[0]) + `,"result":"0x01"},` +
				`{"jsonrpc":"2.0","id":` + string(ids[1]) + `,"error":` + errorObject + `}]`)
		})
		pool, err := NewRPCPool([]string{url})
		c.Assert(err, qt.IsNil)
		defer pool.Close()
		var ok, failed hexutil.Bytes
		batch := []rpc.BatchElem{
			{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &ok},
			{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &failed},
		}
		c.Assert(NewPooledBackend(pool).BatchCallContext(ctx, batch), qt.IsNil)
		c.Assert(ok.String(), qt.Equals, "0x01")
		c.Assert(batch[1].Error, qt.IsNotNil)
		c.Assert(strings.Contains(batch[1].Error.Error(), key), qt.IsFalse, qt.Commentf("leaked: %v", batch[1].Error))
	})
}

// go-ethereum matches JSON keys case-insensitively and unescapes them, so
// an error spelled with a "Message" property or an escaped "error" key is
// the same error to it. Its credential must not reach the error text either,
// single or in a batch, while the revert data still decodes and a result in
// the same batch keeps its exact bytes, credential-shaped bytes included.
func TestErrorKeySpellingsAreRedacted(t *testing.T) {
	const key = "deadbeefcafe"
	ctx := context.Background()
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	spellings := map[string]string{
		"Message alias": `{"code":3,"Message":"execution reverted: bad key ` + key + `","data":"0x08c379a0"}`,
		"escaped error": `{"code":3,"message":"execution reverted: bad key ` + key + `","data":"0x08c379a0"}`,
	}
	errorKey := map[string]string{"Message alias": `"error"`, "escaped error": `"\u0065rror"`}
	for name, object := range spellings {
		t.Run(name+"/single", func(t *testing.T) {
			c := qt.New(t)
			url := rawEndpoint(t, "/rpc?apikey="+key, func(ids []json.RawMessage) []byte {
				return []byte(`{"jsonrpc":"2.0","id":` + string(ids[0]) + `,` + errorKey[name] + `:` + object + `}`)
			})
			pool, err := NewRPCPool([]string{url})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			_, err = NewPooledBackend(pool).CallContract(ctx, ethereum.CallMsg{To: &to}, nil)
			c.Assert(err, qt.IsNotNil)
			c.Assert(strings.Contains(err.Error(), key), qt.IsFalse, qt.Commentf("leaked: %v", err))
			data, ok := ethclient.RevertErrorData(err)
			c.Assert(ok, qt.IsTrue, qt.Commentf("revert data lost: %v", err))
			c.Assert(data, qt.DeepEquals, []byte{0x08, 0xc3, 0x79, 0xa0})
		})
		t.Run(name+"/batch", func(t *testing.T) {
			c := qt.New(t)
			url := rawEndpoint(t, "/rpc?apikey="+key, func(ids []json.RawMessage) []byte {
				return []byte(`[{"jsonrpc":"2.0","id":` + string(ids[0]) + `,"result":"0x00` + key + `00"},` +
					`{"jsonrpc":"2.0","id":` + string(ids[1]) + `,` + errorKey[name] + `:` + object + `}]`)
			})
			pool, err := NewRPCPool([]string{url})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			var ok, failed hexutil.Bytes
			batch := []rpc.BatchElem{
				{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &ok},
				{Method: "eth_call", Args: []any{map[string]any{}, "latest"}, Result: &failed},
			}
			c.Assert(NewPooledBackend(pool).BatchCallContext(ctx, batch), qt.IsNil)
			c.Assert(ok.String(), qt.Equals, "0x00"+key+"00")
			c.Assert(batch[1].Error, qt.IsNotNil)
			c.Assert(strings.Contains(batch[1].Error.Error(), key), qt.IsFalse, qt.Commentf("leaked: %v", batch[1].Error))
			data, isRevert := ethclient.RevertErrorData(batch[1].Error)
			c.Assert(isRevert, qt.IsTrue)
			c.Assert(data, qt.DeepEquals, []byte{0x08, 0xc3, 0x79, 0xa0})
		})
	}
}

// fakeDataError is a JSON-RPC error with data, as go-ethereum's client
// returns for a revert (code 3, the revert bytes as hex data).
type fakeDataError struct {
	code int
	msg  string
	data any
}

func (e fakeDataError) Error() string  { return e.msg }
func (e fakeDataError) ErrorCode() int { return e.code }
func (e fakeDataError) ErrorData() any { return e.data }

// The output-side wrapper, on its own: an error carrying an endpoint's
// secret comes back with clean text while errors.Is and errors.As still
// reach what it wraps, and a JSON-RPC error keeps its code and revert data
// without exposing its raw message; an error without secrets is returned
// as it is.
func TestRedactorKeepsErrorBehaviour(t *testing.T) {
	const secret = "QKEY-0123456789"
	r := newRedactor("https://user:PASSWORD1@rpc.example/v2/SEGMENTKEY123?apikey=" + secret)
	secrets := []string{"user", "PASSWORD1", "SEGMENTKEY123", secret}
	clean := func(t *testing.T, err error) {
		t.Helper()
		for _, s := range secrets {
			qt.Assert(t, strings.Contains(err.Error(), s), qt.IsFalse, qt.Commentf("%s in %q", s, err.Error()))
		}
	}

	t.Run("sentinel", func(t *testing.T) {
		c := qt.New(t)
		err := r.error(fmt.Errorf("lookup with apikey=%s: %w", secret, ethereum.NotFound))
		clean(t, err)
		c.Assert(errors.Is(err, ethereum.NotFound), qt.IsTrue)
	})
	t.Run("revert", func(t *testing.T) {
		c := qt.New(t)
		inner := fakeDataError{code: 3, msg: "execution reverted: bad key " + secret + " for user:PASSWORD1", data: "0x08c379a0"}
		err := fmt.Errorf("call: %w", r.error(fmt.Errorf("eth_call: %w", inner)))
		clean(t, err)
		var rpcErr rpc.Error
		c.Assert(errors.As(err, &rpcErr), qt.IsTrue)
		c.Assert(rpcErr.ErrorCode(), qt.Equals, 3)
		clean(t, rpcErr)
		var dataErr rpc.DataError
		c.Assert(errors.As(err, &dataErr), qt.IsTrue)
		c.Assert(dataErr.ErrorData(), qt.Equals, "0x08c379a0")
		clean(t, dataErr)
		data, ok := ethclient.RevertErrorData(err)
		c.Assert(ok, qt.IsTrue)
		c.Assert(data, qt.DeepEquals, []byte{0x08, 0xc3, 0x79, 0xa0})
		c.Assert(IsRateLimitError(err), qt.IsFalse)
	})
	t.Run("rate limit code", func(t *testing.T) {
		err := r.error(fakeRPCError{code: rpcRateLimitCode, msg: "slow down, " + secret})
		clean(t, err)
		qt.Assert(t, IsRateLimitError(err), qt.IsTrue)
	})
	t.Run("http status", func(t *testing.T) {
		c := qt.New(t)
		err := r.error(rpc.HTTPError{StatusCode: http.StatusForbidden, Status: "403 Forbidden", Body: []byte("no access for /v2/SEGMENTKEY123")})
		clean(t, err)
		var httpErr rpc.HTTPError
		c.Assert(errors.As(err, &httpErr), qt.IsTrue)
		c.Assert(httpErr.StatusCode, qt.Equals, http.StatusForbidden)
		c.Assert(IsCapabilityError(err), qt.IsTrue)
	})
	t.Run("no secret", func(t *testing.T) {
		err := errors.New("execution reverted")
		qt.Assert(t, r.error(err), qt.Equals, err)
		qt.Assert(t, r.error(nil), qt.IsNil)
	})
}
