import type { ReactNode } from 'react'
import { Callout, KeyValue, SectionHeader, Stack, type KeyValueItem } from '~kit'

export interface PlaceholderProps {
  label: string
  title: ReactNode
  description: string
  /** Which stream delivers this page — shown in the notice. */
  owner: string
  /** Bullet list of what the finished page must contain (from EXPLORER.md §4). */
  contents: string[]
  /** Route params echoed back, so links can be verified before the data lands. */
  params?: KeyValueItem[]
}

/**
 * Every route in EXPLORER.md §4 exists from day one so navigation, deep links
 * and the Playwright sweep are real before the data layer arrives. Each
 * placeholder states what the page owes and who is writing it.
 */
export function Placeholder({ label, title, description, owner, contents, params }: PlaceholderProps) {
  return (
    <Stack>
      <SectionHeader size='page' label={label} title={title} description={description} />
      {params && params.length > 0 ? <KeyValue items={params} columns={2} /> : null}
      <Callout tone='info' title={`Delivered by ${owner}`}>
        <ul className='mt-1 list-disc space-y-1 pl-4'>
          {contents.map((item) => (
            <li key={item}>{item}</li>
          ))}
        </ul>
      </Callout>
    </Stack>
  )
}
