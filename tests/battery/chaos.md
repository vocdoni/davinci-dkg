# Running the swarm under turbulence

`TestOrganizerSwarm` is written to survive an operator breaking things
underneath it and to *report* what happened rather than fail hard. This is the
manual companion to that test: start the swarm, disrupt the fleet from another
terminal, then read the per-ciphertext rows in the report
(`/tmp/battery-report.json` and `/tmp/battery-report.md`) to see what the fleet
did.

Every step below is optional and independent. Run one, run several, or run
none and use the numbers as a clean baseline.

## Things to try

**Restart a wave-0 node mid-run.** Run `docker restart testnet-dkg-node-<k>`
for one or two nodes while ciphertexts are being served. Partials for the
affected slots then come from the next stagger wave, three blocks later per
wave, combines still land, and the latency of those slots grows by a wave or
two. A restarted node re-scans `decrypt-lookback-blocks` on boot and must not
double-submit; an `AlreadyPartiallyDecrypted` revert in its log is benign.

**Restart the node holding combine slot 0.** The report rows of earlier
ciphertexts name it in `combiner=`. The rotation is seed-derived per ciphertext
index, so the same member tends to go first for the same index. The next node
in the rotation combines three blocks later.

**Pause the RPC.** Run `docker pause testnet-anvil-1` for 30 to 60 seconds,
then `docker unpause`. Anvil keeps its block clock, so blocks catch up quickly
on unpause. Nodes log RPC errors and back off per ciphertext, exponentially and
capped at 64 ticks. The battery's own transaction waits are generous
(`BATTERY_TX_TIMEOUT`, 3 minutes) and `waitCombine` is bounded in blocks rather
than seconds. Do not pause for longer than about two minutes: the receipt waits
expire and the affected rows are reported as failures.

**Stop a few nodes for good.** Use `docker stop` on up to `n - t` committee
members. Later waves fill in the partials and combines still land. From
`n - t + 1` stopped nodes on, slots stop combining and the rows say
`not combined within N blocks`.

**Let the epoch boundary pass.** Just start the swarm late in an epoch. The
epoch stays `Live` on chain and the nodes keep serving its ciphertexts while
they claim and contribute to the next one, so combines slow down during the
next epoch's preparation. The summary row counts
`combinedAfterEpochBoundary`.

## Why the swarm absorbs this

- Every wait is bounded in **blocks** and sized generously
  (`BATTERY_COMBINE_WAIT_BLOCKS`, default 240, which is 8 minutes at 2-second
  blocks). A slowdown shows up as a larger `latencyBlocks`, not as a failure,
  until the bound is exceeded.
- Each ciphertext is judged on its own and gets its own report row with the
  partial count, the combiner address and the plaintext check, so one stuck
  slot never masks the others.
- Withheld slots are asserted after a fixed window. A pause that delays
  partials does not change their verdict: no share means no combine.
- The summary reports throughput and average latency over the ciphertexts that
  did combine, plus the counts of the ones that did not, so even a chaotic run
  yields numbers.

## What breaks the run

- **Sharing an Anvil default key with a running node.** Both signers allocate
  nonces locally and one of them stalls. The battery itself never does this.
- **Pausing Anvil for longer than `BATTERY_TX_TIMEOUT`.**
- **Restarting the deployer.** It rewrites the addresses file.
