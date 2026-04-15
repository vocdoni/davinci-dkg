package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/webapp"
)

// runtimeConfig is the JSON document served at /config.json. The browser
// fetches it on startup so the explorer knows which contracts and RPC to talk
// to without any rebuild.
type runtimeConfig struct {
	RPCURL          string `json:"rpcUrl"`
	ManagerAddress  string `json:"managerAddress"`
	RegistryAddress string `json:"registryAddress"`
	ChainID         uint64 `json:"chainId"`
	ChainName       string `json:"chainName"`
}

// startWebapp serves the embedded explorer SPA on cfg.Webapp.Listen. It returns
// the *http.Server so callers can shut it down with the rest of the node.
// Returns nil if the webapp is disabled.
// registryAddress is the resolved DKGRegistry address (may be empty in idle mode).
func startWebapp(ctx context.Context, cfg *Config, chainID uint64, registryAddress string) (*http.Server, error) {
	if !cfg.Webapp.Enabled {
		log.Infow("webapp disabled")
		return nil, nil
	}
	assets, err := webapp.Assets()
	if err != nil {
		return nil, err
	}

	rc := runtimeConfig{
		ManagerAddress:  cfg.resolvedManagerAddr(),
		RegistryAddress: registryAddress,
		ChainID:         chainID,
		ChainName:       cfg.resolvedNetworkName(),
	}
	if len(cfg.Web3.RPC) > 0 {
		rc.RPCURL = cfg.Web3.RPC[0]
	}
	if cfg.Webapp.PublicRPC != "" {
		rc.RPCURL = cfg.Webapp.PublicRPC
	}
	if rc.ChainID == 0 && len(cfg.Web3.RPC) > 0 {
		if id, err := fetchChainID(cfg.Web3.RPC[0]); err == nil {
			rc.ChainID = id
		} else {
			log.Warnw("failed to fetch chain id for webapp config", "err", err)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		_ = json.NewEncoder(w).Encode(rc)
	})
	mux.Handle("/", spaHandler(assets))

	srv := &http.Server{
		Addr:              cfg.Webapp.Listen,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	go func() {
		log.Infow("webapp listening", "addr", cfg.Webapp.Listen)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorw(err, "webapp server stopped")
		}
	}()
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()
	return srv, nil
}

// fetchChainID queries eth_chainId via a minimal JSON-RPC POST so the webapp
// can display the right chain even when the node is running in idle mode.
func fetchChainID(rpcURL string) (uint64, error) {
	req := []byte(`{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}`)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, rpcURL, bytes.NewReader(req))
	if err != nil {
		return 0, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var out struct {
		Result string `json:"result"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return 0, err
	}
	s := strings.TrimPrefix(out.Result, "0x")
	return strconv.ParseUint(s, 16, 64)
}

// spaHandler serves the embedded assets and falls back to index.html for any
// path that does not map to an asset, so client-side routing works.
func spaHandler(assets fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(assets))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clean := strings.TrimPrefix(r.URL.Path, "/")
		if clean == "" {
			fileServer.ServeHTTP(w, r)
			return
		}
		if f, err := assets.Open(clean); err == nil {
			_ = f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}
		// SPA fallback: rewrite to /index.html so React Router can render the
		// route on the client.
		r2 := r.Clone(r.Context())
		r2.URL.Path = "/"
		fileServer.ServeHTTP(w, r2)
	})
}
