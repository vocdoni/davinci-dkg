# UI cross-impl vectors

Byte-for-byte copies of `tests/vectors/*.json`, which `cmd/protocol-vectors`
generates from the Go side (`make vectors`). They are mirrored here so the UI
test suite can pin the protocol constants it renders and the encodings the SDK
performs in the browser, without reaching outside the `ui/` package at runtime.

`src/lib/protocol-vectors.test.ts` asserts both that the SDK reproduces every
value and that this copy is identical to `tests/vectors/`, so a stale mirror
fails the UI suite rather than silently drifting.

Regenerate with `make vectors`, then copy the files here. Do not edit by hand.
