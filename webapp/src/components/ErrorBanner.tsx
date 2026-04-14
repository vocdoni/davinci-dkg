import { Alert, AlertIcon, Text } from '@chakra-ui/react';

interface Props {
  error: unknown;
  title?: string;
}

export function ErrorBanner({ error, title = 'Failed to load' }: Props) {
  const message = error instanceof Error ? error.message : String(error);
  return (
    <Alert status="error" variant="left-accent" borderRadius="md">
      <AlertIcon />
      <Text>
        {title}: <Text as="span" fontFamily="mono">{message}</Text>
      </Text>
    </Alert>
  );
}
