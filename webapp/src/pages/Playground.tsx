import React, { useState, useEffect, useRef, useCallback } from 'react';
import {
  Alert,
  AlertIcon,
  Badge,
  Box,
  Button,
  Code,
  Divider,
  FormControl,
  FormLabel,
  FormHelperText,
  Grid,
  GridItem,
  HStack,
  Heading,
  Input,
  Spinner,
  Switch,
  Text,
  VStack,
} from '@chakra-ui/react';
import { DKGWriter, RoundStatus, buildElGamal, buildRoundId } from '@vocdoni/davinci-dkg-sdk';
import type { Round, BabyJubPoint, ElGamalCiphertext, RoundEvent } from '@vocdoni/davinci-dkg-sdk';
import type { WalletClient, Address } from 'viem';
import { getDKGClient, getClient, loadConfig } from '../lib/client';
import { blocksRemaining, formatBlocksRemaining } from '../lib/format';
import { connectWallet, hasWallet } from '../lib/wallet';

// ── Types ─────────────────────────────────────────────────────────────────────

type LogLevel = 'info' | 'success' | 'error' | 'tx' | 'chain' | 'crypto';
interface LogEntry { ts: number; level: LogLevel; msg: string }
type StepStatus = 'pending' | 'active' | 'done' | 'error';

// ── Tiny helpers ──────────────────────────────────────────────────────────────

function hex(n: bigint, pad = 64) {
  return '0x' + n.toString(16).padStart(pad, '0');
}
function short(h: string, pre = 10, suf = 8) {
  if (!h || h.length <= pre + suf + 3) return h;
  return `${h.slice(0, pre)}…${h.slice(-suf)}`;
}

// ── UI primitives ─────────────────────────────────────────────────────────────

function StepCard({
  n, title, status, children,
}: { n: number; title: string; status: StepStatus; children: React.ReactNode }) {
  const border = status === 'active' ? 'blue.400'
    : status === 'done' ? 'green.600'
    : status === 'error' ? 'red.600'
    : 'gray.700';
  const bg = status === 'active' ? 'blue.950'
    : status === 'done' ? 'green.950'
    : status === 'error' ? 'red.950'
    : 'gray.800';

  return (
    <Box
      bg={bg}
      border="1px solid" borderColor={border} borderRadius="md" p={5}
      opacity={status === 'pending' ? 0.55 : 1}
      transition="opacity 0.2s, border-color 0.2s"
    >
      <HStack mb={4} spacing={3}>
        <Box
          w={7} h={7} borderRadius="full" flexShrink={0}
          bg={status === 'done' ? 'green.500' : status === 'error' ? 'red.500'
            : status === 'active' ? 'blue.500' : 'gray.600'}
          display="flex" alignItems="center" justifyContent="center"
          fontSize="xs" fontWeight="bold" color="white"
        >
          {status === 'done' ? '✓' : status === 'error' ? '✗' : n}
        </Box>
        <Heading size="sm" color={status === 'pending' ? 'gray.500' : 'white'}>
          {title}
        </Heading>
        {status === 'active' && <Spinner size="sm" color="blue.300" ml="auto" />}
      </HStack>
      {children}
    </Box>
  );
}

function KV({ label, value, mono = true }: { label: string; value: React.ReactNode; mono?: boolean }) {
  return (
    <HStack spacing={3} align="start" wrap="wrap">
      <Text fontSize="xs" color="gray.400" w="180px" flexShrink={0} pt="1px">{label}</Text>
      <Box fontSize="xs" fontFamily={mono ? 'mono' : undefined} color="gray.200" wordBreak="break-all">
        {value}
      </Box>
    </HStack>
  );
}

function MonoBox({ children }: { children: React.ReactNode }) {
  return (
    <Box bg="gray.900" p={3} borderRadius="sm" fontFamily="mono" fontSize="xs"
      border="1px solid" borderColor="gray.700" wordBreak="break-all">
      {children}
    </Box>
  );
}

function TerminalLog({ entries }: { entries: LogEntry[] }) {
  const ref = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (ref.current) ref.current.scrollTop = ref.current.scrollHeight;
  }, [entries]);

  const color: Record<LogLevel, string> = {
    info: 'gray.300', success: 'green.300', error: 'red.400',
    tx: 'yellow.300', chain: 'cyan.300', crypto: 'purple.300',
  };

  return (
    <Box ref={ref} bg="black" p={4} borderRadius="md" fontFamily="mono" fontSize="xs"
      maxH="280px" overflowY="auto" border="1px solid" borderColor="gray.700">
      {entries.length === 0 && <Text color="gray.600">Activity log will appear here…</Text>}
      {entries.map((e, i) => (
        <Box key={i} color={color[e.level]} mb="1px">
          <Text as="span" color="gray.600" userSelect="none">
            [{new Date(e.ts).toLocaleTimeString('en', { hour12: false })}]{' '}
          </Text>
          {e.msg}
        </Box>
      ))}
    </Box>
  );
}

// ── Main component ────────────────────────────────────────────────────────────

export function Playground() {
  // ── Wallet
  const [address, setAddress] = useState<Address | null>(null);
  const [walletClient, setWalletClient] = useState<WalletClient | null>(null);
  const [walletError, setWalletError] = useState('');
  const [walletBusy, setWalletBusy] = useState(false);

  // ── Round creation form
  const [form, setForm] = useState({
    threshold: '2', committeeSize: '3', minValidContributions: '2',
    lotteryAlphaBps: '15000', seedDelay: '2',
    regDeadlineOffset: '10', contribDeadlineOffset: '20',
    disclosureAllowed: false,
  });

  // ── Round state
  const [writer, setWriter] = useState<DKGWriter | null>(null);
  const [createBusy, setCreateBusy] = useState(false);
  const [createError, setCreateError] = useState('');
  const [txHash, setTxHash] = useState('');
  const [roundId, setRoundId] = useState<`0x${string}` | null>(null);
  const [round, setRound] = useState<Round | null>(null);
  const [currentBlock, setCurrentBlock] = useState<bigint | null>(null);
  const [participants, setParticipants] = useState<Address[]>([]);
  const [events, setEvents] = useState<RoundEvent[]>([]);
  const [abortBusy, setAbortBusy] = useState(false);
  const [finalizedInfo, setFinalizedInfo] = useState<{
    collectivePublicKeyHash: string;
    aggregateCommitmentsHash: string;
    shareCommitmentHash: string;
    blockNumber: bigint;
  } | null>(null);

  // ── Crypto state
  const [collectivePubKey, setCollectivePubKey] = useState<{ x: bigint; y: bigint } | null>(null);
  const [collectivePubKeyBusy, setCollectivePubKeyBusy] = useState(false);
  const [plaintext, setPlaintext] = useState('42');
  const [ciphertext, setCiphertext] = useState<ElGamalCiphertext | null>(null);
  const [encryptDetails, setEncryptDetails] = useState<{
    k: bigint; c1: BabyJubPoint; s: BabyJubPoint; mPoint: BabyJubPoint;
  } | null>(null);
  const [submittedCiphertext, setSubmittedCiphertext] = useState(false);
  const [submitBusy, setSubmitBusy] = useState(false);
  const [combinedRecord, setCombinedRecord] = useState<{
    completed: boolean; plaintextHash: string; ciphertextIndex: number;
  } | null>(null);

  // ── Activity log
  const [log, setLog] = useState<LogEntry[]>([]);

  const addLog = useCallback((msg: string, level: LogLevel = 'info') => {
    setLog(prev => [...prev, { ts: Date.now(), level, msg }]);
  }, []);

  // ── Polling ref
  const pollRef = useRef<ReturnType<typeof setInterval> | null>(null);

  // ── Step statuses
  const stepWallet: StepStatus = address ? 'done' : walletError ? 'error' : walletBusy ? 'active' : 'active';
  const stepCreate: StepStatus = !address ? 'pending' : roundId ? 'done' : createError ? 'error' : createBusy ? 'active' : 'active';
  const stepWatch: StepStatus = !roundId ? 'pending' : (finalizedInfo || collectivePubKey) ? 'done' : round?.status === RoundStatus.Aborted ? 'error' : 'active';
  const stepKey: StepStatus = !roundId ? 'pending' : collectivePubKey ? 'done' : collectivePubKeyBusy ? 'active' : (finalizedInfo || (round && Number(round.contributionCount) >= Number(round.policy.minValidContributions))) ? 'active' : 'pending';
  const stepEncrypt: StepStatus = !collectivePubKey ? 'pending' : ciphertext ? 'done' : 'active';
  const stepSubmit: StepStatus = !ciphertext ? 'pending' : combinedRecord?.completed ? 'done' : submittedCiphertext ? 'active' : 'active';
  const stepVerify: StepStatus = !combinedRecord?.completed ? 'pending' : 'done';

  // ── Connect wallet
  async function handleConnect() {
    setWalletBusy(true);
    setWalletError('');
    try {
      addLog('Requesting wallet connection…', 'chain');
      const { address: addr, walletClient: wc } = await connectWallet();
      setAddress(addr);
      setWalletClient(wc);

      const cfg = await loadConfig();
      const publicClient = await getClient();
      const w = new DKGWriter({
        publicClient,
        walletClient: wc,
        managerAddress: cfg.managerAddress,
        ...(cfg.registryAddress ? { registryAddress: cfg.registryAddress } : {}),
      });
      setWriter(w);

      addLog(`Connected: ${addr}`, 'success');
      addLog(`Chain: ${cfg.chainName} (id=${cfg.chainId})  RPC: ${cfg.rpcUrl}`, 'chain');
      addLog(`Manager:  ${cfg.managerAddress}`, 'chain');
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : String(err);
      setWalletError(msg);
      addLog(`Wallet error: ${msg}`, 'error');
    } finally {
      setWalletBusy(false);
    }
  }

  // ── Create round
  async function handleCreateRound() {
    if (!writer) return;
    setCreateBusy(true);
    setCreateError('');
    try {
      const currentBlock = await writer.blockNumber();
      const regOffset = BigInt(form.regDeadlineOffset || '10');
      const contribOffset = BigInt(form.contribDeadlineOffset || '20');

      const policy = {
        threshold: Number(form.threshold),
        committeeSize: Number(form.committeeSize),
        minValidContributions: Number(form.minValidContributions),
        lotteryAlphaBps: Number(form.lotteryAlphaBps),
        seedDelay: Number(form.seedDelay),
        registrationDeadlineBlock: currentBlock + regOffset,
        contributionDeadlineBlock: currentBlock + contribOffset,
        disclosureAllowed: form.disclosureAllowed,
      };

      addLog('─── Creating round ───', 'info');
      addLog(`Current block: #${currentBlock}`, 'info');
      addLog(`Policy: threshold=${policy.threshold}  committee=${policy.committeeSize}  `
        + `minContribs=${policy.minValidContributions}  alphaBps=${policy.lotteryAlphaBps}`, 'info');
      addLog(`Deadlines: reg=#${policy.registrationDeadlineBlock}  contrib=#${policy.contributionDeadlineBlock}`, 'info');
      addLog(`Seed delay: ${policy.seedDelay} blocks  disclosure: ${policy.disclosureAllowed}`, 'info');
      addLog('Sending createRound transaction…', 'tx');

      const hash = await writer.createRound(policy);
      setTxHash(hash);
      addLog(`TX submitted: ${hash}`, 'tx');

      const receipt = await writer.waitForTransaction(hash);
      addLog(`TX mined: block #${receipt.blockNumber}  gas=${receipt.gasUsed}`, 'tx');

      const nonce = await writer.roundNonce();
      const prefix = await writer.roundPrefix();
      const id = buildRoundId(prefix, nonce);
      setRoundId(id);

      addLog(`Round nonce: ${nonce}`, 'info');
      addLog(`Round ID:   ${id}`, 'success');
      addLog('Watching for round status changes…', 'chain');
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : String(err);
      setCreateError(msg);
      addLog(`createRound failed: ${msg}`, 'error');
    } finally {
      setCreateBusy(false);
    }
  }

  // ── Poll round state
  useEffect(() => {
    if (!roundId || !writer || finalizedInfo || round?.status === RoundStatus.Aborted) return;

    let lastStatus = -1;
    let lastClaimed = -1;
    let lastContribs = -1;
    let lastEventCount = 0;
    let earlyKeyFetched = false;  // guard: only attempt early extraction once

    async function poll() {
      try {
        const [r, parts, blk] = await Promise.all([
          writer!.getRound(roundId!),
          writer!.selectedParticipants(roundId!),
          writer!.blockNumber(),
        ]);
        setCurrentBlock(blk);
        setRound(r);
        setParticipants(parts);

        if (Number(r.status) !== lastStatus) {
          lastStatus = Number(r.status);
          const labels = ['None', 'Registration', 'Contribution', 'Finalized', 'Aborted', 'Completed'];
          addLog(`Round status → ${labels[lastStatus] ?? lastStatus}`, lastStatus === 3 ? 'success' : lastStatus === 4 ? 'error' : 'chain');
        }
        if (parts.length !== lastClaimed) {
          lastClaimed = parts.length;
          if (parts.length > 0) addLog(`Participants: ${parts.map(p => short(p)).join(', ')}`, 'info');
        }
        if (Number(r.contributionCount) !== lastContribs) {
          lastContribs = Number(r.contributionCount);
          if (lastContribs > 0) addLog(`Contributions submitted: ${lastContribs}/${r.policy.minValidContributions}`, 'info');
        }

        // ── Early collective public key extraction ──────────────────────────
        // Once we have at least minValidContributions, compute the collective
        // public key directly from contribution calldata without waiting for
        // finalizeRound.  This is mathematically identical to the verified key.
        if (
          !earlyKeyFetched &&
          !collectivePubKey &&
          r.status === RoundStatus.Contribution &&
          Number(r.contributionCount) >= Number(r.policy.minValidContributions) &&
          parts.length > 0
        ) {
          earlyKeyFetched = true;
          setCollectivePubKeyBusy(true);
          try {
            addLog('─── Enough contributions — deriving collective public key ───', 'crypto');
            const pk = await writer!.getCollectivePublicKeyFromContributions(roundId!, parts);
            setCollectivePubKey(pk);
            addLog('Collective public key derived from contribution calldata:', 'crypto');
            addLog(`  x: ${hex(pk.x)}`, 'crypto');
            addLog(`  y: ${hex(pk.y)}`, 'crypto');
            addLog('Steps 4–6 are now unlocked.  The key will be verified on-chain once finalizeRound is mined.', 'success');
          } catch (err) {
            earlyKeyFetched = false; // allow retry on next tick
            addLog(`Early key extraction failed (will retry): ${err instanceof Error ? err.message : String(err)}`, 'error');
          } finally {
            setCollectivePubKeyBusy(false);
          }
        }

        // Fetch events
        const evs = await writer!.getAllRoundEvents(roundId!);
        if (evs.length !== lastEventCount) {
          for (let i = lastEventCount; i < evs.length; i++) {
            addLog(`Event: ${evs[i].eventName}  block=#${evs[i].blockNumber}  tx=${short(evs[i].transactionHash)}`, 'chain');
          }
          lastEventCount = evs.length;
          setEvents(evs);
        }

        if (r.status === RoundStatus.Finalized || r.status === RoundStatus.Completed) {
          const finEvents = await writer!.getRoundFinalizedEvents(roundId!);
          if (finEvents.length > 0) {
            const ev = finEvents[finEvents.length - 1];
            setFinalizedInfo({
              collectivePublicKeyHash: ev.collectivePublicKeyHash,
              aggregateCommitmentsHash: ev.aggregateCommitmentsHash,
              shareCommitmentHash: ev.shareCommitmentHash,
              blockNumber: ev.blockNumber,
            });
            addLog('─── Round finalized on-chain ───', 'success');
            addLog(`collectivePublicKeyHash: ${ev.collectivePublicKeyHash}`, 'success');
            addLog(`aggregateCommitmentsHash: ${ev.aggregateCommitmentsHash}`, 'info');
            addLog(`shareCommitmentHash:      ${ev.shareCommitmentHash}`, 'info');

            // If the early key wasn't extracted yet, fall back to finalizeRound calldata.
            if (!collectivePubKey) {
              setCollectivePubKeyBusy(true);
              try {
                const pk = await writer!.getCollectivePublicKey(roundId!);
                setCollectivePubKey(pk);
                addLog('Collective public key extracted from finalizeRound calldata:', 'crypto');
                addLog(`  x: ${hex(pk.x)}`, 'crypto');
                addLog(`  y: ${hex(pk.y)}`, 'crypto');
              } catch (err) {
                addLog(`Failed to extract collective pubkey: ${err instanceof Error ? err.message : String(err)}`, 'error');
              } finally {
                setCollectivePubKeyBusy(false);
              }
            }
            // Stop polling only after we've successfully processed the finalized event.
            clearInterval(pollRef.current!);
            pollRef.current = null;
          }
          // If finEvents is empty, don't stop — keep polling until the event is indexed.
        }

        if (r.status === RoundStatus.Aborted) {
          addLog('Round was aborted.', 'error');
          clearInterval(pollRef.current!);
          pollRef.current = null;
        }
      } catch { /* ignore transient RPC errors */ }
    }

    poll(); // immediate first tick
    pollRef.current = setInterval(poll, 2000);
    return () => { if (pollRef.current) clearInterval(pollRef.current); };
  }, [roundId, writer, finalizedInfo, addLog]); // eslint-disable-line

  // ── Poll for combined decryption once ciphertext is submitted
  const combinePollRef = useRef<ReturnType<typeof setInterval> | null>(null);
  useEffect(() => {
    if (!submittedCiphertext || !roundId || !writer || combinedRecord?.completed) return;

    let lastPartials = -1;
    async function pollCombined() {
      try {
        const [r, rec] = await Promise.all([
          writer!.getRound(roundId!),
          writer!.getCombinedDecryption(roundId!, 1),
        ]);
        const partials = Number(r.partialDecryptionCount);
        if (partials !== lastPartials) {
          lastPartials = partials;
          addLog(`Partial decryptions on-chain: ${partials} / ${r.policy.threshold}`, 'chain');
        }
        if (rec.completed) {
          setCombinedRecord({ completed: true, plaintextHash: rec.plaintextHash, ciphertextIndex: rec.ciphertextIndex });
          addLog('─── Decryption combined on-chain ───', 'success');
          addLog(`plaintextHash: ${rec.plaintextHash}`, 'success');
          clearInterval(combinePollRef.current!);
          combinePollRef.current = null;
        }
      } catch { /* ignore transient RPC errors */ }
    }

    pollCombined();
    combinePollRef.current = setInterval(pollCombined, 3000);
    return () => { if (combinePollRef.current) clearInterval(combinePollRef.current); };
  }, [submittedCiphertext, roundId, writer, combinedRecord, addLog]); // eslint-disable-line

  // ── Abort round (organizer only)
  async function handleAbort() {
    if (!writer || !roundId) return;
    setAbortBusy(true);
    try {
      addLog('─── Aborting round ───', 'error');
      const hash = await writer.abortRound(roundId);
      addLog(`Abort TX submitted: ${hash}`, 'tx');
      const receipt = await writer.waitForTransaction(hash);
      addLog(`Abort TX mined: block #${receipt.blockNumber}`, 'tx');
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : String(err);
      addLog(`Abort failed: ${msg}`, 'error');
    } finally {
      setAbortBusy(false);
    }
  }

  // ── Encrypt with the DKG collective public key
  async function handleEncrypt() {
    if (!collectivePubKey) return;
    const msg = BigInt(plaintext || '0');
    addLog(`─── Encrypting plaintext m=${msg} with DKG collective public key ───`, 'crypto');

    const pubKey: BabyJubPoint = [collectivePubKey.x, collectivePubKey.y];
    const eg = await buildElGamal();
    const k = eg.randomScalar();
    addLog(`Ephemeral k: ${hex(k)}`, 'crypto');

    // BabyJubJub generator = Base8
    const G: BabyJubPoint = [
      5299619240641551281634865583518297030282874472190772894086521144482721001553n,
      16950150798460657717958625567821834550301663161624707787222815936182638968203n,
    ];
    const c1 = eg.mulPoint(G, k);
    const s = eg.mulPoint(pubKey, k);
    const mPoint = eg.mulPoint(G, msg);

    addLog(`c1 = k × G:     (${hex(c1[0])}, ${hex(c1[1])})`, 'crypto');
    addLog(`s  = k × Q:     (${hex(s[0])}, ${hex(s[1])})`, 'crypto');
    addLog(`m × G:          (${hex(mPoint[0])}, ${hex(mPoint[1])})`, 'crypto');

    const ct = eg.encrypt(msg, pubKey, k);
    addLog(`c2 = m×G + s:   (${hex(ct.c2[0])}, ${hex(ct.c2[1])})`, 'crypto');
    addLog(`Ciphertext ready.`, 'success');

    setEncryptDetails({ k, c1, s, mPoint });
    setCiphertext(ct);
    setSubmittedCiphertext(false);
    setCombinedRecord(null);
  }

  // ── Submit ciphertext to node for threshold decryption
  async function handleSubmitCiphertext() {
    if (!ciphertext || !roundId) return;
    setSubmitBusy(true);
    try {
      addLog('Submitting ciphertext to node…', 'chain');
      const roundHex = roundId.slice(2); // strip '0x'
      const body = JSON.stringify({
        ciphertext_index: 1,
        c1x: ciphertext.c1[0].toString(10),
        c1y: ciphertext.c1[1].toString(10),
        c2x: ciphertext.c2[0].toString(10),
        c2y: ciphertext.c2[1].toString(10),
      });
      const res = await fetch(`/api/ciphertext/${roundHex}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body,
      });
      if (!res.ok) {
        const text = await res.text();
        throw new Error(`API error ${res.status}: ${text}`);
      }
      addLog('Ciphertext stored — nodes will compute partial decryptions.', 'success');
      addLog('Polling for combined decryption…', 'chain');
      setSubmittedCiphertext(true);
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : String(err);
      addLog(`Submit failed: ${msg}`, 'error');
    } finally {
      setSubmitBusy(false);
    }
  }

  // ── Render
  return (
    <VStack align="stretch" spacing={5}>
      <Box>
        <Heading size="lg" mb={1}>DKG Playground</Heading>
        <Text color="gray.400" fontSize="sm">
          Interactive developer interface for the full DKG flow — create a round,
          watch nodes participate, then exercise ElGamal encryption/decryption using
          the BabyJubJub curve. Every cryptographic step is shown in full detail.
        </Text>
      </Box>

      {/* ── Step 1: Wallet ────────────────────────────────────────────────── */}
      <StepCard n={1} title="Connect Wallet" status={stepWallet}>
        {!address ? (
          <VStack align="start" spacing={3}>
            {!hasWallet() && (
              <Alert status="warning" variant="left-accent" borderRadius="sm" fontSize="sm">
                <AlertIcon />
                No browser wallet detected. Install MetaMask or a compatible EIP-1193 wallet.
              </Alert>
            )}
            <Text fontSize="sm" color="gray.400">
              The playground requires a browser wallet to sign transactions. The correct
              chain will be auto-added/switched based on <Code fontSize="xs">/config.json</Code>.
            </Text>
            <Button colorScheme="blue" size="sm" isLoading={walletBusy}
              isDisabled={!hasWallet()} onClick={handleConnect}>
              Connect Wallet
            </Button>
            {walletError && <Text color="red.400" fontSize="xs" fontFamily="mono">{walletError}</Text>}
          </VStack>
        ) : (
          <VStack align="start" spacing={2}>
            <KV label="Connected address" value={address} />
            <KV label="Wallet client" value={walletClient ? 'viem WalletClient (EIP-1193)' : '–'} />
            <KV label="Manager contract" value={writer?.managerAddress ?? '–'} />
            <KV label="Registry contract" value={writer ? (() => { try { return writer.registryAddress; } catch { return '(auto-derive on first use)'; } })() : '–'} />
          </VStack>
        )}
      </StepCard>

      {/* ── Step 2: Create Round ─────────────────────────────────────────── */}
      <StepCard n={2} title="Create DKG Round" status={stepCreate}>
        {!roundId ? (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              Configure the round policy. The <Code fontSize="xs">committeeSize</Code> must be ≤
              the number of registered active nodes. <Code fontSize="xs">lotteryAlphaBps</Code> is
              the over-subscription factor in basis points (10000 = 1×, 15000 = 1.5×, min 10000).
              <Code fontSize="xs"> seedDelay</Code> blocks must pass before slot claiming opens.
            </Text>
            <Grid templateColumns="repeat(2, 1fr)" gap={3}>
              {[
                ['Threshold (t)', 'threshold', 'Minimum shares needed to reconstruct the secret'],
                ['Committee size', 'committeeSize', 'Total selected participants'],
                ['Min valid contributions', 'minValidContributions', 'Minimum accepted contributions required to finalize'],
                ['Lottery α (bps)', 'lotteryAlphaBps', 'Over-subscription factor; ≥10000, e.g. 15000 = 1.5×'],
                ['Seed delay (blocks)', 'seedDelay', 'Blocks to wait after creation before claimSlot opens (1–256)'],
                ['Reg. deadline (+blocks)', 'regDeadlineOffset', 'Blocks from now until registration closes'],
                ['Contrib. deadline (+blocks)', 'contribDeadlineOffset', 'Blocks from now until contribution phase closes'],
              ].map(([label, key, help]) => (
                <GridItem key={key}>
                  <FormControl>
                    <FormLabel fontSize="xs" color="gray.300" mb={1}>{label}</FormLabel>
                    <Input
                      size="sm" fontFamily="mono" fontSize="sm"
                      value={form[key as keyof typeof form] as string}
                      onChange={e => setForm(f => ({ ...f, [key]: e.target.value }))}
                      isDisabled={!!roundId || createBusy}
                    />
                    <FormHelperText fontSize="xs" color="gray.500">{help}</FormHelperText>
                  </FormControl>
                </GridItem>
              ))}
              <GridItem>
                <FormControl display="flex" alignItems="center" mt={4}>
                  <FormLabel fontSize="xs" color="gray.300" mb={0} mr={2}>Disclosure allowed</FormLabel>
                  <Switch
                    size="sm"
                    isChecked={form.disclosureAllowed}
                    onChange={e => setForm(f => ({ ...f, disclosureAllowed: e.target.checked }))}
                    isDisabled={!!roundId || createBusy}
                  />
                  <FormHelperText ml={3} fontSize="xs" color="gray.500" mb={0}>
                    Enables share disclosure phase after finalization
                  </FormHelperText>
                </FormControl>
              </GridItem>
            </Grid>
            {createError && (
              <Alert status="error" variant="left-accent" borderRadius="sm" fontSize="sm">
                <AlertIcon />{createError}
              </Alert>
            )}
            <Button colorScheme="cyan" size="sm" isLoading={createBusy}
              isDisabled={!address || !!roundId} onClick={handleCreateRound}>
              Create Round →
            </Button>
          </VStack>
        ) : (
          <VStack align="start" spacing={2}>
            <KV label="Transaction hash" value={txHash} />
            <KV label="Round ID (bytes12)" value={roundId} />
            <KV label="Threshold / committee"
              value={`t=${form.threshold}  n=${form.committeeSize}  minContribs=${form.minValidContributions}`} />
            <KV label="Lottery α (bps)" value={form.lotteryAlphaBps} />
          </VStack>
        )}
      </StepCard>

      {/* ── Step 3: Watch Progress ────────────────────────────────────────── */}
      <StepCard n={3} title="Watch Round Progress" status={stepWatch}>
        {!roundId ? (
          <Text fontSize="sm" color="gray.500">Waiting for a round to be created…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            {round && (
              <>
                <HStack spacing={3} wrap="wrap">
                  <Text fontSize="xs" color="gray.400">Status:</Text>
                  {round.status === RoundStatus.Registration && <Badge colorScheme="yellow">Registration</Badge>}
                  {round.status === RoundStatus.Contribution && <Badge colorScheme="blue">Contribution</Badge>}
                  {round.status === RoundStatus.Finalized && <Badge colorScheme="cyan">Finalized</Badge>}
                  {round.status === RoundStatus.Aborted && <Badge colorScheme="red">Aborted</Badge>}
                  {round.status === RoundStatus.Completed && <Badge colorScheme="green">Completed</Badge>}
                  {stepWatch === 'active' && <Spinner size="xs" color="blue.300" />}
                </HStack>
                <Grid templateColumns="repeat(2, 1fr)" gap={2}>
                  <KV label="Seed block" value={`#${round.seedBlock}`} />
                  <KV label="Seed"
                    value={BigInt(round.seed) === 0n ? <Badge colorScheme="gray">pending</Badge> : short(round.seed)} />
                  <KV label="Claimed slots" value={`${round.claimedCount} / ${round.policy.committeeSize}`} />
                  <KV label="Contributions" value={`${round.contributionCount} / ${round.policy.committeeSize}`} />
                  {round.status === RoundStatus.Registration && (() => {
                    const delta = blocksRemaining(currentBlock, round.policy.registrationDeadlineBlock);
                    const color = delta === null ? 'gray.400' : delta > 20 ? 'green.300' : delta > 0 ? 'yellow.300' : 'red.400';
                    return (
                      <KV label="Registration closes"
                        value={<Text fontFamily="mono" fontSize="xs" color={color}>{formatBlocksRemaining(delta)} (block #{round.policy.registrationDeadlineBlock.toString()})</Text>} />
                    );
                  })()}
                  {round.status === RoundStatus.Contribution && (() => {
                    const delta = blocksRemaining(currentBlock, round.policy.contributionDeadlineBlock);
                    const color = delta === null ? 'gray.400' : delta > 20 ? 'green.300' : delta > 0 ? 'yellow.300' : 'red.400';
                    return (
                      <KV label="Contribution closes"
                        value={<Text fontFamily="mono" fontSize="xs" color={color}>{formatBlocksRemaining(delta)} (block #{round.policy.contributionDeadlineBlock.toString()})</Text>} />
                    );
                  })()}
                </Grid>

                {/* Hint text explaining what the round is waiting for */}
                {round.status === RoundStatus.Registration && (
                  <Text fontSize="xs" color="gray.500">
                    Waiting for DKG nodes to claim committee slots. After the registration deadline,
                    participants submit their key contributions.
                  </Text>
                )}
                {round.status === RoundStatus.Contribution && (
                  <Text fontSize="xs" color="gray.500">
                    Waiting for DKG nodes to submit ZK contributions and for one node to call{' '}
                    <Code fontSize="2xs">finalizeRound</Code> once the threshold is reached.
                    Contributions: {round.contributionCount}/{round.policy.minValidContributions} needed.
                  </Text>
                )}
                {round.status === RoundStatus.Finalized && (
                  <Text fontSize="xs" color="cyan.400">
                    Round finalized on-chain. Waiting to index the RoundFinalized event…
                  </Text>
                )}

                {/* Abort button — visible while round is non-terminal and not yet finalized */}
                {(round.status === RoundStatus.Registration || round.status === RoundStatus.Contribution) && (
                  <Box pt={1}>
                    <Button size="xs" colorScheme="red" variant="outline"
                      isLoading={abortBusy} onClick={handleAbort}>
                      Abort Round
                    </Button>
                    <Text fontSize="2xs" color="gray.600" mt={1}>
                      Organizer only — aborts the round so you can start fresh.
                    </Text>
                  </Box>
                )}

                {participants.length > 0 && (
                  <Box>
                    <Text fontSize="xs" color="gray.400" mb={1}>Selected participants:</Text>
                    <VStack align="start" spacing={1}>
                      {participants.map((p, i) => (
                        <HStack key={p} spacing={2}>
                          <Badge colorScheme="gray" fontFamily="mono" fontSize="2xs">#{i}</Badge>
                          <Code fontSize="xs">{p}</Code>
                        </HStack>
                      ))}
                    </VStack>
                  </Box>
                )}
              </>
            )}

            {finalizedInfo && (
              <>
                <Divider borderColor="gray.700" />
                <Box>
                  <Text fontSize="xs" color="green.400" mb={2} fontWeight="semibold">
                    ✓ Finalized at block #{finalizedInfo.blockNumber.toString()}
                  </Text>
                  <VStack align="start" spacing={2}>
                    <KV label="collectivePublicKeyHash" value={finalizedInfo.collectivePublicKeyHash} />
                    <KV label="aggregateCommitmentsHash" value={finalizedInfo.aggregateCommitmentsHash} />
                    <KV label="shareCommitmentHash" value={finalizedInfo.shareCommitmentHash} />
                  </VStack>
                  <Alert status="info" variant="left-accent" borderRadius="sm" mt={3} fontSize="xs">
                    <AlertIcon />
                    The <strong>collectivePublicKeyHash</strong> = keccak256(pubKey.x, pubKey.y). The actual
                    coordinates are extracted from the <Code fontSize="2xs">finalizeRound</Code> calldata
                    (transcript field) and shown in Step 4.
                  </Alert>
                </Box>
              </>
            )}

            {events.length > 0 && (
              <Box>
                <Text fontSize="xs" color="gray.400" mb={2}>On-chain events ({events.length}):</Text>
                <VStack align="stretch" spacing={1}>
                  {events.map((ev, i) => (
                    <HStack key={i} fontSize="xs" fontFamily="mono" bg="gray.900"
                      px={2} py={1} borderRadius="sm" spacing={3}>
                      <Badge colorScheme="cyan" fontSize="2xs">{ev.eventName}</Badge>
                      <Text color="gray.500">#{ev.blockNumber.toString()}</Text>
                      <Text color="gray.600" isTruncated>{short(ev.transactionHash)}</Text>
                    </HStack>
                  ))}
                </VStack>
              </Box>
            )}
          </VStack>
        )}
      </StepCard>

      {/* ── Step 4: Collective Public Key ────────────────────────────────── */}
      <StepCard n={4} title="DKG Collective Public Key" status={stepKey}>
        {!roundId || (Number(round?.contributionCount ?? 0) < Number(round?.policy.minValidContributions ?? 1) && !finalizedInfo) ? (
          <Text fontSize="sm" color="gray.500">
            Waiting for at least {round?.policy.minValidContributions ?? '?'} contribution(s)…
          </Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              The collective public key = sum of each contributor's zeroth Feldman commitment
              point (a<sub>i,0</sub>·G).{' '}
              {finalizedInfo
                ? 'Extracted from the on-chain finalizeRound calldata.'
                : 'Derived directly from contribution calldatas — identical to the finalized key.'}
            </Text>
            {collectivePubKeyBusy && (
              <HStack spacing={2}>
                <Spinner size="sm" color="purple.300" />
                <Text fontSize="xs" color="gray.400">
                  {finalizedInfo ? 'Fetching from finalizeRound calldata…' : 'Summing contribution commitments…'}
                </Text>
              </HStack>
            )}
            {collectivePubKey && (
              <VStack align="start" spacing={2}>
                <KV label="Curve" value="BabyJubJub (twisted Edwards / BN254 scalar field)" mono={false} />
                <KV label="Public key x" value={hex(collectivePubKey.x)} />
                <KV label="Public key y" value={hex(collectivePubKey.y)} />
                <Text fontSize="xs" color="gray.500">
                  Q = privKey₁·G + privKey₂·G + … (Shamir secret sharing over BabyJubJub)
                </Text>
              </VStack>
            )}
          </VStack>
        )}
      </StepCard>

      {/* ── Step 5: Encrypt ────────────────────────────────────────────────── */}
      <StepCard n={5} title="ElGamal Encrypt" status={stepEncrypt}>
        {!collectivePubKey ? (
          <Text fontSize="sm" color="gray.500">Waiting for collective public key…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              ElGamal on BabyJubJub: the message <em>m</em> is encoded as a scalar,
              a random ephemeral key <em>k</em> is chosen, and the ciphertext is
              (c₁, c₂) = (k·G, m·G + k·Q) where Q is the DKG collective public key and G = Base8.
            </Text>
            <HStack spacing={3} align="end">
              <FormControl w="200px">
                <FormLabel fontSize="xs" color="gray.300" mb={1}>Plaintext (integer)</FormLabel>
                <Input size="sm" fontFamily="mono" value={plaintext}
                  onChange={e => { setPlaintext(e.target.value); setCiphertext(null); setSubmittedCiphertext(false); setCombinedRecord(null); }} />
              </FormControl>
              <Button colorScheme="purple" size="sm" onClick={handleEncrypt}>Encrypt →</Button>
            </HStack>

            {encryptDetails && ciphertext && (
              <>
                <Divider borderColor="gray.700" />
                <Text fontSize="xs" color="gray.400" fontWeight="semibold">Encryption trace:</Text>
                <VStack align="start" spacing={2} pl={2} borderLeft="2px solid" borderColor="purple.700">
                  <KV label="m (plaintext)" value={`${plaintext}  = ${hex(BigInt(plaintext || '0'))}`} />
                  <KV label="k (ephemeral)" value={hex(encryptDetails.k)} />
                  <KV label="c₁ = k × G" value={
                    <VStack align="start" spacing={0}>
                      <Text>x: {hex(encryptDetails.c1[0])}</Text>
                      <Text>y: {hex(encryptDetails.c1[1])}</Text>
                    </VStack>
                  } />
                  <KV label="s = k × Q" value={
                    <VStack align="start" spacing={0}>
                      <Text>x: {hex(encryptDetails.s[0])}</Text>
                      <Text>y: {hex(encryptDetails.s[1])}</Text>
                    </VStack>
                  } />
                  <KV label="m·G" value={
                    <VStack align="start" spacing={0}>
                      <Text>x: {hex(encryptDetails.mPoint[0])}</Text>
                      <Text>y: {hex(encryptDetails.mPoint[1])}</Text>
                    </VStack>
                  } />
                  <KV label="c₂ = m·G + s" value={
                    <VStack align="start" spacing={0}>
                      <Text>x: {hex(ciphertext.c2[0])}</Text>
                      <Text>y: {hex(ciphertext.c2[1])}</Text>
                    </VStack>
                  } />
                </VStack>
                <Box>
                  <Text fontSize="xs" color="gray.400" mb={2}>Ciphertext (to be transmitted):</Text>
                  <MonoBox>
                    <Text color="gray.400">c₁.x = <Text as="span" color="yellow.300">{hex(ciphertext.c1[0])}</Text></Text>
                    <Text color="gray.400">c₁.y = <Text as="span" color="yellow.300">{hex(ciphertext.c1[1])}</Text></Text>
                    <Text color="gray.400">c₂.x = <Text as="span" color="yellow.300">{hex(ciphertext.c2[0])}</Text></Text>
                    <Text color="gray.400">c₂.y = <Text as="span" color="yellow.300">{hex(ciphertext.c2[1])}</Text></Text>
                  </MonoBox>
                </Box>
              </>
            )}
          </VStack>
        )}
      </StepCard>

      {/* ── Step 6: Submit to Nodes ───────────────────────────────────────── */}
      <StepCard n={6} title="Submit Ciphertext & Await Threshold Decryption" status={stepSubmit}>
        {!ciphertext ? (
          <Text fontSize="sm" color="gray.500">Waiting for ciphertext…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              The ciphertext is stored via the node's <Code fontSize="xs">/api/ciphertext/</Code> endpoint.
              Nodes pick it up on their next poll cycle, compute a ZK partial decryption, and submit it
              on-chain. Once the threshold is reached, one node runs <Code fontSize="xs">combineDecryptions</Code>{' '}
              to produce the final combined decryption and write it on-chain.
            </Text>
            {!submittedCiphertext ? (
              <Button colorScheme="teal" size="sm" alignSelf="start"
                isLoading={submitBusy} onClick={handleSubmitCiphertext}>
                Send Ciphertext to Nodes →
              </Button>
            ) : (
              <VStack align="start" spacing={3}>
                <HStack spacing={2}>
                  <Badge colorScheme="teal">Submitted</Badge>
                  <Text fontSize="xs" color="gray.400">
                    Ciphertext stored — waiting for threshold decryption…
                  </Text>
                  {!combinedRecord?.completed && <Spinner size="xs" color="teal.300" />}
                </HStack>
                {combinedRecord?.completed ? (
                  <HStack spacing={2}>
                    <Badge colorScheme="green">Combined on-chain</Badge>
                    <Text fontSize="xs" color="green.300">
                      plaintextHash: {short(combinedRecord.plaintextHash)}
                    </Text>
                  </HStack>
                ) : (
                  <Text fontSize="xs" color="gray.500" fontFamily="mono">
                    Polling every 3 s for CombinedDecryption event…
                  </Text>
                )}
                <Button variant="outline" colorScheme="gray" size="xs" alignSelf="start"
                  onClick={() => { setCiphertext(null); setEncryptDetails(null); setSubmittedCiphertext(false); setCombinedRecord(null); }}>
                  ← Reset
                </Button>
              </VStack>
            )}
          </VStack>
        )}
      </StepCard>

      {/* ── Step 7: Verify ───────────────────────────────────────────────── */}
      <StepCard n={7} title="Verify Recovered Plaintext" status={stepVerify}>
        {!combinedRecord?.completed ? (
          <Text fontSize="sm" color="gray.500">Waiting for combined decryption on-chain…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              The <Code fontSize="xs">plaintextHash</Code> stored on-chain is the raw plaintext
              scalar (not a keccak256 hash — the circuit stores the plaintext itself as bytes32).
              We compare it to the original plaintext entered in Step 5.
            </Text>
            {(() => {
              const onChainPlaintext = BigInt(combinedRecord.plaintextHash);
              const originalPlaintext = BigInt(plaintext || '0');
              const match = onChainPlaintext === originalPlaintext;
              return (
                <Box
                  bg={match ? 'green.900' : 'red.900'}
                  border="1px solid" borderColor={match ? 'green.500' : 'red.500'}
                  borderRadius="md" p={4}
                >
                  <Text fontSize="lg" fontFamily="mono" fontWeight="bold" color="white" mb={2}>
                    {match ? '✓' : '✗'} Recovered plaintext: {onChainPlaintext.toString()}
                  </Text>
                  <VStack align="start" spacing={1}>
                    <KV label="Original plaintext" value={originalPlaintext.toString()} />
                    <KV label="On-chain plaintextHash (bytes32)" value={combinedRecord.plaintextHash} />
                    <KV label="Ciphertext index" value={combinedRecord.ciphertextIndex.toString()} />
                  </VStack>
                  <Text fontSize="xs" color={match ? 'green.300' : 'red.300'} mt={3}>
                    {match
                      ? `Full roundtrip verified: encrypted ${originalPlaintext} with DKG key → threshold decryption → recovered ${onChainPlaintext}.`
                      : `Mismatch — on-chain value ${onChainPlaintext} ≠ original ${originalPlaintext}.`
                    }
                  </Text>
                </Box>
              );
            })()}
          </VStack>
        )}
      </StepCard>

      {/* ── Activity log ─────────────────────────────────────────────────────── */}
      <Box>
        <Text fontSize="xs" color="gray.500" mb={2} fontFamily="mono">ACTIVITY LOG</Text>
        <TerminalLog entries={log} />
      </Box>
    </VStack>
  );
}
