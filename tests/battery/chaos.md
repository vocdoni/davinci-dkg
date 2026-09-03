# Chaos notes for the swarm run

`TestOrganizerSwarm` is written to survive operator-inflicted turbulence
and *report* it rather than fail hard. While it runs, the integrator does
the following by hand and then reads the per-ciphertext rows in the report
(`/tmp/battery-report.json` / `.md`) to see what the fleet did.

## What the integrator does

1. **Restart a wave-0 node mid-run**
   `docker restart testnet-dkg-node-<k>` for one or two nodes while
   ciphertexts are being served. Expected: partials for the affected slots
   come from the next stagger wave (`staggerBlocks` = 3 blocks later per
   wave), combines still land, latency of those slots grows by a wave or
   two. The restarted node re-scans `decrypt-lookback-blocks` on boot and
   must not double-submit (`AlreadyPartiallyDecrypted` is benign).
2. **Restart the node holding the combine slot 0** for a slot — visible as
   `combiner=` in the report rows of earlier ciphertexts (the rotation is
   seed-derived per ciphertext index, so the same member tends to go first
   for the same index). Expected: the next slot in the rotation combines
   3 blocks later.
3. **Pause the RPC** — `docker pause testnet-anvil-1` for 30–60 s, then
   `docker unpause`. Anvil keeps its block clock, so on unpause blocks catch
   up quickly. Expected: nodes log RPC errors and back off per ciphertext
   (exponential, capped at 64 ticks); the battery's own tx waits are
   generous (`BATTERY_TX_TIMEOUT`, 3 min) and `waitCombine` is bounded by
   blocks, not seconds. Do not pause longer than ~2 minutes or the battery's
   receipt waits expire and the affected rows are reported as failures.
4. **Kill and never restart a few nodes** (`docker stop`, up to `n − t` of
   the committee). Expected: later waves fill the partials, combines still
   land; from `n − t + 1` stopped nodes on, slots stop combining and the
   rows say `not combined within N blocks`.
5. **Let the epoch boundary pass** — just run the swarm late in an epoch.
   The epoch stays `Live` on chain; the nodes keep serving its ciphertexts
   while claiming / contributing to the next one, so combines slow down
   during the next epoch's Preparation. The summary row counts
   `combinedAfterEpochBoundary`.

## How the swarm tolerates it

- Every wait is bounded in **blocks** and sized generously
  (`BATTERY_COMBINE_WAIT_BLOCKS`, default 240 = 8 minutes at 2 s blocks).
  A slowdown shows up as a larger `latencyBlocks`, not a failure, until
  the bound is exceeded.
- Each ciphertext is judged independently and gets its own report row with
  the partial count, the combiner address and the plaintext check, so one
  stuck slot never masks the others.
- Withheld slots are asserted after a fixed window; a pause that delays
  partials does not change their verdict (no share ⇒ no combine).
- The summary reports throughput and average latency over the ciphertexts
  that did combine, plus the counts of the ones that did not, so a chaotic
  run still yields numbers.

## What is *not* tolerated

- Sharing an Anvil default key with a running node: both signers allocate
  nonces locally and one of them will stall. The battery never does this.
- Pausing Anvil for longer than `BATTERY_TX_TIMEOUT`.
- Restarting the deployer (it would rewrite the addresses file).
