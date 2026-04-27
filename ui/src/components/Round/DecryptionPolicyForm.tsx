import { Field, Input, SimpleGrid, Stack, Switch, Text } from '@chakra-ui/react'

export interface DecryptionPolicyFormState {
  ownerOnly: boolean
  maxDecryptions: string
  notBeforeBlock: string
  notBeforeTimestamp: string
  notAfterBlock: string
  notAfterTimestamp: string
}

export const defaultDecryptionPolicyForm: DecryptionPolicyFormState = {
  ownerOnly: false,
  maxDecryptions: '',
  notBeforeBlock: '',
  notBeforeTimestamp: '',
  notAfterBlock: '',
  notAfterTimestamp: '',
}

interface Props {
  value: DecryptionPolicyFormState
  onChange: (next: DecryptionPolicyFormState) => void
  disabled?: boolean
}

const fields: { key: keyof Omit<DecryptionPolicyFormState, 'ownerOnly'>; label: string; help: string }[] = [
  { key: 'maxDecryptions', label: 'Max ciphertexts', help: 'Per round; blank = unlimited.' },
  { key: 'notBeforeBlock', label: 'Not-before block', help: 'submitCiphertext reverts before this block.' },
  { key: 'notAfterBlock', label: 'Not-after block', help: 'submitCiphertext reverts after this block.' },
  { key: 'notBeforeTimestamp', label: 'Not-before timestamp', help: 'Unix seconds; 0 disables.' },
  { key: 'notAfterTimestamp', label: 'Not-after timestamp', help: 'Unix seconds; 0 disables.' },
]

export function DecryptionPolicyForm({ value, onChange, disabled }: Props) {
  return (
    <Stack gap={4}>
      <Field.Root disabled={disabled} display='flex' alignItems='center' gap={3}>
        <Switch.Root
          checked={value.ownerOnly}
          onCheckedChange={(d) => onChange({ ...value, ownerOnly: d.checked })}
        >
          <Switch.HiddenInput />
          <Switch.Control />
        </Switch.Root>
        <Text fontSize='sm'>Owner-only — only the round organizer may submit ciphertexts</Text>
      </Field.Root>
      <SimpleGrid columns={{ base: 1, md: 2 }} gap={3}>
        {fields.map((f) => (
          <Field.Root key={f.key} disabled={disabled}>
            <Field.Label fontSize='xs'>{f.label}</Field.Label>
            <Input
              size='sm'
              fontFamily='mono'
              value={value[f.key]}
              onChange={(e) => onChange({ ...value, [f.key]: e.target.value })}
              placeholder='0'
              inputMode='numeric'
            />
            <Field.HelperText fontSize='2xs'>{f.help}</Field.HelperText>
          </Field.Root>
        ))}
      </SimpleGrid>
    </Stack>
  )
}
