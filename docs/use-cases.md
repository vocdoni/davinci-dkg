# Use cases

What the DKG gives an application, in one sentence: a public encryption key whose secret nobody
holds, under which anyone can post small numbers on chain that stay unreadable until a
contract-enforced moment, after which a committee of `t` out of `n` operators decrypts them for
everyone. This document lists applications beyond voting that fit that shape, what each needs from
the protocol, and what each has to add on top of it.

## What the protocol provides

Every application registers against a live epoch and claims one of the epoch's 16 pool keys, so
its ciphertexts decrypt under a key that no other application shares. At registration the organizer
fixes the policy that the contract enforces afterwards:

| Knob | Options | Enforced by |
|---|---|---|
| Mode | **automatic**: any `t` committee members decrypt once the window opens; **organizer-locked**: nothing decrypts until the organizer publishes `sk_org`, once, for the whole application | `DKGAppManager`, `DKGManager` |
| Who may submit | registrant only (default), an allow-list of up to 32 addresses, or anyone | `submitCiphertext` |
| Submission window | block range and a cap on the number of ciphertexts | `submitCiphertext` |
| Decryption window | `decryptNotBefore` / `decryptNotAfter` timestamps | `submitPartialDecryption`, `combineDecryption` |

Plaintexts are scalars encrypted with exponential ElGamal on BabyJubJub. Two properties follow and
shape every use case below:

- **Additive homomorphism.** Ciphertexts of `a` and `b` can be added off chain into a ciphertext
  of `a + b` without any key. An application can therefore aggregate before it asks for a
  decryption, and reveal only the sum.
- **Small plaintexts.** Decryption ends in a discrete-logarithm search, capped at 2⁵⁰ on the
  committee and 2³² in the browser SDK. A plaintext is a counter, an amount in coarse units, an
  index or a 32-bit chunk, not a document. Anything larger is carried by a symmetric key encrypted
  as eight 32-bit chunks, eight ciphertexts.

The costs (Sepolia, 16-key pool, `n = 32`, `t = 22`; see [BENCHMARKS.md](../BENCHMARKS.md)):
registration 0.23 M gas (automatic) or 0.62 M (locked), one ciphertext 0.10 M paid by the
submitter, one decryption `t × 0.40 M + 0.48 M` paid by the committee, one reveal 0.23 M. Any
number of ciphertexts can wait in an application; only decrypted ones cost the committee anything.

### What it does not do

- It does not check who wrote a ciphertext or what it contains. `submitCiphertext` takes no proof.
  Exponential ElGamal is malleable: given someone's ciphertext of `x`, anyone admitted to submit can
  post a ciphertext of `x + 1` without learning `x`. **Any application where a participant benefits
  from copying or shifting another's value must bind ciphertexts to their authors itself**, for
  example with a proof of knowledge of the encryption randomness `r` (`C₁ = rG`, a Schnorr proof)
  or a hash commitment to `(m, r)` posted alongside, verified by the application contract before it
  forwards the ciphertext as the sole admitted submitter.
- It does not hide who submitted, when, or how many ciphertexts exist. Metadata is public.
- It does not decrypt for one party. Every decryption is public: the plaintext ends up in
  `getPlaintext` for everyone.
- Its windows bind the contract and honest nodes, not `t` colluding operators, who hold shares and
  could decrypt off chain at any time (for a locked application they would still need `sk_org`).
- A locked application whose organizer loses `sk_org` is undecryptable forever; an organizer who
  reuses one secret across two applications opens both with one reveal.

## Sealed-bid auctions

The canonical fit. Bidders encrypt their bid amount under the auction's key during the bidding
window; the ciphertexts sit on chain, unreadable, so nobody can react to a rival's bid; when the
window closes the committee decrypts every bid and the auction contract picks the winner.

**Setup.** One application per auction, automatic mode, `decryptNotBefore` = end of bidding,
submitter = the auction contract only (registrant policy). Bidders call the auction contract, which
verifies a proof of knowledge of `r` for the ciphertext (or a commitment to `(bid, r)`), records
`msg.sender` against the ciphertext index and forwards the ciphertext to the DKG. Bid amounts are
in coarse units (cents, or a fixed-point tick) so they stay below 2³². A deposit or the escrowed
payment prevents bids the bidder cannot honour.

**Variants.**

- *First-price*: highest decrypted bid wins and pays it.
- *Vickrey (second-price)*: same data, winner pays the second-highest bid; the decrypted bids are
  exactly what the rule needs.
- *Sealed reserve price*: the seller submits an encrypted reserve as the first ciphertext; it is
  revealed together with the bids, so bidders cannot probe it.
- *Procurement tenders*: reverse auction, lowest compliant bid wins. Organizer-locked mode fits
  public procurement, where the tender office opens all bids at a formal session: the reveal is
  that session, and until it happens the contract admits no partial, so no bid can leak early even
  to the committee's honest nodes.
- *Domain names, NFTs, spectrum-style allocations*: a batch of items becomes a batch of auctions,
  one application each, up to 16 per epoch.

**Why the DKG rather than commit-reveal.** Commit-reveal lets a bidder who sees they lost refuse to
reveal, and a bidder who sees they won overpay nothing. Here every bid is opened by the committee
whether or not its author cooperates, and nothing about a bid is known, even to the auctioneer,
until the window closes. Losing bids do become public after the auction; if that matters, the
application can encrypt bids under a per-auction key and only submit the winning-candidate
ciphertexts, which needs an order-preserving step this protocol does not provide.

## Batch auctions and encrypted order flow

Decentralized exchanges that clear in discrete batches (frequent batch auctions) suffer from orders
being visible before the batch clears. Encrypting each order's limit price and size under a
per-batch key removes the information a front-runner needs: orders are placed during the batch,
decrypted by the committee at the batch boundary, then matched by the exchange contract at one
uniform clearing price.

**Setup.** One application per batch (a 24-hour epoch pays for 16 batches; shorter batches need
more epochs or one application per batch with a rolling epoch), automatic mode, `decryptNotBefore`
= batch end, open submission or the exchange contract as the single submitter. An order is two
small scalars, price tick and size, so two ciphertexts, or one packed scalar when both fit in
2⁵⁰ together. The same malleability caveat as for auctions applies: the exchange contract must bind
orders to their senders and their collateral.

**What it buys.** Front-running and sandwiching within a batch become impossible; ordering inside
the batch is irrelevant because the batch clears at one price. What it does not buy: privacy after
the batch, since every order is decrypted, and protection against a batch boundary an adversary can
move.

## Scheduled disclosure and time-lock encryption

Anything that must be published at a fixed future time, by nobody in particular: embargoed
announcements, exam papers, results that must appear simultaneously to all parties, sealed
predictions whose authors want to prove they were made before the event, dead-drop keys.

**Setup.** Automatic mode, `decryptNotBefore` = the disclosure time, `decryptNotAfter` unbounded
or a grace period, registrant-only submission. A payload larger than a scalar is encrypted with a
one-time symmetric key and stored anywhere (IPFS, the application's own contract); the key's eight
32-bit chunks are the eight ciphertexts. At the disclosure time the committee decrypts the chunks
without anybody asking, and the payload becomes readable by everyone.

**Trust.** The disclosure happens iff at least `t` honest committee members are online after the
time and fewer than `t` collude to open it early. Compared with a single escrow agent this replaces
one trusted party by a `t`-of-`n` committee whose members are registered and paid on chain;
compared with a puzzle-based time-lock it needs no assumption about anyone's hardware speed, and it
cannot be opened early by someone with more compute. What cannot be done: cancelling a disclosure
once posted. An author who may want to retract should encrypt under a locked application and hold
`sk_org` as the switch, accepting that then the disclosure depends on them.

## Committed randomness and lotteries

Drawing a random number among untrusting parties is usually commit-reveal, whose known failure is
the last participant who declines to reveal after seeing the others. With the DKG each participant
submits an encrypted contribution during the commit window; after the window the committee decrypts
all of them and the draw is a public function of the sum, for example the sum modulo the number
of tickets.

**Setup.** Automatic mode, `decryptNotBefore` = end of the commit window, open submission or the
lottery contract as submitter, one ciphertext per participant. Contributions must be bound to
their authors (a shifted copy of another's contribution is as good as a fresh one, so binding is
about accountability, not about biasing the sum). The result is unbiasable by any coalition smaller
than `t` committee members, and no participant can abort it.

**Uses.** Raffles and ticket draws, leader or sequencer election among a known set, seeds for
on-chain games, assignment of scarce slots (allow-list spots, validator positions) by verifiable
lottery, and, more prosaically, a public randomness beacon whose rounds are the epochs.

## Simultaneous-move games

Games where players act at the same time and must not see each other's move first: rock–paper–
scissors and its tournament versions, sealed orders in Diplomacy-style strategy games, blind
guesses, simultaneous price-setting in economic experiments.

**Setup.** One application per game or per round, automatic mode, `decryptNotBefore` = the round
deadline, players as the allow-list. A move is a small integer, one ciphertext. After the deadline
the committee opens all moves; a player who never submitted has a public default. The same
structure gives *sealed answers* for quizzes and prediction tournaments: answers are posted before
the question closes and revealed together.

## Private aggregation of numeric inputs

The homomorphism turns the DKG into a threshold aggregator: participants encrypt a number each, the
application adds the ciphertexts off chain, and only the sum is ever submitted for decryption. No
individual value is decrypted at all.

**Setup.** Registrant-only submission (so nobody can submit an individual ciphertext for
decryption), automatic mode, the aggregator as registrant. Participants encrypt under the
application key with the SDK and hand their ciphertexts to the aggregator through any channel, or
post them to the application's own contract, which the aggregator reads. The aggregator publishes
the sum ciphertext, the committee decrypts it, and anyone can check that the published sum
ciphertext is the homomorphic sum of the posted individual ones.

**Uses.**

- *Compensation and benchmark surveys*: each firm reports a salary band or a cost figure; only the
  total (and, with a second application, the count) is revealed.
- *Grant and peer review*: reviewers score proposals; per-proposal sums are decrypted, individual
  scores are not. A committee member cannot see who scored what.
- *Pledges and conditional commitments*: participants encrypt what they would contribute if a
  threshold is met; the total is decrypted at the deadline; individual pledges stay private unless
  the campaign succeeds and a second application reveals them.
- *Participation counts and attestations*: each eligible party encrypts 0 or 1; the sum is the
  count of attesters without a list.

**Limits.** The sum must stay below 2⁵⁰, which bounds the number of participants times the value
range. Range enforcement on individual inputs (that each is 0 or 1, or below a cap) is not provided
by the DKG; an application that needs it adds a range proof at submission. Encrypted inputs that
are not individually submitted also need the aggregator to be honest about which inputs it added;
posting the individual ciphertexts on chain makes that check public.

## Sealed applications and allocations

Fair launches, allow-list sales, token allocations and any "state your demand, then we allocate"
process: participants encrypt the quantity they want (and, for a price-discovery sale, the price
they would pay), the committee decrypts after the subscription window, and the allocation rule runs
on public numbers that nobody could adjust to others' demand. Organizer-locked mode fits an issuer
who wants to open the book at a formal moment; automatic mode removes the issuer from the loop.

## Escrowed unlock codes

A 32-bit code (a locker PIN, a shared-mobility unlock, a license key fragment) encrypted under a
locked application: the seller is the organizer and reveals `sk_org` when payment settles, at which
point the committee decrypts the code for the buyer, publicly. Since the reveal is application-wide
and public, this is a one-shot pattern for one buyer per application, cheap enough at 0.23 M gas
for the reveal plus one decryption.

## Choosing the mode

| You want | Mode | Reveal |
|---|---|---|
| Disclosure at a time, with no party in the loop | automatic | none |
| Disclosure when an authority decides, never before | organizer-locked | the authority, once |
| Only the sum of many inputs, never the inputs | automatic, registrant-only submission | none |
| Retractable disclosure | organizer-locked, organizer holds the switch | the organizer, or never |

## Integration checklist

1. Register the application with `dkgapp register` or the SDK (`registerApplication`), choosing
   mode, submitter policy, ciphertext cap and windows.
2. If participants may gain by copying or shifting each other's inputs, put an application contract
   in front as the sole submitter and have it verify a proof of knowledge of `r` (or a commitment)
   before forwarding.
3. Encode plaintexts as scalars below 2³² (browser) or 2⁵⁰ (committee); split anything larger into
   32-bit chunks.
4. Encrypt with the SDK (`encryptForApplication`) or `dkgapp encrypt`; read results with
   `waitForDecryption` / `getPlaintext` or `dkgapp plaintext`.
5. For a locked application, store `sk_org` where the reveal will be performed from; losing it
   loses the application.
6. Budget: one epoch serves 16 applications; ciphertexts cost the submitter 0.10 M gas each; the
   committee pays for decryptions, so an application that decrypts thousands of ciphertexts should
   aggregate first.
