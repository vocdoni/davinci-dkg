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
import { useParams } from 'react-router-dom';
import { HashCell } from '../components/HashCell';
import { StatusBadge } from '../components/StatusBadge';
import { isZeroHash } from '../lib/format';
import { useRound, useRoundEvents } from '../lib/hooks';

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

export function RoundDetail() {
  const { id } = useParams<{ id: string }>();
  const roundId = parseRoundId(id);
  const roundQ = useRound(roundId ?? undefined);
  const eventsQ = useRoundEvents(roundId ?? undefined);

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
  if (roundQ.isError) {
    return (
      <Alert status="error" variant="left-accent" borderRadius="md">
        <AlertIcon />
        <Text>Failed to load round: {(roundQ.error as Error).message}</Text>
      </Alert>
    );
  }
  if (roundQ.isLoading || !roundQ.data) {
    return <Text color="gray.400">Loading round…</Text>;
  }
  const r = roundQ.data;
  const policy = r.policy;
  const committeeSize = Number(policy.committeeSize);
  const threshold = Number(policy.threshold);
  const claimed = Number(r.claimedCount);
  const contribs = Number(r.contributionCount);
  const partials = Number(r.partialDecryptionCount);
  const reveals = Number(r.revealedShareCount);

  const phaseProgress = (n: number) =>
    committeeSize > 0 ? Math.min(100, (n / committeeSize) * 100) : 0;

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
          <Field label="Registration deadline">
            #{policy.registrationDeadlineBlock.toString()}
          </Field>
          <Field label="Contribution deadline">
            #{policy.contributionDeadlineBlock.toString()}
          </Field>
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
