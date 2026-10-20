package battery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// The connect timeout error names the RPC without the API key its URL may
// carry in the path, query or userinfo; it needs no fleet.
func TestRPCTimeoutHidesTheEndpointSecrets(t *testing.T) {
	t.Setenv("BATTERY_CONNECT_TIMEOUT", "1ms")
	down := httptest.NewServer(http.NotFoundHandler())
	down.Close()
	url := strings.Replace(down.URL, "http://", "http://batuser:BATPASSWORD@", 1) + "/v2/BATPATHKEY?apikey=BATQUERYKEY"
	_, err := waitRPC(context.Background(), url)
	if err == nil {
		t.Fatal("waitRPC on a closed endpoint returned no error")
	}
	for _, secret := range []string{"batuser", "BATPASSWORD", "BATPATHKEY", "BATQUERYKEY"} {
		if strings.Contains(err.Error(), secret) {
			t.Errorf("%s leaked: %v", secret, err)
		}
	}
}

// A reachable endpoint whose errors echo the request (an HTTP 403 body, a
// JSON-RPC error message) until the connect deadline must not put the API
// key in the timeout error either.
func TestRPCTimeoutHidesProviderEchoes(t *testing.T) {
	t.Setenv("BATTERY_CONNECT_TIMEOUT", "1ms")
	for _, mode := range []string{"http", "jsonrpc"} {
		t.Run(mode, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				user, pass, _ := r.BasicAuth()
				echo := "no access for " + r.RequestURI + " as " + user + ":" + pass
				if mode == "http" {
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte(echo))
					return
				}
				var req struct {
					ID json.RawMessage `json:"id"`
				}
				_ = json.NewDecoder(r.Body).Decode(&req)
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"jsonrpc": "2.0", "id": req.ID, "error": map[string]any{"code": -32000, "message": echo},
				})
			}))
			defer srv.Close()
			url := strings.Replace(srv.URL, "http://", "http://batuser:BATPASSWORD@", 1) + "/v2/BATPATHKEY?apikey=BATQUERYKEY"
			_, err := waitRPC(context.Background(), url)
			if err == nil {
				t.Fatal("waitRPC on a refusing endpoint returned no error")
			}
			for _, secret := range []string{"batuser", "BATPASSWORD", "BATPATHKEY", "BATQUERYKEY"} {
				if strings.Contains(err.Error(), secret) {
					t.Errorf("%s leaked: %v", secret, err)
				}
			}
		})
	}
}
