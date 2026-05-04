# Smart Contract Gas Optimization Review

Review date: 2026-05-01

Scope: `solidity/src` with emphasis on cryptographic paths: BabyJubJub arithmetic, Schnorr verification, Groth16 verifier wrappers, BRLC transcript checks, and storage around proof submissions.

Baseline command run:

```text
cd solidity && forge test --gas-report --no-match-test '_Heavy|Stress'
```

Result: 117 tests passed. Important caveat: the `DKGManager` gas numbers in this Foundry suite use mock verifiers, so they measure manager-side storage, calldata, hashing, and transcript work, but not production Groth16 pairing/public-input MSM cost. Registry Schnorr and BabyJubJub tests do exercise the real native crypto.

## Current Hotspots

- `DKGRegistry.registerKey`: median `2,110,279` gas.
- `DKGRegistry.updateKey`: median `2,028,546` gas.
- `DKGManager.finalizeEpoch`: median `727,896` gas with mock verifier.
- `DKGManager.submitContribution`: median `238,286` gas with mock verifier.
- `DKGManager.submitPartialDecryption`: median `159,339` gas with mock verifier.
- `DKGManager.combineDecryption`: median `101,258` gas with mock verifier.
- Poseidon helper deployment is very large: `PoseidonT6` is `23,225,077` deploy gas and `115,920` bytes.
- `DKGManager` runtime size is reported as `26,937` bytes, above the EIP-170 24,576-byte mainnet limit. Treat this as a deployment risk, not only a gas concern.

## Blocking Crypto Caveat

### `BabyJubJub.isInPrimeSubgroup` is currently cheap because it is ineffective

`isInPrimeSubgroup(x, y)` calls `scalarMul(SUBGROUP_ORDER, x, y)` (`solidity/src/libraries/BabyJubJub.sol:96-98`), but `scalarMul` first reduces the scalar modulo `SUBGROUP_ORDER` (`solidity/src/libraries/BabyJubJub.sol:181`). That means `[L]P` becomes `[0]P` and returns the identity for every point.

Do not remove subgroup checks based on current gas. First either:

- add a `scalarMulRaw` path that does not reduce the scalar before `[L]P`, or
- move subgroup membership into a proof-gated statement and document where each externally supplied point is constrained.

Fixing this naively will increase `submitCiphertext` gas because `_requireValidEncryptionPoint` calls `isInPrimeSubgroup` for both ciphertext points (`solidity/src/DKGManager.sol:972-977`). The recommendations below account for that.

## Highest Impact Recommendations

### 1. Move collective-key accumulation from `submitContribution` to `finalizeEpoch`

Current path:

- Each contribution writes a compact contribution record, then updates `_collectiveKey` by calling `BabyJubJub.pointAdd` (`solidity/src/DKGManager.sol:573-588`).
- `pointAdd` uses modular inversions and is expensive.
- The contribution circuit exposes `CommitmentX0/Y0` only so the manager can update the running key (`solidity/src/DKGManager.sol:543-545`).

Optimization:

- Stop updating `_collectiveKey` in `submitContribution`.
- In `finalizeEpoch`, after `_verifyFinalizeTranscript`, read `aggregateCommitments[0]` from calldata and store it as `_collectiveKey[epochId]`.
- Remove `commitment0X` / `commitment0Y` from `submitContribution` and remove the two public inputs from the contribution circuit/verifier.

Why this is safe:

- The finalize proof already proves that aggregate commitments are the sum of the accepted contribution rows.
- `_verifyFinalizeTranscript` now rejects duplicate participant rows and checks each active row against the accepted contribution digest (`solidity/src/DKGManager.sol:785-815`).

Expected savings:

- Removes one BabyJubJub addition and two running-key storage writes per contribution.
- Removes two Groth16 public inputs from the contribution verifier, reducing production verifier MSM gas.
- Adds only two storage writes at finalize time.

### 2. Store partial decryption hashes plus bitmaps, not full delta points

Current path:

- `submitPartialDecryption` stores `participantIndex`, `ciphertextIndex`, `accepted`, `delta.x`, and `delta.y` in `epochPartialDecryptions` (`solidity/src/DKGManager.sol:893-901`).
- `combineDecryption` later reads the stored point and compares it to transcript coordinates (`solidity/src/DKGManager.sol:747-752`).

Optimization:

- Store `deltaHash = keccak256(delta.x, delta.y)` only.
- Track accepted participant indexes with a per `(epochId, aid, ciphertextIndex)` bitmap.
- During combine, read `pdX/pdY` from calldata as today and compare `keccak256(pdX, pdY)` to the stored hash.
- Key by participant index directly instead of address so combine does not need `epochParticipants[epochId][participantIndex - 1]` just to find the mapping key.

Suggested storage shape:

```solidity
mapping(bytes12 => mapping(bytes32 => mapping(uint16 => uint256))) partialBitmap;
mapping(bytes12 => mapping(bytes32 => mapping(uint16 => mapping(uint16 => bytes32)))) partialDeltaHash;
```

Expected savings:

- Saves one to two cold `SSTORE`s per partial decryption.
- Reduces combine-time storage reads and mapping hashing.
- Makes duplicate checks cheaper because bitmap existence replaces `record.accepted`.

Security notes:

- The combine transcript already carries `pdX/pdY`; hashing those coordinates back to the stored hash preserves the same binding.
- Continue to check the submitter owns `participantIndex` at submission time.

### 3. Optimize BabyJubJub `pointAdd` to use one inversion instead of two

Current `pointAdd` computes two denominators and inverts both:

- `x3` divides by `1 + d*x1*x2*y1*y2`.
- `y3` divides by `1 - d*x1*x2*y1*y2`.
- Each inversion uses the bigModExp precompile (`solidity/src/libraries/BabyJubJub.sol:141-150`, `326-341`).

Optimization:

Use one inverse:

```text
t = d*x1*x2*y1*y2
denX = 1 + t
denY = 1 - t
invProduct = inverse(denX * denY)
invDenX = denY * invProduct
invDenY = denX * invProduct
```

Then multiply `x3Num` by `invDenX` and `y3Num` by `invDenY`.

Expected savings:

- Roughly one bigModExp precompile call per BabyJubJub addition.
- This affects native Schnorr verification, scalar multiplication, subgroup checks once fixed, and `_collectiveKey` accumulation if recommendation 1 is not applied.

### 4. Replace in-SNARK Chaum-Pedersen transcript emulation with a direct DLEQ relation

Current partial-decrypt circuit proves both:

- direct relation: `PublicKey = secret*G` and `Delta = secret*C1`;
- plus a Chaum-Pedersen nonce/challenge/response transcript.

For an on-chain SNARK verifier, the direct relation is already the proof of knowledge relation. The `A1`, `A2`, `Response`, and in-circuit Fiat-Shamir hash are redundant for soundness if the SNARK statement keeps `eid`, `aid`, `ctIdx`, `role`, `participantIndex`, `C1`, `PublicKey`, and `Delta` public.

Optimization:

- Remove `Nonce`, `A1`, `A2`, `Response`, and the in-circuit challenge hash from `PartialDecryptCircuit`.
- Reduce the public input vector from 16 fields to about 11 fields.
- Simplify `submitPartialDecryption` and `submitOrganizerShare` public-input checks accordingly.

Expected savings:

- Smaller partial-decrypt circuit.
- Fewer Groth16 public inputs, so lower production verifier MSM gas for every committee partial and organizer share.
- Less calldata and less manager-side input validation.

### 5. Rewrite Groth16 verifier wrappers to avoid self-`staticcall` and duplicate decoding

Current wrappers decode `proof` and `input`, ABI-encode a typed call, then `staticcall` themselves (`solidity/src/verifiers/ContributionVerifier.sol:18-51`, `PartialDecryptVerifier.sol:18-51`, `DecryptCombineVerifier.sol:23-56`). The manager then decodes the same `input` again.

Optimization:

- Generate verifier wrappers with a direct `verifyProof(bytes calldata proof, bytes calldata input)` entry that reads calldata words directly.
- Or generate base verifier functions that accept memory arrays and can be called internally without `address(this).staticcall`.
- In `DKGManager`, read public inputs via `calldataload` from `bytes calldata input` instead of `abi.decode` where only a few fields are needed.
- Validate cheap public-input bindings before invoking the expensive verifier so bad submissions fail early.

Expected savings:

- A few thousand gas per proof-gated call from removing wrapper self-call, memory allocation, and duplicate decode.
- Larger savings on invalid/reverted transactions because malformed public inputs can fail before pairing verification.

### 6. Make transcript commitments active-length instead of `MAX_N`-length

Current BRLC checks stream fixed-size transcripts:

- contribution: `8 * MAX_N` words (`solidity/src/DKGManager.sol:54`, `566`);
- finalize: `2 * MAX_N^2 + 5 * MAX_N` words (`solidity/src/DKGManager.sol:55`, `814-815`);
- combine: `4 + 3 * MAX_N` words (`solidity/src/DKGManager.sol:56`, `1087-1089`).

This means small committees still pay for MaxN padding. `finalizeEpoch` is especially affected because it streams 2,208 words at MaxN=32.

Optimization:

- Redesign circuit transcript commitments to cover only active rows and active coefficients.
- For finalize, target approximately:
  - `participantIndexes[acceptedCount]`;
  - `contributionCommitments[acceptedCount][threshold]`;
  - `aggregateCommitments[threshold]`;
  - `shareCommitments[acceptedCount]`.
- For combine, commit only `shareCount` participant indexes and deltas.
- Keep public `threshold`, `committeeSize`, `acceptedCount`, and `shareCount` so lengths remain unambiguous.

Expected savings:

- Large `finalizeEpoch` savings when `n < MAX_N`.
- Moderate `submitContribution` and `combineDecryption` savings.
- Also reduces calldata size.

Tradeoff:

- Requires circuit changes and new proving/verifying keys.
- The contract and witness builders must agree exactly on dynamic transcript ordering.

## Medium Impact Recommendations

### 7. Compact internal contribution and partial record storage

`DKGTypes.ContributionRecord` and `DKGTypes.PartialDecryptionRecord` contain fields that are not persisted by the manager (`solidity/src/libraries/DKGTypes.sol:43-59`). The manager writes only the fields it needs, but the unused fields still force wider storage layout.

Optimization:

- Use compact internal storage structs and keep external getters ABI-compatible by constructing the old return struct in memory.
- Contribution storage can be reduced to `commitmentVectorDigest` plus packed `contributorIndex/accepted`, or replaced by index-keyed digest plus a contribution bitmap.
- Partial storage can be replaced by recommendation 2.

Expected savings:

- Fewer cold `SSTORE`s in `submitContribution` and `submitPartialDecryption`.
- Fewer `SLOAD`s in finalize/combine verification loops.

### 8. Remove unused stored hashes from organizer share records

`OrganizerShareRecord` stores `deltaOrg`, `dleqHash`, and `accepted` (`solidity/src/libraries/DKGTypes.sol:127-130`). `combineDecryption` only needs `accepted` and `deltaOrg` (`solidity/src/DKGManager.sol:1058-1060`).

Optimization:

- Do not store `dleqHash`; emit it if off-chain consumers need it.
- Consider storing only `deltaOrgHash` and passing `deltaOrgX/Y` to `combineDecryption`, mirroring the partial-delta hash approach.

Expected savings:

- Saves one cold `SSTORE` per organizer share if `dleqHash` is removed.
- Can save another slot if `deltaOrg` is stored as a hash and supplied in combine calldata.

### 9. Do not store identity organizer keys for mode-0 applications

`registerApplication` stores `organizerPK = (0, 1)` for public-derivation apps (`solidity/src/DKGManager.sol:1215-1221`). Mode 0 never needs to read that point in combine.

Optimization:

- Leave `organizerPK` zeroed for mode 0.
- Emit `(0, 1)` in the event if that representation is useful off-chain.
- In getters, optionally normalize zero organizerPK to identity for mode 0.

Expected savings:

- Saves the nonzero `organizerPK.y = 1` storage write per public-derivation app.

### 10. Use scratch-memory hashing for fixed-size tuples

Several hot paths use `abi.encode(...)` only to hash fixed-size field tuples:

- ciphertext hash (`solidity/src/DKGManager.sol:861-864`, `961`);
- share commitment hash (`solidity/src/DKGManager.sol:705`);
- stored share check (`solidity/src/DKGManager.sol:874-887`);
- organizer ciphertext binding (`solidity/src/DKGManager.sol:1132-1134`).

Optimization:

- Replace with small assembly helpers that `mstore` fixed words into scratch memory and call `keccak256(ptr, len)`.
- Use one helper for two-word point hashes and one for four-word ciphertext hashes.

Expected savings:

- Small per call, but repeated in finalize loops and proof submissions.

### 11. Prefer fail-fast public-input checks before verifier calls

Several proof-gated functions call the verifier before decoding and checking public inputs, for example `submitContribution` (`solidity/src/DKGManager.sol:529-545`), `finalizeEpoch` (`solidity/src/DKGManager.sol:639-654`), and `submitPartialDecryption` (`solidity/src/DKGManager.sol:868-888`).

Optimization:

- Check input length and cheap public-input fields directly from calldata first.
- Call the verifier only after those fields match expected epoch/application/ciphertext state.

Expected savings:

- Successful transactions improve slightly if combined with no-decode wrappers.
- Invalid transactions become much cheaper to reject.

### 12. Reconsider Poseidon for native Schnorr challenges

Registry Schnorr challenge calculation uses `PoseidonT6` then `PoseidonT3` (`solidity/src/DKGRegistry.sol:170-190`). The gas report shows:

- `PoseidonT6.hash`: `241,358` gas;
- `PoseidonT3.hash`: `58,278` gas.

Options:

- If the challenge is only verified on-chain and does not need circuit-native compatibility, switch the native Schnorr challenge to `keccak256` domain separation. This removes Poseidon runtime cost and avoids massive Poseidon helper deployments.
- If Poseidon compatibility is required, keep Poseidon but prioritize recommendation 3 because BabyJubJub point additions dominate the rest of Schnorr verification.
- If adding a registration SNARK is acceptable, replace native Schnorr verification with a small proof that `PK = sk*G`. This should be benchmarked against the native path after `pointAdd` is optimized.

Security note:

- Changing Fiat-Shamir hash changes proof vectors and domain separation. Treat this as a protocol version bump.

### 13. Remove compressed-proof support if production always sends uncompressed proofs

The contribution, partial-decrypt, and decrypt-combine wrappers support both 8-word and 4-word proofs (`solidity/src/verifiers/ContributionVerifier.sol:18-41`, `PartialDecryptVerifier.sol:18-41`, `DecryptCombineVerifier.sol:23-46`). Compressed proofs save calldata but require decompression logic in the generated verifier.

Optimization:

- Pick one proof format per deployment target.
- On L1, uncompressed proofs often win because calldata savings are smaller than decompression cost.
- Remove the unused branch and generated decompression functions.

Expected savings:

- Smaller verifier bytecode.
- Less wrapper branching.
- Avoids accidentally paying decompression gas when calldata savings do not justify it.

### 14. Fix the epoch eviction gas bomb with lazy cleanup

`_evictRound` loops over every participant, every registered aid plus legacy aid, and every ciphertext index up to 256 (`solidity/src/DKGManager.sol:454-499`). This can become too expensive when the history ring evicts a heavily used epoch.

Optimization:

- Do not eagerly delete all nested mapping entries.
- Use epoch nonce/generation checks and let old mapping entries become unreachable.
- Track only compact per-epoch roots or tombstones for recent epochs.
- If refunds are desired, expose an optional cleanup function that can be called in chunks.

Expected savings:

- Prevents `createEpoch` from becoming uncallable after a busy historical epoch.
- Reduces worst-case gas even if average gas does not change much.

## Deployment-Size Recommendations

### 15. Split `DKGManager` or move code behind external helper contracts

Current gas report shows `DKGManager` runtime size at `26,937` bytes. That is above the EIP-170 runtime code size limit.

Options:

- Move application registration / organizer co-decryption into a companion contract.
- Move eviction and view helper logic out of the manager.
- Generate smaller verifier wrappers and remove unused compressed proof paths.
- Consider a minimal core manager plus extension contracts for app-specific flows.

Tradeoff:

- External helper calls add runtime gas. Use this mainly to solve deployability and bytecode pressure, then benchmark hot paths again.

### 16. Benchmark optimizer settings after structural changes

Current `foundry.toml` uses `optimizer_runs = 200` and `via_ir = true`.

Optimization:

- Re-run gas reports with low runs (`1`, `10`) and higher runs (`500`, `1000`) after bytecode-size reductions.
- Separate deployment-size goals from hot-path runtime goals.

## Suggested Priority Order

1. Fix `isInPrimeSubgroup` semantics before relying on ciphertext subgroup validation gas.
2. Implement one-inversion `BabyJubJub.pointAdd`; rerun registry Schnorr and ciphertext validation gas tests.
3. Remove per-contribution `_collectiveKey` accumulation and store aggregate C0 only at finalize.
4. Replace partial-decryption point storage with delta hashes plus bitmaps.
5. Reduce partial-decrypt SNARK public inputs by proving the direct DLEQ relation instead of carrying Chaum-Pedersen transcript fields.
6. Regenerate verifier wrappers to avoid self-`staticcall`, duplicate decode, and unused compressed-proof branches.
7. Redesign BRLC transcript commitments to active-length transcripts.
8. Compact storage structs and remove unused stored hashes.
9. Address `DKGManager` code size by splitting non-hot-path logic or shrinking generated wrappers.

## Verification Notes

- Ran `forge test --gas-report --no-match-test '_Heavy|Stress'` from `solidity/`.
- Result: all 117 tests passed.
- The recommendations above are static-analysis suggestions; each should be benchmarked after implementation because several trade runtime gas against deployment size or circuit/verifier changes.
