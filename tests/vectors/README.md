# Cross-impl test vectors

Canonical fixtures consumed by all four layers (Solidity / Go / TypeScript SDK / UI).
The Go side at `internal/protocol/protocol.go`, `crypto/schnorr/*`,
`circuits/contribution` and `circuits/finalize` is the single source of truth;
this directory is **generated**.

## Files

| file                         | covers                                                                  |
|------------------------------|-------------------------------------------------------------------------|
| `protocol.json`              | transcript-domain digests + the BN254 / subgroup field constants         |
| `schnorr.json`               | operator + organizer Schnorr proof-of-knowledge                          |
| `dleq.json`                  | committee partial-decryption Chaum-Pedersen challenges + responses       |
| `contribution_compact.json`  | compact contribution transcripts (v4 §3–§5): words, offsets, digests, anchor, challenge, BRLC |
| `finalize_transcript.json`   | batched finalization transcripts (v4 §6–§9): words, Poseidon digest levels, anchor, challenge, BRLC, Merkle roots |

`protocol.json` carries two kinds of domain rows, each with its UTF-8
preimage, its keccak256 and the digest reduced into the BN254 scalar field:

- Schnorr registration transcript domains (`OperatorRegisterV1`,
  `OrganizerRegisterV1`) and the in-circuit partial-decrypt domain
  (`PartialDecryptCircuit`, consumed via `SetBytes`, not keccak'd).
- The three BRLC transcript domains every proof-carrying call binds into its
  Fiat–Shamir challenge `keccak(eid ‖ domain ‖ anchor) mod p` (see
  `BRLC.deriveChallenge`): `ContributionTranscriptV2`
  (`davinci-dkg:contribution:v2`, `submitContribution`), `FinalizeTranscriptV2`
  (`davinci-dkg:finalize:v2`, `finalizeEpoch` — replaces the former
  `davinci-dkg:poolkey:v1` of `activatePoolKey`) and
  `DecryptCombineTranscriptV1` (`davinci-dkg:decrypt-combine:v1`,
  `combineDecryption`). Their source is `internal/protocol/protocol.go`; the
  circuits' witness builders and the `*_TRANSCRIPT_DOMAIN` constants in
  `DKGManager.sol` must hash the same strings.

`contribution_compact.json` (docs/pool-keys-v4.md §3–§5): for each `(t, n,
contributorIndex)` case, the recipient secrets and nonces that regenerate the
committee keys and ephemerals, the `MaxK × t` coefficients, the plaintext
shares, the exact `L_C = MaxK·(2t+n) + 5n` transcript words (decimal, no
padding), the region offsets, `keccak256` of the committee region
`[2Kt, 2Kt+3n)` (what `_snapshotCommittee` must equal), the transcript
keccak, both Poseidon digests, the anchor
`keccak(commitmentsHash ‖ encryptedSharesHash ‖ keccak(transcript))`, the
challenge, the BRLC commitment and the eight public inputs in verifier order.

`finalize_transcript.json` (docs/pool-keys-v4.md §6–§9): for each accepted
set (contiguous, non-contiguous with a silent member, descending order with
`a = t`), the dealers' coefficients and stored `commitmentsHash`, all `MaxK`
pool keys, the `n` share commitments per key, the fixed `L_F = 1120`-word
transcript, the three digest levels `R`, `B_j`, `T`, the anchor
`keccak(transcriptDigest ‖ keccak(transcript))`, the challenge, the BRLC
commitment, the seven public inputs in verifier order and the `MaxK` keccak
Merkle roots (`leaf = keccak(0x00 ‖ x ‖ y)`, empty leaf
`keccak("davinci-dkg:merkle-empty:v1")`, node `keccak(0x01 ‖ l ‖ r)`).

## Regenerating

```
make vectors          # writes tests/vectors/*.json and mirrors them into ui/tests/vectors/
make vectors-check    # regenerate and fail if the working tree changed
```

The generator is `cmd/protocol-vectors/` (`main.go`, `v4.go`). Determinism is
load-bearing: re-running on a clean checkout must produce byte-identical
files. CI runs `make vectors-check` to catch silent drift.

## Consumers

- **SDK** — `sdk/tests/vectors.test.ts` reads every file and asserts the
  TypeScript verifiers reproduce each value.
- **Solidity** — `solidity/test/DKGProtocol.t.sol` mirrors the digests and
  values inline; SDK vector tests fail if Solidity drifts because both
  layers re-derive `c` over the same transcript.
- **Go** — `crypto/schnorr/*`, `internal/protocol`, `circuits/contribution`
  and `circuits/finalize` are the canonical sources; the generator simply
  re-emits values from those packages, so any Go-side drift breaks the
  generator output (and therefore the SDK tests).
- **UI** — `ui/tests/vectors/` is a byte-for-byte mirror (`make vectors`
  refreshes it); `ui/src/lib/protocol-vectors.test.ts` fails when the mirror
  or the SDK constants drift.

## Adding a new vector type

1. Add a `build*()` function in `cmd/protocol-vectors/`.
2. Add a `write(*dir, "name.json", ...)` call in `main()`.
3. Add a matching `describe('vectors / name.json', ...)` block in
   `sdk/tests/vectors.test.ts`.
4. Run `make vectors && pnpm --filter sdk test`.
