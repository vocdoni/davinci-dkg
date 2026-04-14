// Package prover wraps the gnark Groth16 backend with a disk-backed
// artifact cache.
//
// Proving keys and constraint systems are large (hundreds of MiB for the
// contribution circuit) and expensive to materialize, so the package caches
// them on disk keyed by circuit hash under $DAVINCI_DKG_ARTIFACTS_DIR. A
// cold node pays the setup cost once; subsequent boots hit the cache.
//
// Build with the gpu tag to enable an experimental CUDA-accelerated prover
// on supported hardware. Without the tag the GPU entry points return a
// clear error at runtime.
package prover
