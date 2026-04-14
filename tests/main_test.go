package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
)

var services *helpers.TestServices

func TestMain(m *testing.M) {
	if !helpers.IsIntegrationEnabled() {
		log.Info("skipping integration tests...")
		os.Exit(0)
	}

	log.Init("debug", "stdout", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var cleanup func()
	var err error
	services, cleanup, err = helpers.NewTestServices(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup integration services: %v\n", err)
		os.Exit(1)
	}

	code := m.Run()
	cleanup()
	os.Exit(code)
}
