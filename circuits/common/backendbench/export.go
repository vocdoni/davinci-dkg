package backendbench

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/backend/plonk"
	plonkbn254 "github.com/consensys/gnark/backend/plonk/bn254"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

// ExportSolidity writes, for one circuit and assignment, a Groth16 and a
// PLONK Solidity verifier plus a proof and the public inputs each accepts,
// so a Foundry test can measure on-chain verification gas for both.
func ExportSolidity(dir string, circuit, assignment frontend.Circuit) error {
	field := ecc.BN254.ScalarField()
	full, err := frontend.NewWitness(assignment, field)
	if err != nil {
		return err
	}
	public, err := full.Public()
	if err != nil {
		return err
	}
	inputs := publicInputs(public)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "inputs.json"), inputs); err != nil {
		return err
	}

	// Groth16
	ccs, err := frontend.Compile(field, r1cs.NewBuilder, circuit)
	if err != nil {
		return err
	}
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		return err
	}
	// The exported Solidity verifier hashes the commitment with sha256.
	proof, err := groth16.Prove(ccs, pk, full, backend.WithProverHashToFieldFunction(sha256.New()))
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, "Groth16Verifier.sol"))
	if err != nil {
		return err
	}
	if err := vk.(*groth16bn254.VerifyingKey).ExportSolidity(f, solidity.WithPragmaVersion("^0.8.28")); err != nil {
		return err
	}
	_ = f.Close()
	if err := os.WriteFile(filepath.Join(dir, "groth16-proof.hex"),
		[]byte(hex.EncodeToString(proof.(*groth16bn254.Proof).MarshalSolidity())), 0o644); err != nil {
		return err
	}

	// PLONK
	sccs, err := frontend.Compile(field, scs.NewBuilder, circuit)
	if err != nil {
		return err
	}
	srs, srsL, err := unsafekzg.NewSRS(sccs)
	if err != nil {
		return err
	}
	ppk, pvk, err := plonk.Setup(sccs, srs, srsL)
	if err != nil {
		return err
	}
	pproof, err := plonk.Prove(sccs, ppk, full)
	if err != nil {
		return err
	}
	f, err = os.Create(filepath.Join(dir, "PlonkVerifier.sol"))
	if err != nil {
		return err
	}
	if err := pvk.(*plonkbn254.VerifyingKey).ExportSolidity(f, solidity.WithPragmaVersion("^0.8.28")); err != nil {
		return err
	}
	_ = f.Close()
	if err := os.WriteFile(filepath.Join(dir, "plonk-proof.hex"),
		[]byte(hex.EncodeToString(pproof.(*plonkbn254.Proof).MarshalSolidity())), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "exported %d public inputs, r1cs=%d scs=%d to %s\n",
		len(inputs), ccs.GetNbConstraints(), sccs.GetNbConstraints(), dir)
	return nil
}

func publicInputs(public interface{ Vector() any }) []string {
	vec, ok := public.Vector().(fr.Vector)
	if !ok {
		return nil
	}
	out := make([]string, 0, len(vec))
	for i := range vec {
		out = append(out, vec[i].BigInt(new(big.Int)).String())
	}
	return out
}

func writeJSON(path string, v any) error {
	raw, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}
