import { useState } from 'react'
import { Box, HStack, IconButton, Text, Tooltip } from '@chakra-ui/react'
import { LuCheck, LuClipboardCopy } from 'react-icons/lu'
import { Highlight, themes, type Language } from 'prism-react-renderer'

interface Props {
  /** Code body. Multiline strings render as a <pre>; the copy button copies it verbatim. */
  children: string
  /** Optional caption shown above the block (e.g. file name, language label). */
  caption?: string
  /**
   * Prism language token for syntax highlighting (e.g. 'tsx', 'ts',
   * 'bash'). Omit for terminal-style plaintext rendering — useful for
   * shell snippets where highlighting just adds noise.
   */
  language?: Language | 'plain'
  /** Defaults to a sensible terminal-y dark style; override via Chakra props if needed. */
  maxH?: string | number
}

// Pre-formatted, copyable code block with optional Prism syntax
// highlighting. The copy button echoes the HashCell pattern (clipboard
// icon → checkmark for ~1.5s) so users across the app form one mental
// model for "this is copyable".
//
// We use the `vsDark` theme because (a) it's the closest visual match to
// the rest of the dark UI, and (b) it ships in prism-react-renderer with
// no extra CSS file to wire up.
export function CodeBlock({ children, caption, language = 'plain', maxH }: Props) {
  const [copied, setCopied] = useState(false)
  const onCopy = async () => {
    try {
      await navigator.clipboard.writeText(children)
      setCopied(true)
      setTimeout(() => setCopied(false), 1500)
    } catch {
      console.warn('Clipboard write failed; printing snippet to console:\n', children)
    }
  }

  return (
    <Box position='relative'>
      {caption && (
        <Text fontSize='2xs' color='gray.500' mb={1.5} fontFamily='mono'>
          {caption}
        </Text>
      )}
      <Box
        bg='#1e1e1e'
        borderWidth='1px'
        borderColor='gray.800'
        borderRadius='md'
        overflow='hidden'
      >
        {language === 'plain' ? (
          <Box
            as='pre'
            m={0}
            p={3}
            pr={10}
            whiteSpace='pre'
            overflowX='auto'
            maxH={maxH}
            overflowY={maxH ? 'auto' : undefined}
            fontFamily='mono'
            fontSize='xs'
            color='gray.200'
            style={{ tabSize: 2 }}
          >
            {children}
          </Box>
        ) : (
          <Highlight theme={themes.vsDark} code={children.trimEnd()} language={language}>
            {({ className, style, tokens, getLineProps, getTokenProps }) => (
              <Box
                as='pre'
                className={className}
                style={{ ...style, background: 'transparent', tabSize: 2, margin: 0 }}
                m={0}
                p={3}
                pr={10}
                whiteSpace='pre'
                overflowX='auto'
                maxH={maxH}
                overflowY={maxH ? 'auto' : undefined}
                fontFamily='mono'
                fontSize='xs'
              >
                {tokens.map((line, i) => (
                  <div key={i} {...getLineProps({ line })}>
                    {line.map((token, key) => (
                      <span key={key} {...getTokenProps({ token })} />
                    ))}
                  </div>
                ))}
              </Box>
            )}
          </Highlight>
        )}
      </Box>
      <HStack position='absolute' top={caption ? 8 : 2} right={2}>
        <Tooltip.Root>
          <Tooltip.Trigger asChild>
            <IconButton
              aria-label='Copy snippet'
              size='xs'
              variant='ghost'
              onClick={onCopy}
              color='gray.400'
              _hover={{ color: 'cyan.300', bg: 'transparent' }}
            >
              {copied ? <LuCheck /> : <LuClipboardCopy />}
            </IconButton>
          </Tooltip.Trigger>
          <Tooltip.Positioner>
            <Tooltip.Content>{copied ? 'Copied' : 'Copy'}</Tooltip.Content>
          </Tooltip.Positioner>
        </Tooltip.Root>
      </HStack>
    </Box>
  )
}
