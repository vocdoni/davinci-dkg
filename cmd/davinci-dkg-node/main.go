package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vocdoni/davinci-dkg/internal/version"
	"github.com/vocdoni/davinci-dkg/log"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading configuration: %v\n", err)
		os.Exit(1)
	}

	log.Init(cfg.Log.Level, cfg.Log.Output, nil)
	log.Infow("starting davinci-dkg-node", "version", version.Version, "network", cfg.resolvedNetworkName())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !cfg.HasChainConfig() {
		log.Infow("no chain config provided — node is idle (set --privkey, --manager to enable participation)")
		if _, err := startWebapp(ctx, cfg, 0, "", ""); err != nil {
			log.Errorw(err, "failed to start webapp")
		}
		waitForSignal()
		return
	}

	node, err := newNode(cfg)
	if err != nil {
		log.Errorw(err, "failed to initialize node")
		os.Exit(1)
	}

	if _, err := startWebapp(ctx, cfg, node.contracts.ChainID, node.contracts.Addresses.Registry.Hex(), cfg.SharedDir); err != nil {
		log.Errorw(err, "failed to start webapp")
	}

	if err := node.EnsureRegistered(ctx); err != nil {
		log.Errorw(err, "key registration failed")
		os.Exit(1)
	}

	// Emit the full startup banner *after* EnsureRegistered so the on-chain
	// snapshot reflects the post-registration state (status=ACTIVE, fresh
	// lastActiveBlock).
	node.LogStartupSnapshot(ctx, cfg)

	go node.Run(ctx, cfg.PollInterval)
	waitForSignal()
}

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	log.Infow("shutdown signal received")
}
