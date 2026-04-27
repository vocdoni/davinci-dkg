// Build a self-contained error report string suitable for pasting into a
// GitHub issue. The shape is intentionally human-readable rather than JSON
// — the consumer is a maintainer reading the issue body, not a bot.

export interface ErrorReportContext {
  route?: string
  chainId?: number
  chainName?: string
  walletAddress?: string
  roundId?: string
  blockNumber?: string | bigint | number
  buildVersion?: string
}

export function buildErrorReport(error: unknown, ctx: ErrorReportContext = {}): string {
  const errMsg = error instanceof Error ? error.message : String(error)
  const stack = error instanceof Error && error.stack ? error.stack : '(no stack)'
  const ts = new Date().toISOString()
  const ua = typeof navigator !== 'undefined' ? navigator.userAgent : '(no navigator)'

  const lines = [
    '## davinci-dkg UI error report',
    '',
    `- timestamp: ${ts}`,
    ctx.route ? `- route: ${ctx.route}` : null,
    ctx.chainName || ctx.chainId
      ? `- chain: ${ctx.chainName ?? '(unknown)'} (id ${ctx.chainId ?? '?'})`
      : null,
    ctx.walletAddress ? `- wallet: ${ctx.walletAddress}` : null,
    ctx.roundId ? `- round: ${ctx.roundId}` : null,
    ctx.blockNumber != null ? `- block: ${ctx.blockNumber.toString()}` : null,
    ctx.buildVersion ? `- build: ${ctx.buildVersion}` : null,
    `- user agent: ${ua}`,
    '',
    '### Error',
    '',
    '```',
    errMsg,
    '```',
    '',
    '### Stack',
    '',
    '```',
    stack,
    '```',
  ].filter(Boolean) as string[]

  return lines.join('\n')
}
