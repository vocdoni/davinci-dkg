# Cross-impl test vectors

Canonical fixtures consumed by all four layers (Solidity / Go / TypeScript SDK / UI).
The Go side at `internal/protocol/protocol.go` and `crypto/schnorr/*` is the
single source of truth; this directory is **generated**.

## Files

| file               | covers                                                          |
|--------------------|-----------------------------------------------------------------|
| `protocol.json`    | transcript-domain digests + the BN254 / subgroup field constants |
| `schnorr.json`     | operator + organizer Schnorr proof-of-knowledge                  |
| `dleq.json`        | committee partial-decryption Chaum-Pedersen challenges + responses |

`protocol.json` carries two kinds of domain rows, each with its UTF-8
preimage, its keccak256 and the digest reduced into the BN254 scalar field:

- Schnorr registration transcript domains (`OperatorRegisterV1`,
  `OrganizerRegisterV1`) and the in-circuit partial-decrypt domain
  (`PartialDecryptCircuit`, consumed via `SetBytes`, not keccak'd).
- The three BRLC transcript domains every proof-carrying call binds into its
  Fiat–Shamir challenge `keccak(eid ‖ domain ‖ anchor) mod p` (see
  `BRLC.deriveChallenge`): `ContributionTranscriptV1`
  (`davinci-dkg:contribution:v1`, `submitContribution`), `PoolKeyTranscriptV1`
  (`davinci-dkg:poolkey:v1`, `activatePoolKey` — replaces the former
  `davinci-dkg:finalize:v1`) and `DecryptCombineTranscriptV1`
  (`davinci-dkg:decrypt-combine:v1`, `combineDecryption`). Their source is
  `internal/protocol/protocol.go`; the circuits' witness builders and the
  `*_TRANSCRIPT_DOMAIN` constants in `DKGManager.sol` must hash the same
  strings.

## Regenerating

```
make vectors          # writes tests/vectors/*.json
make vectors-check    # regenerate and fail if the working tree changed
```

The generator is `cmd/protocol-vectors/main.go`. Determinism is load-bearing:
re-running on a clean checkout must produce byte-identical files. CI runs
`make vectors-check` to catch silent drift.

## Consumers

- **SDK** — `sdk/tests/vectors.test.ts` reads every file and asserts the
  TypeScript verifiers reproduce each value.
- **Solidity** — `solidity/test/DKGProtocol.t.sol` mirrors the digests and
  values inline; SDK vector tests fail if Solidity drifts because both
  layers re-derive `c` over the same transcript.
- **Go** — `crypto/schnorr/*` and `internal/protocol` are the canonical
  sources; the generator simply re-emits values from those packages, so any
  Go-side drift breaks the generator output (and therefore the SDK tests).
  `tests/helpers` additionally checks that the pool-key circuit derives its
  challenge under `PoolKeyTranscriptV1`.
- **UI** — `ui/tests/vectors/` is a byte-for-byte mirror (`make vectors`
  refreshes it); `ui/src/lib/protocol-vectors.test.ts` fails when the mirror
  or the SDK constants drift.

## Adding a new vector type

1. Add a `build*()` function in `cmd/protocol-vectors/main.go`.
2. Add a `write(*dir, "name.json", ...)` call in `main()`.
3. Add a matching `describe('vectors / name.json', ...)` block in
   `sdk/tests/vectors.test.ts`.
4. Run `make vectors && pnpm --filter sdk test`.
