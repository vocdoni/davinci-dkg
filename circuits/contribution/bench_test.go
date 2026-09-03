package contribution

import (
	"context"
	"testing"

	"github.com/vocdoni/davinci-dkg/circuits/common/backendbench"
)

func BenchmarkProve(b *testing.B) {
	witness, _, err := BuildWitness(testAssignment())
	if err != nil {
		b.Fatal(err)
	}
	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := runtime.ProveAndVerify(witness); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkBackends compares Groth16 and PLONK on this circuit (see
// circuits/common/backendbench). Run with -bench BenchmarkBackends -benchtime=1x.
func BenchmarkBackends(b *testing.B) {
	witness, _, err := BuildWitness(testAssignment())
	if err != nil {
		b.Fatal(err)
	}
	backendbench.Compare(b, "contribution", &ContributionCircuit{}, witness, 3)
}
