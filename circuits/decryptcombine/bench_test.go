package decryptcombine

import (
	"context"
	"os"
	"testing"

	"github.com/vocdoni/davinci-dkg/circuits/common/backendbench"
)

func BenchmarkProve(b *testing.B) {
	witness, _, err := BuildWitness(testAssignment())
	if err != nil {
		b.Fatal(err)
	}
	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &DecryptCombineCircuit{})
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
	backendbench.Compare(b, "decryptcombine", &DecryptCombineCircuit{}, witness, 3)
}

// TestExportBackendVerifiers writes Groth16 and PLONK Solidity verifiers with
// a matching proof to $BACKEND_EXPORT_DIR for the Foundry gas comparison.
func TestExportBackendVerifiers(t *testing.T) {
	dir := os.Getenv("BACKEND_EXPORT_DIR")
	if dir == "" {
		t.Skip("BACKEND_EXPORT_DIR not set")
	}
	witness, _, err := BuildWitness(testAssignment())
	if err != nil {
		t.Fatal(err)
	}
	if err := backendbench.ExportSolidity(dir, &DecryptCombineCircuit{}, witness); err != nil {
		t.Fatal(err)
	}
}
