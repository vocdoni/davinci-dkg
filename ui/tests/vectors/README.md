# Explorer copy of the cross-implementation vectors

Identical copies of `tests/vectors/*.json`, kept inside `ui/` so the explorer's tests can pin the
protocol constants and the SDK's encodings without reading outside the package.
`src/lib/protocol-vectors.test.ts` checks both that the SDK reproduces every value and that these
files match `tests/vectors/`.

`make vectors` regenerates the originals and refreshes this copy. Do not edit these files.
