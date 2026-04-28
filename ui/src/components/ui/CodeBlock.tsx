import { useState } from 'react'
import { Box, HStack, IconButton, Text, Tooltip } from '@chakra-ui/react'
import { LuCheck, LuClipboardCopy } from 'react-icons/lu'
import { Highlight, themes, type Language } from 'prism-react-renderer'

interface Props {
  /** Code body. Multiline strings render as a <pre>; the copy button copies it verbatim. */
  children: string
  /** Optional caption shown above the block (e.g. file name). Rendered as
   *  a window-chrome strip above the code itself. */
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
// highlighting. Editorial chrome:
//   ── caption strip   small mono filename + language tag, sat above the
//                      code on a darker plate; reads like a manuscript
//                      window header.
//   ── copy button     subtle, only colour-shifts on hover; success
//                      confirmation by a phosphor checkmark, not a toast.
//   ── code            JetBrains Mono on a deep ink panel, with a soft
//                      gold rule on the left edge marking it as code.
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

  const langTag = language && language !== 'plain' ? language : null

  return (
    <Box
      position='relative'
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='md'
      bg='surface.mono'
      overflow='hidden'
      _hover={{ borderColor: 'border' }}
      transition='border-color 0.15s'
    >
      {/* Caption strip — file name on the left, language tag on the right.
          Always renders if caption OR language is set, so even un-named
          snippets get the small "tsx" / "bash" tag. */}
      {(caption || langTag) && (
        <HStack
          justify='space-between'
          px={3}
          py={2}
          borderBottomWidth='1px'
          borderColor='border.subtle'
          bg='canvas.deep'
        >
          {caption ? (
            <Text fontFamily='mono' fontSize='2xs' color='ink.2' letterSpacing='0.04em'>
              {caption}
            </Text>
          ) : (
            <Box />
          )}
          {langTag && (
            <Text
              fontFamily='mono'
              fontSize='2xs'
              color='ink.4'
              letterSpacing='0.06em'
              textTransform='uppercase'
            >
              {langTag}
            </Text>
          )}
        </HStack>
      )}

      <Box position='relative'>
        {/* Soft gold rule on the left edge — marks the block as code/data
            in the editorial layout. */}
        <Box
          position='absolute'
          left={0}
          top={0}
          bottom={0}
          w='2px'
          bg='accent.fg'
          opacity={0.32}
        />

        {language === 'plain' ? (
          <Box
            as='pre'
            m={0}
            py={3.5}
            pl={4}
            pr={11}
            whiteSpace='pre'
            overflowX='auto'
            maxH={maxH}
            overflowY={maxH ? 'auto' : undefined}
            fontFamily='mono'
            fontSize='xs'
            color='ink.1'
            lineHeight='1.6'
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
                py={3.5}
                pl={4}
                pr={11}
                whiteSpace='pre'
                overflowX='auto'
                maxH={maxH}
                overflowY={maxH ? 'auto' : undefined}
                fontFamily='mono'
                fontSize='xs'
                lineHeight='1.6'
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

        {/* Copy button — quietly absent until hover (and accessible via
            keyboard). Coloured shift on success. */}
        <Tooltip.Root>
          <Tooltip.Trigger asChild>
            <IconButton
              aria-label={copied ? 'Copied' : 'Copy snippet'}
              size='xs'
              variant='ghost'
              onClick={onCopy}
              position='absolute'
              top={2}
              right={2}
              color={copied ? 'live.fg' : 'ink.3'}
              bg='transparent'
              _hover={{ color: copied ? 'live.fg' : 'accent.fg', bg: 'rgba(255,255,255,0.04)' }}
            >
              {copied ? <LuCheck /> : <LuClipboardCopy />}
            </IconButton>
          </Tooltip.Trigger>
          <Tooltip.Positioner>
            <Tooltip.Content fontFamily='sans' fontSize='xs'>
              {copied ? 'Copied' : 'Copy snippet'}
            </Tooltip.Content>
          </Tooltip.Positioner>
        </Tooltip.Root>
      </Box>
    </Box>
  )
}
