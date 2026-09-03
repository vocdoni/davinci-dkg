// Package backendbench compares the Groth16 and PLONK backends on one
// circuit: constraint counts, setup, proving and verification time, and
// key/proof sizes. It exists to answer "is the universal setup worth it?"
// with numbers; nothing in production imports it.
package backendbench

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

// Row is one backend's measurements for one circuit.
type Row struct {
	Circuit, Backend      string
	Constraints           int
	Compile, Setup, Prove time.Duration
	Verify                time.Duration
	ProofBytes, PKBytes   int64
	VKBytes, SRSBytes     int64
}

func (r Row) String() string {
	return fmt.Sprintf("%-16s %-8s constraints=%9d compile=%6.1fs setup=%7.1fs prove=%7.2fs verify=%5.1fms proof=%5dB pk=%7.1fMB vk=%6dB srs=%7.1fMB",
		r.Circuit, r.Backend, r.Constraints, r.Compile.Seconds(), r.Setup.Seconds(), r.Prove.Seconds(),
		float64(r.Verify.Microseconds())/1000, r.ProofBytes, float64(r.PKBytes)/1e6, r.VKBytes, float64(r.SRSBytes)/1e6)
}

type counter struct{ n int64 }

func (c *counter) Write(p []byte) (int, error) { c.n += int64(len(p)); return len(p), nil }

func size(w io.WriterTo) int64 {
	c := &counter{}
	_, _ = w.WriteTo(c)
	return c.n
}

// Compare runs both backends on circuit with the given full assignment and
// prints one row per backend. `runs` proofs are timed and averaged.
func Compare(b *testing.B, name string, circuit, assignment frontend.Circuit, runs int) {
	b.Helper()
	field := ecc.BN254.ScalarField()
	full, err := frontend.NewWitness(assignment, field)
	if err != nil {
		b.Fatal(err)
	}
	public, err := full.Public()
	if err != nil {
		b.Fatal(err)
	}

	// ── Groth16 ────────────────────────────────────────────────────────────
	g := Row{Circuit: name, Backend: "groth16"}
	t0 := time.Now()
	ccs, err := frontend.Compile(field, r1cs.NewBuilder, circuit)
	if err != nil {
		b.Fatal(err)
	}
	g.Compile = time.Since(t0)
	g.Constraints = ccs.GetNbConstraints()
	t0 = time.Now()
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		b.Fatal(err)
	}
	g.Setup = time.Since(t0)
	g.PKBytes, g.VKBytes = size(pk), size(vk)
	var proof groth16.Proof
	t0 = time.Now()
	for i := 0; i < runs; i++ {
		if proof, err = groth16.Prove(ccs, pk, full); err != nil {
			b.Fatal(err)
		}
	}
	g.Prove = time.Since(t0) / time.Duration(runs)
	g.ProofBytes = size(proof)
	t0 = time.Now()
	if err := groth16.Verify(proof, vk, public); err != nil {
		b.Fatal(err)
	}
	g.Verify = time.Since(t0)
	report(b, g)
	ccs, pk, vk, proof = nil, nil, nil, nil //nolint:ineffassign,staticcheck // release before PLONK

	// ── PLONK ──────────────────────────────────────────────────────────────
	p := Row{Circuit: name, Backend: "plonk"}
	t0 = time.Now()
	scsCCS, err := frontend.Compile(field, scs.NewBuilder, circuit)
	if err != nil {
		b.Fatal(err)
	}
	p.Compile = time.Since(t0)
	p.Constraints = scsCCS.GetNbConstraints()
	t0 = time.Now()
	srs, srsLagrange, err := unsafekzg.NewSRS(scsCCS)
	if err != nil {
		b.Fatal(err)
	}
	srsTime := time.Since(t0)
	p.SRSBytes = size(srs) + size(srsLagrange)
	t0 = time.Now()
	ppk, pvk, err := plonk.Setup(scsCCS, srs, srsLagrange)
	if err != nil {
		b.Fatal(err)
	}
	p.Setup = time.Since(t0)
	p.PKBytes, p.VKBytes = size(ppk), size(pvk)
	var pproof plonk.Proof
	t0 = time.Now()
	for i := 0; i < runs; i++ {
		if pproof, err = plonk.Prove(scsCCS, ppk, full); err != nil {
			b.Fatal(err)
		}
	}
	p.Prove = time.Since(t0) / time.Duration(runs)
	p.ProofBytes = size(pproof)
	t0 = time.Now()
	if err := plonk.Verify(pproof, pvk, public); err != nil {
		b.Fatal(err)
	}
	p.Verify = time.Since(t0)
	report(b, p)
	fmt.Fprintf(os.Stderr, "%-16s plonk    srs generation (unsafe, test only) %.1fs, domain size %d\n", name, srsTime.Seconds(), domainSize(scsCCS))
}

func domainSize(ccs constraint.ConstraintSystem) int {
	n := ccs.GetNbConstraints() + ccs.GetNbPublicVariables()
	d := 1
	for d < n {
		d <<= 1
	}
	return d
}

func report(b *testing.B, r Row) {
	b.Helper()
	fmt.Fprintln(os.Stderr, r.String())
}
