package retail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

// AnyFixture, as an exception's only fixture, matches every fixture of its
// decoder, case and kind.
const AnyFixture = "*"

// Exception accepts one decoder's disagreements of one kind and case on the
// listed fixtures, which share one cause. Every listed fixture must still
// disagree, or the exception is stale.
type Exception struct {
	Decoder  string   `json:"decoder"`
	Case     string   `json:"case"`
	Fixtures []string `json:"fixtures"`
	Kind     string   `json:"kind"`
	Reason   string   `json:"reason"`
	Evidence []string `json:"evidence"`
	// Oracle pins an exception caused by gophertunnel to the module version
	// this harness links; it is required for gophertunnel's own verdicts.
	Oracle string `json:"oracle,omitempty"`
}

// ExceptionFile is differential/retail/exceptions.json.
type ExceptionFile struct {
	SchemaVersion    int         `json:"schema_version"`
	MinecraftVersion string      `json:"minecraft_version"`
	ProtocolVersion  int         `json:"protocol_version"`
	Exceptions       []Exception `json:"exceptions"`
}

// LoadExceptions reads an exception file for corpus and the linked
// gophertunnel ("path version"). It rejects entries without a reason and
// evidence, with unknown names, repeated, excusing a panic, or pinned to
// another oracle.
func LoadExceptions(path string, corpus Corpus, oracle string) ([]Exception, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var file ExceptionFile
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&file); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if file.SchemaVersion != 1 || file.MinecraftVersion != corpus.MinecraftVersion || file.ProtocolVersion != corpus.ProtocolVersion {
		return nil, fmt.Errorf("%s does not target corpus %s/%d", path, corpus.MinecraftVersion, corpus.ProtocolVersion)
	}
	seen := map[[4]string]bool{}
	for i, e := range file.Exceptions {
		fail := func(problem string) error { return fmt.Errorf("%s: exception %d %s", path, i, problem) }
		switch {
		case !slices.Contains(Decoders, e.Decoder):
			return nil, fail(fmt.Sprintf("names decoder %q", e.Decoder))
		case rules[e.Case].decoders == nil:
			return nil, fail(fmt.Sprintf("names case %q", e.Case))
		case !slices.Contains(kinds, e.Kind) || e.Kind == KindPanic:
			return nil, fail(fmt.Sprintf("names kind %q, which is unknown or never acceptable", e.Kind))
		case len(e.Fixtures) == 0 || slices.Contains(e.Fixtures, "") || len(e.Fixtures) > 1 && slices.Contains(e.Fixtures, AnyFixture):
			return nil, fail("needs fixtures, or * alone")
		case strings.TrimSpace(e.Reason) == "" || len(e.Evidence) == 0 || slices.Contains(e.Evidence, ""):
			return nil, fail("lacks a reason or evidence")
		case e.Decoder == DecoderGophertunnel && e.Oracle == "":
			return nil, fail("excuses gophertunnel without pinning the oracle")
		case e.Oracle != "" && e.Oracle != oracle:
			return nil, fail(fmt.Sprintf("is pinned to %s, the harness links %s", e.Oracle, oracle))
		}
		for _, fixture := range e.Fixtures {
			key := [4]string{e.Decoder, e.Case, fixture, e.Kind}
			if seen[key] {
				return nil, fail(fmt.Sprintf("repeats %v", key))
			}
			seen[key] = true
		}
	}
	return file.Exceptions, nil
}

func (e Exception) matches(d Disagreement, fixture string) bool {
	return e.Decoder == d.Decoder && e.Case == d.Case && e.Kind == d.Kind && (fixture == AnyFixture || fixture == d.Fixture)
}

// Stale is an exception fixture that no disagreement matched.
type Stale struct {
	Exception Exception
	Fixture   string
}

// Report splits disagreements by whether an exception accepts them, and lists
// exception fixtures that accept nothing.
type Report struct {
	Unexpected []Disagreement
	Accepted   []Disagreement
	Stale      []Stale
}

// Failed reports whether the comparison must fail: an unaccepted
// disagreement or a stale exception.
func (r Report) Failed() bool {
	return len(r.Unexpected) != 0 || len(r.Stale) != 0
}

// Apply matches disagreements against exceptions.
func Apply(found []Disagreement, exceptions []Exception) Report {
	var report Report
	used := map[[2]int]bool{} // exception index, fixture index
	for _, d := range found {
		accepted := false
		for i, e := range exceptions {
			for j, fixture := range e.Fixtures {
				if e.matches(d, fixture) {
					used[[2]int{i, j}], accepted = true, true
				}
			}
		}
		if accepted {
			report.Accepted = append(report.Accepted, d)
		} else {
			report.Unexpected = append(report.Unexpected, d)
		}
	}
	for i, e := range exceptions {
		for j, fixture := range e.Fixtures {
			if !used[[2]int{i, j}] {
				report.Stale = append(report.Stale, Stale{e, fixture})
			}
		}
	}
	return report
}
