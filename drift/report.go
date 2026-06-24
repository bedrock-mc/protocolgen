package main

import (
	"fmt"
	"strings"
)

// render produces a human-readable report. When verbose is true, every field is
// shown; otherwise only non-ok findings are listed.
func render(rep Report, verbose bool) string {
	var b strings.Builder
	var mismatches, reviews, missing, extra int

	if len(rep.MissingInSource) > 0 {
		fmt.Fprintf(&b, "NEW packets (in docs, not in gophertunnel):\n")
		for _, d := range rep.MissingInSource {
			fmt.Fprintf(&b, "  + %s (id %d)\n", d.Name, d.ID)
		}
		b.WriteByte('\n')
	}
	if len(rep.MissingInDocs) > 0 {
		fmt.Fprintf(&b, "REMOVED packets (in gophertunnel, not in docs):\n")
		for _, s := range rep.MissingInDocs {
			fmt.Fprintf(&b, "  - %s (%s, id %d)\n", s.TypeName, s.IDConst, s.ID)
		}
		b.WriteByte('\n')
	}

	for _, d := range rep.Diffs {
		var rows []FieldDiff
		for _, f := range d.Fields {
			switch f.Status {
			case statusMismatch:
				mismatches++
			case statusReview:
				reviews++
			case statusMissingSrc:
				missing++
			case statusExtraSrc:
				extra++
			}
			if verbose || f.Status != statusOK {
				rows = append(rows, f)
			}
		}
		if len(rows) == 0 {
			continue
		}
		fmt.Fprintf(&b, "%s ↔ %s (id %d)\n", d.DocName, d.SrcName, d.ID)
		for _, f := range rows {
			guard := ""
			if f.Guarded {
				guard = " [guarded]"
			}
			fmt.Fprintf(&b, "  [%-17s] #%d  doc:%s=%s  src:%s=%s%s\n",
				f.Status, f.Ordinal, fieldOrDash(f.DocName), fieldOrDash(f.DocType),
				fieldOrDash(f.SrcOp), fieldOrDash(f.SrcField), guard)
			if f.Note != "" {
				fmt.Fprintf(&b, "      %s\n", f.Note)
			}
		}
		b.WriteByte('\n')
	}

	fmt.Fprintf(&b, "Summary: %d packets compared, %d new, %d removed | %d mismatch, %d missing, %d extra, %d review\n",
		len(rep.Diffs), len(rep.MissingInSource), len(rep.MissingInDocs), mismatches, missing, extra, reviews)
	if !rep.HasDrift() {
		b.WriteString("Result: no drift detected.\n")
	} else {
		b.WriteString("Result: drift detected — review the findings above.\n")
	}
	return b.String()
}

func fieldOrDash(s string) string {
	if s == "" {
		return "—"
	}
	return s
}
