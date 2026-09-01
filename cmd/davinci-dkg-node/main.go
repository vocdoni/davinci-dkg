package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vocdoni/davinci-dkg/internal/version"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/node"
)

func main() {
	cfg, err := node.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading configuration: %v\n", err)
		os.Exit(1)
	}

	log.Init(cfg.Log.Level, cfg.Log.Output, nil)
	log.Infow("starting davinci-dkg-node", "version", version.Version, "network", cfg.ResolvedNetworkName())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !cfg.HasChainConfig() {
		log.Infow("no chain config provided — node is idle (set --privkey, --manager to enable participation)")
		waitForSignal()
		return
	}

	n, err := node.New(cfg)
	if err != nil {
		log.Errorw(err, "failed to initialize node")
		os.Exit(1)
	}

	if err := n.EnsureRegistered(ctx); err != nil {
		log.Errorw(err, "key registration failed")
		os.Exit(1)
	}

	// Emit the full startup banner *after* EnsureRegistered so the on-chain
	// snapshot reflects the post-registration state (status=ACTIVE, fresh
	// lastActiveBlock).
	n.LogStartupSnapshot(ctx, cfg)

	go n.Run(ctx, cfg)
	waitForSignal()
}

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	log.Infow("shutdown signal received")
}
