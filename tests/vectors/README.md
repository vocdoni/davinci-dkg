# Cross-implementation vectors

Fixtures shared by the Go, Solidity, TypeScript SDK and explorer tests. They are generated from
the Go code (`internal/protocol`, `crypto/schnorr`, `circuits/contribution`, `circuits/finalize`)
and must not be edited by hand.

| File | Covers |
|---|---|
| `protocol.json` | Transcript domain digests and the BN254 and subgroup constants |
| `schnorr.json` | Operator and organizer Schnorr proofs of possession |
| `dleq.json` | Chaum–Pedersen challenges and responses of committee partial decryptions |
| `contribution_compact.json` | Compact contribution transcripts: words, offsets, digests, anchor, challenge, BRLC |
| `finalize_transcript.json` | Finalization transcripts: words, Poseidon digest levels, anchor, challenge, BRLC, Merkle roots |

`protocol.json` holds each domain's UTF-8 preimage, its keccak256 and the digest reduced into the
BN254 scalar field: the Schnorr registration domains (`OperatorRegisterV1`,
`OrganizerRegisterV1`), the in-circuit partial-decryption domain (`PartialDecryptCircuit`, used as
a field element, not hashed), and the three BRLC domains every proof-carrying call binds into its
challenge `keccak(eid ‖ domain ‖ anchor) mod p`: `ContributionTranscriptV2`
(`davinci-dkg:contribution:v2`), `FinalizeTranscriptV2` (`davinci-dkg:finalize:v2`) and
`DecryptCombineTranscriptV1` (`davinci-dkg:decrypt-combine:v1`).

`contribution_compact.json` has, for each `(t, n, contributorIndex)` case, the recipient secrets
and nonces, the `MaxK × t` coefficients, the plaintext shares, the exact `MaxK·(2t+n) + 5n`
transcript words, the region offsets, the keccak256 of the committee region, the transcript
keccak, both Poseidon digests, the anchor, the challenge, the BRLC commitment and the eight public
inputs in verifier order.

`finalize_transcript.json` has, for each accepted set (contiguous, with a silent member, in
descending order with `a = t`), the dealers' coefficients and stored `commitmentsHash`, all `MaxK`
pool keys, the share commitments per key, the 1,120-word transcript, the digest levels `R`, `B_j`
and `T`, the anchor, the challenge, the BRLC commitment, the seven public inputs and the `MaxK`
Merkle roots. The layouts are specified in [docs/pool-keys.md](../../docs/pool-keys.md).

## Regenerating

```bash
make vectors          # writes tests/vectors/*.json and copies them to ui/tests/vectors/
make vectors-check    # regenerates and fails if the files changed
```

The generator is `cmd/protocol-vectors`. Its output is deterministic: running it on a clean
checkout must produce identical files.

## Consumers

- SDK: `sdk/tests/vectors.test.ts` reproduces every value.
- Solidity: `solidity/test/DKGProtocol.t.sol` mirrors the digests inline.
- Explorer: `ui/tests/vectors/` is an identical copy, checked by
  `ui/src/lib/protocol-vectors.test.ts`.
- Go: the generator re-emits values from the packages above, so Go-side drift changes the files.

To add a vector type, add a builder and a `write` call in `cmd/protocol-vectors`, a matching
`describe` block in `sdk/tests/vectors.test.ts`, then run `make vectors` and the SDK tests.
