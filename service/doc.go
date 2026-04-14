// Package service implements the per-phase workers that make up a running
// davinci-dkg node.
//
// Each phase of the DKG protocol — registration, contribution, finalization,
// partial decryption, decryption combination, reveal submission, secret
// reconstruction — is handled by a dedicated worker type. A round monitor
// watches on-chain state and routes each active round to the workers whose
// phase capabilities match the round's current status.
//
// Workers are stateless across rounds; all persistent state lives in the
// storage package. This makes it safe to restart the node at any point —
// any work that was already committed on-chain is simply skipped on resume.
package service
