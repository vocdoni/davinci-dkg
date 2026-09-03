package battery

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

// Result is one structured observation of the battery: a transaction that
// was sent (with its gas and latency), a revert that was expected, or a
// scenario-level measurement such as a throughput figure.
type Result struct {
	Scenario       string  `json:"scenario"`
	Step           string  `json:"step"`
	Kind           string  `json:"kind"` // tx kind (submitCiphertext, ...) or "measure" / "revert"
	Tx             string  `json:"tx,omitempty"`
	Gas            uint64  `json:"gas,omitempty"`
	Block          uint64  `json:"block,omitempty"`
	LatencyBlocks  int64   `json:"latencyBlocks,omitempty"`
	LatencySeconds float64 `json:"latencySeconds,omitempty"`
	Pass           bool    `json:"pass"`
	Notes          string  `json:"notes,omitempty"`
	Time           string  `json:"time"`
}

// reporter appends results to the JSON report on every record so a run that
// is killed half-way still leaves its observations on disk.
type reporter struct {
	mu      sync.Mutex
	path    string
	results []Result
}

var report = newReporter()

func newReporter() *reporter {
	path := os.Getenv("DAVINCI_DKG_BATTERY_REPORT")
	if path == "" {
		path = "/tmp/battery-report.json"
	}
	return &reporter{path: path}
}

// record stores one result, logs it through t and rewrites the JSON file.
func record(t *testing.T, r Result) {
	t.Helper()
	if r.Time == "" {
		r.Time = time.Now().UTC().Format(time.RFC3339)
	}
	if r.Scenario == "" {
		r.Scenario = t.Name()
	}
	report.mu.Lock()
	report.results = append(report.results, r)
	snapshot := make([]Result, len(report.results))
	copy(snapshot, report.results)
	report.mu.Unlock()

	status := "ok"
	if !r.Pass {
		status = "FAIL"
	}
	t.Logf("[%s] %s/%s %s gas=%d block=%d lat=%db/%.1fs %s",
		status, r.Scenario, r.Step, r.Kind, r.Gas, r.Block, r.LatencyBlocks, r.LatencySeconds, r.Notes)
	if err := writeJSON(report.path, snapshot); err != nil {
		t.Logf("report: write %s: %v", report.path, err)
	}
}

func writeJSON(path string, results []Result) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// markdownPath returns the sibling .md file of the JSON report.
func (r *reporter) markdownPath() string {
	return strings.TrimSuffix(r.path, filepath.Ext(r.path)) + ".md"
}

type gasStat struct {
	count         int
	min, max, sum uint64
}

func (g *gasStat) add(gas uint64) {
	if g.count == 0 || gas < g.min {
		g.min = gas
	}
	if gas > g.max {
		g.max = gas
	}
	g.sum += gas
	g.count++
}

// writeMarkdown renders a human-readable summary: per-scenario pass/fail
// tables and a gas digest per transaction kind.
func (r *reporter) writeMarkdown() error {
	r.mu.Lock()
	results := make([]Result, len(r.results))
	copy(results, r.results)
	r.mu.Unlock()

	var b strings.Builder
	fmt.Fprintf(&b, "# davinci-dkg battery report\n\nGenerated %s — %d observations.\n\n",
		time.Now().UTC().Format(time.RFC3339), len(results))
	writeScenarioTables(&b, results)
	writeGasDigest(&b, results)
	return os.WriteFile(r.markdownPath(), []byte(b.String()), 0o644)
}

func writeScenarioTables(b *strings.Builder, results []Result) {
	byScenario := map[string][]Result{}
	var order []string
	for _, res := range results {
		if _, seen := byScenario[res.Scenario]; !seen {
			order = append(order, res.Scenario)
		}
		byScenario[res.Scenario] = append(byScenario[res.Scenario], res)
	}
	for _, scenario := range order {
		rows := byScenario[scenario]
		pass, fail := 0, 0
		for _, res := range rows {
			if res.Pass {
				pass++
			} else {
				fail++
			}
		}
		fmt.Fprintf(b, "## %s — %d pass, %d fail\n\n", scenario, pass, fail)
		b.WriteString("| status | step | kind | gas | block | latency (blocks / s) | notes |\n")
		b.WriteString("|---|---|---|---|---|---|---|\n")
		for _, res := range rows {
			status := "ok"
			if !res.Pass {
				status = "**FAIL**"
			}
			fmt.Fprintf(b, "| %s | %s | %s | %d | %d | %d / %.1f | %s |\n",
				status, res.Step, res.Kind, res.Gas, res.Block, res.LatencyBlocks, res.LatencySeconds,
				strings.ReplaceAll(res.Notes, "|", "\\|"))
		}
		b.WriteString("\n")
	}
}

func writeGasDigest(b *strings.Builder, results []Result) {
	stats := map[string]*gasStat{}
	for _, res := range results {
		if res.Gas == 0 {
			continue
		}
		st, ok := stats[res.Kind]
		if !ok {
			st = &gasStat{}
			stats[res.Kind] = st
		}
		st.add(res.Gas)
	}
	if len(stats) == 0 {
		return
	}
	kinds := make([]string, 0, len(stats))
	for kind := range stats {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)
	b.WriteString("## Gas per transaction kind\n\n| kind | count | min | avg | max |\n|---|---|---|---|---|\n")
	for _, kind := range kinds {
		st := stats[kind]
		fmt.Fprintf(b, "| %s | %d | %d | %d | %d |\n", kind, st.count, st.min, st.sum/uint64(st.count), st.max)
	}
	b.WriteString("\n")
}
