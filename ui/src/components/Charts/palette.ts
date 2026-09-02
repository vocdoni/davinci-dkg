// Categorical palette for the explorer's charts.
//
// The editorial theme's own colours are a *status* vocabulary — gold for
// "active", phosphor for "live", coral for "failed" — so reusing them to tell
// four data series apart would overload them. These four hues are a separate,
// fixed-order categorical ramp, checked (not eyeballed) against the dark chart
// surface #0f0f13 with the data-viz validator:
//
//   lightness band  OKLCH L 0.48–0.67 ✓   chroma floor C ≥ 0.10 ✓
//   CVD separation  worst all-pairs ΔE 11.1 (protan/deutan sim) ✓
//   normal vision   worst all-pairs ΔE 21.8 ✓
//   contrast        all ≥ 3:1 vs surface ✓
//
// The mapping is by *metric*, not by rank: "contributions" is blue on the home
// page, on the registry leaderboard and anywhere else it appears, whatever its
// position in the chart. Never cycle these for a fifth series — fold the tail
// into "other" or facet the chart instead.

export interface ChartSeriesSpec {
  key: string
  label: string
  color: string
}

export const ChartColors = {
  claims: '#ba8b00',
  contributions: '#3280dd',
  ciphertexts: '#137738',
  partials: '#d15697',
} as const

export const SeriesClaims: ChartSeriesSpec = {
  key: 'claims',
  label: 'Slots claimed',
  color: ChartColors.claims,
}
export const SeriesContributions: ChartSeriesSpec = {
  key: 'contributions',
  label: 'Contributions',
  color: ChartColors.contributions,
}
export const SeriesCiphertexts: ChartSeriesSpec = {
  key: 'ciphertexts',
  label: 'Ciphertexts',
  color: ChartColors.ciphertexts,
}
export const SeriesPartials: ChartSeriesSpec = {
  key: 'partials',
  label: 'Partial decryptions',
  color: ChartColors.partials,
}

/** Recessive gridline colour — same hairline the rest of the UI uses. */
export const ChartGrid = 'rgba(236, 232, 223, 0.10)'
/** Axis / tick label colour (matches the `ink.4` semantic token). */
export const ChartAxisInk = '#6b6759'
/** Slightly brighter ink for the axis labels that carry meaning. */
export const ChartLabelInk = '#8d887a'
