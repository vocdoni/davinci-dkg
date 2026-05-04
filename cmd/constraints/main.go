// constraints prints the Groth16 R1CS constraint count for each of the
// four production circuits at the current MaxN setting. Used to refresh
// the constraint-count table in BENCHMARKS.md.
package main

import (
	"fmt"

	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
)

func main() {
	c, err := contribution.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contribution    %d\n", c.GetNbConstraints())

	f, err := finalize.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Finalize        %d\n", f.GetNbConstraints())

	p, err := partialdecrypt.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("PartialDecrypt  %d\n", p.GetNbConstraints())

	d, err := decryptcombine.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DecryptCombine  %d\n", d.GetNbConstraints())
}
