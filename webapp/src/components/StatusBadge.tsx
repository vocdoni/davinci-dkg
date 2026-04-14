import { Badge } from '@chakra-ui/react';
import { roundStatusColor, roundStatusLabel } from '../lib/abi';

export function StatusBadge({ status }: { status: number | bigint }) {
  const s = typeof status === 'bigint' ? Number(status) : status;
  return (
    <Badge colorScheme={roundStatusColor(s)} fontFamily="mono">
      {roundStatusLabel(s)}
    </Badge>
  );
}
