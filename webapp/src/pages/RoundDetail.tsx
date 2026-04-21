import {
  Alert,
  AlertIcon,
  Box,
  Divider,
  Grid,
  GridItem,
  Heading,
  HStack,
  Progress,
  SimpleGrid,
  Tag,
  Text,
  VStack,
} from '@chakra-ui/react';
import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { HashCell } from '../components/HashCell';
import { StatusBadge } from '../components/StatusBadge';
import { blocksRemaining, formatBlocksRemaining, isZeroHash } from '../lib/format';
import { useBlockNumber, useRound, useRoundEvents } from '../lib/hooks';

const ROUND_ID_RE = /^0x[0-9a-fA-F]{24}$/;

function parseRoundId(raw: string | undefined): `0x${string}` | null {
  if (!raw) return null;
  return ROUND_ID_RE.test(raw) ? (raw.toLowerCase() as `0x${string}`) : null;
}

function Field({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <Box>
      <Text fontSize="xs" textTransform="uppercase" color="gray.500" mb={1}>
        {label}
      </Text>
      <Box fontFamily="mono" fontSize="sm">
        {children}
      </Box>
    </Box>
  );
}

function Section({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <Box bg="gray.800" p={5} borderRadius="md" borderWidth="1px" borderColor="gray.700">
      <Heading size="md" mb={4} color="cyan.200">
        {title}
      </Heading>
      {children}
    </Box>
  );
}

function DeadlineField({
  label,
  deadline,
  currentBlock,
  activeWhen,
  status,
}: {
  label: string;
  deadline: bigint;
  currentBlock: bigint | undefined;
  activeWhen: number; // the round status value when this deadline is relevant
  status: number;
}) {
  const delta = blocksRemaining(currentBlock, deadline);
  const isActive = status === activeWhen;
  const color =
    delta === null ? 'gray.400'
    : delta > 20   ? 'green.300'
    : delta > 0    ? 'yellow.300'
    : 'red.400';

  return (
    <Field label={label}>
      <HStack spacing={2} align="baseline" wrap="wrap">
        <Text>#{deadline.toString()}</Text>
        {isActive && delta !== null && (
          <Text fontSize="xs" color={color} fontFamily="mono">
            ({formatBlocksRemaining(delta)})
          </Text>
        )}
      </HStack>
    </Field>
  );
}

export function RoundDetail() {
  const { id } = useParams<{ id: string }>();
  const roundId = parseRoundId(id);
  const roundQ = useRound(roundId ?? undefined);
  const eventsQ = useRoundEvents(roundId ?? undefined);
  const blockQ = useBlockNumber();

  if (!roundId) {
    return (
      <Alert status="error" variant="left-accent" borderRadius="md">
        <AlertIcon />
        <Box>
          <Text fontWeight="semibold">Invalid round ID</Text>
          <Text fontSize="sm" color="gray.400">
            Expected a 12-byte hex string (e.g. 0xd5941d7d0000000000000001).
          </Text>
        </Box>
      </Alert>
    );
  }
  useEffect(() => {
    if (roundQ.isError) console.error('[RoundDetail] Failed to load round:', roundQ.error);
  }, [roundQ.isError, roundQ.error]);

  if (roundQ.isLoading || (!roundQ.data && !roundQ.isError)) {
    return <Text color="gray.400">Loading round…</Text>;
  }
  if (!roundQ.data) {
    return <Text color="gray.400">Round data unavailable (check console for errors).</Text>;
  }
  const r = roundQ.data;
  const policy = r.policy;
  const dp = r.decryptionPolicy;
  const committeeSize = Number(policy.committeeSize);
  const threshold = Number(policy.threshold);
  const claimed = Number(r.claimedCount);
  const contribs = Number(r.contributionCount);
  const partials = Number(r.partialDecryptionCount);
  const reveals = Number(r.revealedShareCount);
  const ciphertexts = Number(r.ciphertextCount ?? 0);

  const phaseProgress = (n: number) =>
    committeeSize > 0 ? Math.min(100, (n / committeeSize) * 100) : 0;

  const dpActive =
    !!dp && (
      dp.ownerOnly ||
      Number(dp.maxDecryptions) > 0 ||
      dp.notBeforeBlock > 0n ||
      dp.notBeforeTimestamp > 0n ||
      dp.notAfterBlock > 0n ||
      dp.notAfterTimestamp > 0n
    );

  const formatTs = (ts: bigint) =>
    ts === 0n ? '—' : new Date(Number(ts) * 1000).toISOString().replace('T', ' ').slice(0, 19) + ' UTC';

  return (
    <VStack align="stretch" spacing={5}>
      <HStack>
        <Heading size="lg">Round</Heading>
        <HashCell value={roundId} head={12} tail={8} />
        <StatusBadge status={Number(r.status)} />
      </HStack>

      <Section title="Round info">
        <SimpleGrid columns={{ base: 2, md: 4 }} spacing={4}>
          <Field label="Nonce">{r.nonce.toString()}</Field>
          <Field label="Organizer">
            <HashCell value={r.organizer} />
          </Field>
          <Field label="Seed block">#{r.seedBlock.toString()}</Field>
          <Field label="Seed">
            {isZeroHash(r.seed) ? (
              <Tag size="sm" colorScheme="gray">
                pending
              </Tag>
            ) : (
              <HashCell value={r.seed} />
            )}
          </Field>
          <Field label="Lottery threshold">
            <HashCell value={`0x${r.lotteryThreshold.toString(16).padStart(64, '0')}`} />
          </Field>
        </SimpleGrid>
      </Section>

      <Section title="Policy">
        <SimpleGrid columns={{ base: 2, md: 4 }} spacing={4}>
          <Field label="Committee size">{committeeSize}</Field>
          <Field label="Threshold (t)">{threshold}</Field>
          <Field label="Min valid contribs">{policy.minValidContributions}</Field>
          <Field label="Lottery α (bps)">{policy.lotteryAlphaBps}</Field>
          <Field label="Seed delay">{policy.seedDelay}</Field>
          <DeadlineField
            label="Registration deadline"
            deadline={policy.registrationDeadlineBlock}
            currentBlock={blockQ.data}
            activeWhen={1}
            status={Number(r.status)}
          />
          <DeadlineField
            label="Contribution deadline"
            deadline={policy.contributionDeadlineBlock}
            currentBlock={blockQ.data}
            activeWhen={2}
            status={Number(r.status)}
          />
          <Field label="Disclosure">
            {policy.disclosureAllowed ? (
              <Tag size="sm" colorScheme="green">
                allowed
              </Tag>
            ) : (
              <Tag size="sm" colorScheme="gray">
                disabled
              </Tag>
            )}
          </Field>
        </SimpleGrid>
      </Section>

      <Section title="Decryption policy">
        {!dp || !dpActive ? (
          <HStack spacing={3} color="gray.400" fontSize="sm">
            <Tag size="sm" colorScheme="gray">open</Tag>
            <Text>Anyone can call <code>submitCiphertext</code> at any time (no caps).</Text>
          </HStack>
        ) : (
          <SimpleGrid columns={{ base: 2, md: 3 }} spacing={4}>
            <Field label="Owner-only">
              {dp.ownerOnly ? (
                <Tag size="sm" colorScheme="orange">owner-only</Tag>
              ) : (
                <Tag size="sm" colorScheme="gray">open</Tag>
              )}
            </Field>
            <Field label="Max decryptions">
              {Number(dp.maxDecryptions) > 0
                ? `${ciphertexts} / ${dp.maxDecryptions}`
                : `${ciphertexts} / unlimited`}
            </Field>
            <Field label="Not-before block">
              {dp.notBeforeBlock === 0n ? '—' : `#${dp.notBeforeBlock.toString()}`}
            </Field>
            <Field label="Not-before timestamp">{formatTs(dp.notBeforeTimestamp)}</Field>
            <Field label="Not-after block">
              {dp.notAfterBlock === 0n ? '—' : `#${dp.notAfterBlock.toString()}`}
            </Field>
            <Field label="Not-after timestamp">{formatTs(dp.notAfterTimestamp)}</Field>
          </SimpleGrid>
        )}
      </Section>

      <Section title="Phase progress">
        <Grid templateColumns="160px 1fr 80px" gap={3} alignItems="center">
          <GridItem>
            <Text fontSize="sm">Slot claims</Text>
          </GridItem>
          <GridItem>
            <Progress value={phaseProgress(claimed)} colorScheme="yellow" size="sm" rounded="sm" />
          </GridItem>
          <GridItem>
            <Text fontFamily="mono" fontSize="sm" textAlign="right">
              {claimed}/{committeeSize}
            </Text>
          </GridItem>

          <GridItem>
            <Text fontSize="sm">Contributions</Text>
          </GridItem>
          <GridItem>
            <Progress value={phaseProgress(contribs)} colorScheme="blue" size="sm" rounded="sm" />
          </GridItem>
          <GridItem>
            <Text fontFamily="mono" fontSize="sm" textAlign="right">
              {contribs}/{committeeSize}
            </Text>
          </GridItem>

          <GridItem>
            <Text fontSize="sm">Ciphertexts</Text>
          </GridItem>
          <GridItem>
            <Progress
              value={
                Number(dp?.maxDecryptions ?? 0) > 0
                  ? Math.min(100, (ciphertexts / Number(dp!.maxDecryptions)) * 100)
                  : ciphertexts > 0 ? 100 : 0
              }
              colorScheme="teal"
              size="sm"
              rounded="sm"
            />
          </GridItem>
          <GridItem>
            <Text fontFamily="mono" fontSize="sm" textAlign="right">
              {Number(dp?.maxDecryptions ?? 0) > 0
                ? `${ciphertexts}/${dp!.maxDecryptions}`
                : `${ciphertexts}`}
            </Text>
          </GridItem>

          <GridItem>
            <Text fontSize="sm">Partial decryptions</Text>
          </GridItem>
          <GridItem>
            <Progress value={phaseProgress(partials)} colorScheme="purple" size="sm" rounded="sm" />
          </GridItem>
          <GridItem>
            <Text fontFamily="mono" fontSize="sm" textAlign="right">
              {partials}/{committeeSize}
            </Text>
          </GridItem>

          <GridItem>
            <Text fontSize="sm">Revealed shares</Text>
          </GridItem>
          <GridItem>
            <Progress value={phaseProgress(reveals)} colorScheme="pink" size="sm" rounded="sm" />
          </GridItem>
          <GridItem>
            <Text fontFamily="mono" fontSize="sm" textAlign="right">
              {reveals}/{committeeSize}
            </Text>
          </GridItem>
        </Grid>
      </Section>

      <Section title={`Selected participants (${r.participants?.length ?? 0})`}>
        {r.participants && r.participants.length > 0 ? (
          <SimpleGrid columns={{ base: 1, md: 2 }} spacing={2}>
            {r.participants.map((p: string, i: number) => (
              <HStack
                key={p}
                bg="gray.900"
                px={3}
                py={2}
                borderRadius="sm"
                borderWidth="1px"
                borderColor="gray.700"
              >
                <Text fontFamily="mono" fontSize="xs" color="gray.500" w="30px">
                  #{i}
                </Text>
                <HashCell value={p} head={10} tail={8} />
              </HStack>
            ))}
          </SimpleGrid>
        ) : (
          <Text color="gray.400">No participants selected yet.</Text>
        )}
      </Section>

      <Section title={`Events (${eventsQ.data?.length ?? 0})`}>
        {eventsQ.isLoading && <Text color="gray.400">Loading events…</Text>}
        {eventsQ.data && eventsQ.data.length === 0 && (
          <Text color="gray.400">No events found.</Text>
        )}
        <VStack align="stretch" spacing={2}>
          {eventsQ.data?.map((ev: any, i: number) => (
            <Box
              key={`${ev.transactionHash}-${i}`}
              bg="gray.900"
              p={3}
              borderRadius="sm"
              borderWidth="1px"
              borderColor="gray.700"
            >
              <HStack justify="space-between" mb={1}>
                <Tag colorScheme="cyan" fontFamily="mono">
                  {ev.eventName}
                </Tag>
                <Text fontFamily="mono" fontSize="xs" color="gray.500">
                  block #{ev.blockNumber?.toString()}
                </Text>
              </HStack>
              <HStack fontSize="xs" color="gray.400">
                <Text>tx</Text>
                <HashCell value={ev.transactionHash} head={8} tail={6} />
              </HStack>
              {ev.args && (
                <Box mt={2} pl={3} borderLeft="2px solid" borderColor="gray.700">
                  {Object.entries(ev.args).map(([k, v]) => {
                    if (k === 'roundId') return null;
                    const s = typeof v === 'bigint' ? v.toString() : String(v);
                    const isHash = s.startsWith('0x') && s.length >= 42;
                    return (
                      <HStack key={k} spacing={2} fontSize="xs">
                        <Text color="gray.500" w="160px">
                          {k}
                        </Text>
                        {isHash ? <HashCell value={s} /> : <Text fontFamily="mono">{s}</Text>}
                      </HStack>
                    );
                  })}
                </Box>
              )}
            </Box>
          ))}
        </VStack>
      </Section>

      <Divider />
    </VStack>
  );
}
