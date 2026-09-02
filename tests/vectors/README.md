# Cross-impl test vectors

Canonical fixtures consumed by all four layers (Solidity / Go / TypeScript SDK / UI).
The Go side at `internal/protocol/protocol.go` and `crypto/schnorr/*` is the
single source of truth; this directory is **generated**.

## Files

| file               | covers                                                          |
|--------------------|-----------------------------------------------------------------|
| `protocol.json`    | transcript-domain digests + the organizer-share DLEQ encoding    |
| `schnorr.json`     | operator + organizer Schnorr proof-of-knowledge                  |
| `dleq.json`        | committee partial-decryption Chaum-Pedersen challenges + responses |

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
- **Go** — `crypto/schnorr/*` is the canonical source; the generator simply
  re-emits values from those packages, so any Go-side drift breaks the
  generator output (and therefore the SDK tests).

## Adding a new vector type

1. Add a `build*()` function in `cmd/protocol-vectors/main.go`.
2. Add a `write(*dir, "name.json", ...)` call in `main()`.
3. Add a matching `describe('vectors / name.json', ...)` block in
   `sdk/tests/vectors.test.ts`.
4. Run `make vectors && pnpm --filter sdk test`.
