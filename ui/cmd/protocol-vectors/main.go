package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

type Vector struct {
	ModePublicDerivation         uint8  `json:"mode_public_derivation"`
	ModeOrganizerCoDec           uint8  `json:"mode_organizer_codec"`
	RoleCommittee                uint8  `json:"role_committee"`
	RoleOrganizer                uint8  `json:"role_organizer"`
	DomainOperatorRegisterV1Str  string `json:"domain_operator_register_v1_str"`
	DomainOperatorRegisterV1     string `json:"domain_operator_register_v1"`
	DomainOrganizerRegisterV1Str string `json:"domain_organizer_register_v1_str"`
	DomainOrganizerRegisterV1    string `json:"domain_organizer_register_v1"`
	DomainDLEQV1Str              string `json:"domain_dleq_v1_str"`
	DomainDLEQV1                 string `json:"domain_dleq_v1"`
}

func main() {
	v := Vector{
		ModePublicDerivation:         uint8(protocol.ModePublicDerivation),
		ModeOrganizerCoDec:           uint8(protocol.ModeOrganizerCoDec),
		RoleCommittee:                uint8(protocol.RoleCommittee),
		RoleOrganizer:                uint8(protocol.RoleOrganizer),
		DomainOperatorRegisterV1Str:  protocol.DomainOperatorRegisterV1Str,
		DomainOperatorRegisterV1:     "0x" + hex.EncodeToString(protocol.DomainOperatorRegisterV1[:]),
		DomainOrganizerRegisterV1Str: protocol.DomainOrganizerRegisterV1Str,
		DomainOrganizerRegisterV1:    "0x" + hex.EncodeToString(protocol.DomainOrganizerRegisterV1[:]),
		DomainDLEQV1Str:              protocol.DomainDLEQV1Str,
		DomainDLEQV1:                 "0x" + hex.EncodeToString(protocol.DomainDLEQV1[:]),
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
