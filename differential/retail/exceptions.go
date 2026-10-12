package retail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

// AnyFixture matches every fixture of an exception's decoder, case and kind.
const AnyFixture = "*"

// Exception accepts the disagreements of one decoder, case, fixture and kind.
type Exception struct {
	Decoder  string   `json:"decoder"`
	Case     string   `json:"case"`
	Fixture  string   `json:"fixture"`
	Kind     string   `json:"kind"`
	Reason   string   `json:"reason"`
	Evidence []string `json:"evidence"`
}

// ExceptionFile is differential/retail/exceptions.json.
type ExceptionFile struct {
	SchemaVersion    int         `json:"schema_version"`
	MinecraftVersion string      `json:"minecraft_version"`
	ProtocolVersion  int         `json:"protocol_version"`
	Exceptions       []Exception `json:"exceptions"`
}

// LoadExceptions reads an exception file for corpus, rejecting entries
// without a reason and evidence, with unknown names, or repeated.
func LoadExceptions(path string, corpus Corpus) ([]Exception, error) {
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
		key := [4]string{e.Decoder, e.Case, e.Fixture, e.Kind}
		switch {
		case !slices.Contains(Decoders, e.Decoder):
			return nil, fmt.Errorf("%s: exception %d names decoder %q", path, i, e.Decoder)
		case rules[e.Case].decoders == nil:
			return nil, fmt.Errorf("%s: exception %d names case %q", path, i, e.Case)
		case !slices.Contains(kinds, e.Kind):
			return nil, fmt.Errorf("%s: exception %d names kind %q", path, i, e.Kind)
		case e.Fixture == "":
			return nil, fmt.Errorf("%s: exception %d names no fixture", path, i)
		case strings.TrimSpace(e.Reason) == "" || len(e.Evidence) == 0 || slices.Contains(e.Evidence, ""):
			return nil, fmt.Errorf("%s: exception %d lacks a reason or evidence", path, i)
		case seen[key]:
			return nil, fmt.Errorf("%s: exception %d repeats %v", path, i, key)
		}
		seen[key] = true
	}
	return file.Exceptions, nil
}

func (e Exception) matches(d Disagreement) bool {
	return e.Decoder == d.Decoder && e.Case == d.Case && e.Kind == d.Kind && (e.Fixture == AnyFixture || e.Fixture == d.Fixture)
}

// Report splits disagreements by whether an exception accepts them, and lists
// exceptions that accept nothing.
type Report struct {
	Unexpected []Disagreement
	Accepted   []Disagreement
	Stale      []Exception
}

// Failed reports whether the comparison must fail: an unaccepted
// disagreement or a stale exception.
func (r Report) Failed() bool {
	return len(r.Unexpected) != 0 || len(r.Stale) != 0
}

// Apply matches disagreements against exceptions.
func Apply(found []Disagreement, exceptions []Exception) Report {
	var report Report
	used := make([]bool, len(exceptions))
	for _, d := range found {
		accepted := false
		for i, e := range exceptions {
			if e.matches(d) {
				used[i], accepted = true, true
			}
		}
		if accepted {
			report.Accepted = append(report.Accepted, d)
		} else {
			report.Unexpected = append(report.Unexpected, d)
		}
	}
	for i, e := range exceptions {
		if !used[i] {
			report.Stale = append(report.Stale, e)
		}
	}
	return report
}
