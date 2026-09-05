// constraints prints the Groth16 R1CS constraint count for each of the
// four production circuits at the current MaxN setting. Used to refresh
// the constraint-count table in BENCHMARKS.md.
package main

import (
	"fmt"

	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/circuits/poolkey"
)

func main() {
	c, err := contribution.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contribution    %d\n", c.GetNbConstraints())

	k, err := poolkey.Compile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("PoolKey         %d\n", k.GetNbConstraints())

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
