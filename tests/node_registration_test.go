package tests

import (
	"context"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
)

func TestHarnessRegistersNodeKeys(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	for i := 0; i < 3; i++ {
		actor, err := services.Actor(i)
		c.Assert(err, qt.IsNil)

		node, err := services.Contracts.GetNode(ctx, actor.Address())
		c.Assert(err, qt.IsNil)
		c.Assert(node.Operator, qt.Equals, actor.Address())
		c.Assert(node.PubX.Sign(), qt.Equals, 1)
		c.Assert(node.PubY.Sign(), qt.Equals, 1)
	}
}
