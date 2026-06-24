// Package driftcheck compares gophertunnel's hand-written packet codecs against
// Mojang's bedrock-protocol-docs JSON and reports where they have drifted apart.
// It reads the docs as JSON and the packets as Go source (go/ast) — it never
// imports the target.
package driftcheck

import (
	"encoding/json"
	"fmt"
	"os"
)

// Check loads the docs and packets, compares them, prints the report, and reports
// whether drift was found.
func Check(docsDir, packetsDir string, verbose, asJSON bool) (drift bool, err error) {
	docs, err := loadDocPackets(docsDir)
	if err != nil {
		return false, err
	}
	srcs, err := loadSourcePackets(packetsDir)
	if err != nil {
		return false, err
	}

	rep := Compare(docs, srcs)

	if asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rep); err != nil {
			return rep.HasDrift(), err
		}
	} else {
		fmt.Print(render(rep, verbose))
	}
	return rep.HasDrift(), nil
}
