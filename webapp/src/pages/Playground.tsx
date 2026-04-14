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
    lotteryAlphaBps: '15000', seedDelay: '5',
    regDeadlineOffset: '100', contribDeadlineOffset: '200',
    disclosureAllowed: false,
  });

  // ── Round state
  const [writer, setWriter] = useState<DKGWriter | null>(null);
  const [createBusy, setCreateBusy] = useState(false);
  const [createError, setCreateError] = useState('');
  const [txHash, setTxHash] = useState('');
  const [roundId, setRoundId] = useState<`0x${string}` | null>(null);
  const [round, setRound] = useState<Round | null>(null);
  const [participants, setParticipants] = useState<Address[]>([]);
  const [events, setEvents] = useState<RoundEvent[]>([]);
  const [finalizedInfo, setFinalizedInfo] = useState<{
    collectivePublicKeyHash: string;
    aggregateCommitmentsHash: string;
    shareCommitmentHash: string;
    blockNumber: bigint;
  } | null>(null);

  // ── Crypto state
  const [keyPair, setKeyPair] = useState<{ privKey: bigint; pubKey: BabyJubPoint } | null>(null);
  const [plaintext, setPlaintext] = useState('42');
  const [ciphertext, setCiphertext] = useState<ElGamalCiphertext | null>(null);
  const [encryptDetails, setEncryptDetails] = useState<{
    k: bigint; c1: BabyJubPoint; s: BabyJubPoint; mPoint: BabyJubPoint;
  } | null>(null);
  const [decryptDetails, setDecryptDetails] = useState<{
    s: BabyJubPoint; negS: BabyJubPoint; mPoint: BabyJubPoint; iterations: number;
  } | null>(null);
  const [recovered, setRecovered] = useState<bigint | null>(null);

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
  const stepWatch: StepStatus = !roundId ? 'pending' : finalizedInfo ? 'done' : round?.status === RoundStatus.Aborted ? 'error' : 'active';
  const stepKey: StepStatus = !finalizedInfo ? 'pending' : keyPair ? 'done' : 'active';
  const stepEncrypt: StepStatus = !keyPair ? 'pending' : ciphertext ? 'done' : 'active';
  const stepDecrypt: StepStatus = !ciphertext ? 'pending' : recovered !== null ? 'done' : 'active';

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
        registryAddress: cfg.registryAddress,
      });
      setWriter(w);

      addLog(`Connected: ${addr}`, 'success');
      addLog(`Chain: ${cfg.chainName} (id=${cfg.chainId})  RPC: ${cfg.rpcUrl}`, 'chain');
      addLog(`Manager:  ${cfg.managerAddress}`, 'chain');
      addLog(`Registry: ${cfg.registryAddress}`, 'chain');
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
      const regOffset = BigInt(form.regDeadlineOffset || '100');
      const contribOffset = BigInt(form.contribDeadlineOffset || '200');

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

    async function poll() {
      try {
        const [r, parts] = await Promise.all([
          writer!.getRound(roundId!),
          writer!.selectedParticipants(roundId!),
        ]);
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
          if (lastContribs > 0) addLog(`Contributions submitted: ${lastContribs}/${r.policy.committeeSize}`, 'info');
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
            addLog('─── Round finalized ───', 'success');
            addLog(`collectivePublicKeyHash: ${ev.collectivePublicKeyHash}`, 'success');
            addLog(`aggregateCommitmentsHash: ${ev.aggregateCommitmentsHash}`, 'info');
            addLog(`shareCommitmentHash:      ${ev.shareCommitmentHash}`, 'info');
          }
          clearInterval(pollRef.current!);
          pollRef.current = null;
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

  // ── Generate demo key pair
  async function handleGenerateKey() {
    addLog('─── Generating BabyJubJub key pair ───', 'crypto');
    const eg = await buildElGamal();
    const kp = eg.generateKeyPair();
    setKeyPair(kp);
    addLog(`privKey: ${hex(kp.privKey)}`, 'crypto');
    addLog(`pubKey.x: ${hex(kp.pubKey[0])}`, 'crypto');
    addLog(`pubKey.y: ${hex(kp.pubKey[1])}`, 'crypto');
  }

  // ── Encrypt
  async function handleEncrypt() {
    if (!keyPair) return;
    const msg = BigInt(plaintext || '0');
    addLog(`─── Encrypting plaintext m=${msg} ───`, 'crypto');

    const eg = await buildElGamal();
    const k = eg.randomScalar();
    addLog(`Ephemeral k: ${hex(k)}`, 'crypto');

    const c1 = eg.mulPoint([5299619240641551281634865583518297030282874472190772894086521144482721001553n,
      16950150798460657717958625567821834550301663161624707787222815936182638968203n], k);
    const s = eg.mulPoint(keyPair.pubKey, k);
    const mPoint = eg.mulPoint([5299619240641551281634865583518297030282874472190772894086521144482721001553n,
      16950150798460657717958625567821834550301663161624707787222815936182638968203n], msg);

    addLog(`c1 = k × G:     (${hex(c1[0])}, ${hex(c1[1])})`, 'crypto');
    addLog(`s  = k × Q:     (${hex(s[0])}, ${hex(s[1])})   Q = pubKey`, 'crypto');
    addLog(`m × G:          (${hex(mPoint[0])}, ${hex(mPoint[1])})`, 'crypto');

    const ct = eg.encrypt(msg, keyPair.pubKey, k);
    addLog(`c2 = m×G + s:   (${hex(ct.c2[0])}, ${hex(ct.c2[1])})`, 'crypto');
    addLog(`Ciphertext ready.`, 'success');

    setEncryptDetails({ k, c1, s, mPoint });
    setCiphertext(ct);
    setDecryptDetails(null);
    setRecovered(null);
  }

  // ── Decrypt
  async function handleDecrypt() {
    if (!ciphertext || !keyPair) return;
    const msg = BigInt(plaintext || '0');
    addLog('─── Decrypting ───', 'crypto');

    const eg = await buildElGamal();

    const s = eg.mulPoint(ciphertext.c1, keyPair.privKey);
    addLog(`s = privKey × c1: (${hex(s[0])}, ${hex(s[1])})`, 'crypto');

    // Negate s: -(x,y) = (-x mod p, y) on twisted Edwards
    // BabyJubJub base field prime (same as the BN254 scalar field)
    const BJJ_P = 21888242871839275222246405745257275088548364400416034343698204186575808495617n;
    const negS: BabyJubPoint = [(BJJ_P - s[0]) % BJJ_P, s[1]];
    addLog(`-s (x negated):   (${hex(negS[0])}, ${hex(negS[1])})`, 'crypto');

    const mPoint = eg.addPoint(ciphertext.c2, negS);
    addLog(`mPoint = c2 + (-s): (${hex(mPoint[0])}, ${hex(mPoint[1])})`, 'crypto');
    addLog('Solving discrete log (brute-force, valid for m < 2²⁰)…', 'crypto');

    const recovered = eg.decrypt(ciphertext, keyPair.privKey);
    addLog(`Found m = ${recovered} (after ${recovered + 1n} iterations)`, 'success');

    setDecryptDetails({ s, negS, mPoint, iterations: Number(recovered) + 1 });
    setRecovered(recovered);

    if (recovered === msg) {
      addLog(`✓ Verified: decrypted value ${recovered} matches original plaintext ${msg}`, 'success');
    } else {
      addLog(`✗ Mismatch: decrypted ${recovered} ≠ original ${msg}`, 'error');
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
            <KV label="Registry contract" value={writer?.registryAddress ?? '–'} />
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
                </Grid>
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
                    coordinates are encoded in the <Code fontSize="2xs">finalizeRound</Code> calldata
                    (transcript field). For this demo we generate a local BabyJubJub key pair in Step 4.
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

      {/* ── Step 4: Public Key ────────────────────────────────────────────── */}
      <StepCard n={4} title="Encryption Key Setup" status={stepKey}>
        {!finalizedInfo ? (
          <Text fontSize="sm" color="gray.500">Waiting for round finalization…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              In a production DKG flow the collective public key coordinates are recovered
              by decoding the <Code fontSize="xs">transcript</Code> parameter of the
              <Code fontSize="xs"> finalizeRound</Code> transaction. Below we generate a
              fresh BabyJubJub key pair to demonstrate the cryptographic primitives.
              The curve and operations are identical to those used for the actual DKG key.
            </Text>
            <Text fontSize="xs" color="gray.500" fontFamily="mono">
              Curve: BabyJubJub (twisted Edwards form on BN254 scalar field)
              ax² + y² = 1 + dx²y²  where a=168700, d=168696
            </Text>
            {!keyPair ? (
              <Button colorScheme="purple" size="sm" onClick={handleGenerateKey} alignSelf="start">
                Generate Key Pair
              </Button>
            ) : (
              <VStack align="start" spacing={2}>
                <KV label="Private key (scalar)" value={hex(keyPair.privKey)} />
                <KV label="Public key x" value={hex(keyPair.pubKey[0])} />
                <KV label="Public key y" value={hex(keyPair.pubKey[1])} />
                <Text fontSize="xs" color="gray.500">
                  pubKey = privKey × Base8  where Base8 is the standard BabyJubJub generator
                </Text>
                <Button colorScheme="purple" variant="outline" size="xs"
                  onClick={async () => { setKeyPair(null); setCiphertext(null); setRecovered(null); setTimeout(handleGenerateKey, 50); }}>
                  Regenerate
                </Button>
              </VStack>
            )}
          </VStack>
        )}
      </StepCard>

      {/* ── Step 5: Encrypt ────────────────────────────────────────────────── */}
      <StepCard n={5} title="ElGamal Encrypt" status={stepEncrypt}>
        {!keyPair ? (
          <Text fontSize="sm" color="gray.500">Waiting for key pair generation…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              ElGamal on BabyJubJub: the message <em>m</em> is encoded as a scalar,
              a random ephemeral key <em>k</em> is chosen, and the ciphertext is
              (c₁, c₂) = (k·G, m·G + k·Q) where Q is the public key and G = Base8.
            </Text>
            <HStack spacing={3} align="end">
              <FormControl w="200px">
                <FormLabel fontSize="xs" color="gray.300" mb={1}>Plaintext (integer)</FormLabel>
                <Input size="sm" fontFamily="mono" value={plaintext}
                  onChange={e => { setPlaintext(e.target.value); setCiphertext(null); setRecovered(null); }} />
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

      {/* ── Step 6: Decrypt ────────────────────────────────────────────────── */}
      <StepCard n={6} title="ElGamal Decrypt" status={stepDecrypt}>
        {!ciphertext ? (
          <Text fontSize="sm" color="gray.500">Waiting for ciphertext…</Text>
        ) : (
          <VStack align="stretch" spacing={4}>
            <Text fontSize="sm" color="gray.400">
              Decryption: compute <em>s = privKey · c₁</em> (the shared secret), then
              <em> mPoint = c₂ − s = c₂ + (−s)</em>. In twisted Edwards −(x,y) = (−x, y).
              Finally recover <em>m</em> by brute-force discrete log: iterate <em>i·G</em>
              until it equals mPoint (practical only for small m).
            </Text>
            {recovered === null ? (
              <Button colorScheme="green" size="sm" alignSelf="start" onClick={handleDecrypt}>
                Decrypt →
              </Button>
            ) : (
              <>
                {decryptDetails && (
                  <VStack align="start" spacing={2} pl={2} borderLeft="2px solid" borderColor="green.700">
                    <KV label="s = privKey · c₁" value={
                      <VStack align="start" spacing={0}>
                        <Text>x: {hex(decryptDetails.s[0])}</Text>
                        <Text>y: {hex(decryptDetails.s[1])}</Text>
                      </VStack>
                    } />
                    <KV label="−s  (negate x)" value={
                      <VStack align="start" spacing={0}>
                        <Text>x: {hex(decryptDetails.negS[0])}</Text>
                        <Text>y: {hex(decryptDetails.negS[1])}</Text>
                      </VStack>
                    } />
                    <KV label="mPoint = c₂ + (−s)" value={
                      <VStack align="start" spacing={0}>
                        <Text>x: {hex(decryptDetails.mPoint[0])}</Text>
                        <Text>y: {hex(decryptDetails.mPoint[1])}</Text>
                      </VStack>
                    } />
                    <KV label="DLOG search" value={`${decryptDetails.iterations} iteration(s)`} />
                  </VStack>
                )}
                <Box bg={recovered === BigInt(plaintext || '0') ? 'green.900' : 'red.900'}
                  border="1px solid" borderColor={recovered === BigInt(plaintext || '0') ? 'green.500' : 'red.500'}
                  borderRadius="md" p={4}>
                  <HStack>
                    <Text fontSize="lg" fontFamily="mono" fontWeight="bold" color="white">
                      {recovered === BigInt(plaintext || '0') ? '✓' : '✗'} Recovered plaintext: {recovered.toString()}
                    </Text>
                  </HStack>
                  {recovered === BigInt(plaintext || '0') ? (
                    <Text fontSize="xs" color="green.300" mt={1}>
                      Matches original plaintext {plaintext}. Encrypt/decrypt roundtrip verified.
                    </Text>
                  ) : (
                    <Text fontSize="xs" color="red.300" mt={1}>
                      Mismatch — recovered {recovered.toString()} ≠ {plaintext}.
                    </Text>
                  )}
                </Box>
                <Button variant="outline" colorScheme="gray" size="xs" alignSelf="start"
                  onClick={() => { setCiphertext(null); setRecovered(null); setEncryptDetails(null); }}>
                  ← Reset encrypt/decrypt
                </Button>
              </>
            )}
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
