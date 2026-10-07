package main

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestRegistrationAid(t *testing.T) {
	me := common.HexToAddress("0x00000000000000000000000000000000000A11CE")
	other := common.HexToAddress("0x000000000000000000000000000000000000BEEF")

	random, err := registrationAid("", me)
	if err != nil || types.ApplicationRegistrant(random) != me {
		t.Fatalf("random id: got %x, %v", random, err)
	}

	// salt 7, registrant 0xA11CE
	mine := "0x" + "000000000000000000000007" + "00000000000000000000000000000000000a11ce"
	got, err := registrationAid(mine, me)
	if err != nil || types.ApplicationRegistrant(got) != me || got[11] != 7 {
		t.Fatalf("own id: got %x, %v", got, err)
	}
	if _, err := registrationAid(mine, other); err == nil {
		t.Fatal("an id in another account's namespace must be rejected")
	}
	// The pre-#14 shape, a random id with the top bits cleared, is refused.
	if _, err := registrationAid("0x1fab"+strings.Repeat("00", 30), me); err == nil {
		t.Fatal("an unbound id must be rejected")
	}
	if _, err := registrationAid("0x00", me); err == nil {
		t.Fatal("the zero id must be rejected")
	}
}
