package contribution

import (
	"context"
	"testing"
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
