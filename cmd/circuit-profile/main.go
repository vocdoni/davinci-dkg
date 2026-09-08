// circuit-profile compiles each circuit under gnark's constraint profiler and
// writes /tmp/<circuit>.pprof; read it with
// `go tool pprof -sample_index=0 -top -nodecount=40 /tmp/<circuit>.pprof`.
// Pass a circuit name to profile only that one.
package main

import (
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/profile"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
)

func main() {
	circuitsByName := []struct {
		name string
		c    frontend.Circuit
	}{
		{"contribution", &contribution.ContributionCircuit{}},
		{"finalize", &finalize.FinalizeCircuit{}},
		{"decryptcombine", &decryptcombine.DecryptCombineCircuit{}},
		{"partialdecrypt", &partialdecrypt.PartialDecryptCircuit{}},
	}
	for _, e := range circuitsByName {
		if len(os.Args) > 1 && os.Args[1] != e.name {
			continue
		}
		p := profile.Start(profile.WithPath("/tmp/" + e.name + ".pprof"))
		ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, e.c)
		p.Stop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("==== %s: %d constraints\n", e.name, ccs.GetNbConstraints())
		fmt.Printf("%d samples written to /tmp/%s.pprof\n", p.NbConstraints(), e.name)
	}
}
