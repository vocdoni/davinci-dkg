import { HStack, IconButton, Text, Tooltip, useToast } from '@chakra-ui/react';
import { copyToClipboard, shortHash } from '../lib/format';

interface Props {
  value: string | undefined | null;
  head?: number;
  tail?: number;
  mono?: boolean;
  full?: boolean;
}

export function HashCell({ value, head = 6, tail = 4, mono = true, full = false }: Props) {
  const toast = useToast();
  if (!value) return <Text color="gray.500">—</Text>;
  const display = full ? value : shortHash(value, head, tail);
  return (
    <HStack spacing={1}>
      <Tooltip label={value} placement="top" hasArrow>
        <Text fontFamily={mono ? 'mono' : undefined} fontSize="sm">
          {display}
        </Text>
      </Tooltip>
      <IconButton
        aria-label="copy"
        size="xs"
        variant="ghost"
        icon={<span style={{ fontSize: 11 }}>⧉</span>}
        onClick={() => {
          copyToClipboard(value);
          toast({ title: 'Copied', status: 'success', duration: 800, isClosable: false });
        }}
      />
    </HStack>
  );
}
