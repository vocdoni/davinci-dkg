import type { RowData } from '@tanstack/react-table'

export type CellAlign = 'left' | 'right' | 'center'

// Column-level presentation hints, merged into TanStack's `meta` slot so a
// column definition stays one object. `numeric` is the important one: every
// number in a table is JetBrains Mono with tabular figures, right-aligned, so
// digits line up down the column.
declare module '@tanstack/react-table' {
  interface ColumnMeta<TData extends RowData, TValue> {
    align?: CellAlign
    numeric?: boolean
    /** Fixed track width, e.g. `'160px'`. Omit for a flexible column. */
    width?: string
    /** Tooltip on the header cell — room for the protocol detail. */
    headerTooltip?: string
    /** Plain-text projection of the cell, for a future CSV export. */
    exportValue?: (value: TValue, row: TData) => string
  }
}

// TanStack makes `TValue` invariant, so a heterogeneous array of column
// definitions (a string column next to a bigint column) has no single sound
// element type. `any` here is the library's own documented escape hatch.
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type AnyColumnDef<TData> = import('@tanstack/react-table').ColumnDef<TData, any>
