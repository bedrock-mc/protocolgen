// Command protocoldrift compares gophertunnel's hand-written packet codecs
// against Mojang's official bedrock-protocol-docs JSON schema, reporting where
// the two have drifted apart (new/removed packets, field-count divergences, and
// wire-encoding mismatches).
//
// It reads the docs from a local clone passed via -docs (Mojang's repo is under
// the Minecraft EULA, so its content is never vendored here) and analyses the
// gophertunnel packets as source via go/ast — it never imports them, so it works
// across protocol versions without compiling the target.
//
// Exit codes: 0 = no drift, 1 = drift found, 2 = error.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	docsDir := flag.String("docs", "", "path to the Mojang bedrock-protocol-docs json/ directory (local clone)")
	packetsDir := flag.String("packets", "minecraft/protocol/packet", "path to gophertunnel's protocol/packet directory")
	verbose := flag.Bool("verbose", false, "show every field, not just findings")
	asJSON := flag.Bool("json", false, "emit the report as JSON")
	failOnDrift := flag.Bool("fail-on-drift", true, "exit non-zero when drift is detected")
	flag.Parse()

	if *docsDir == "" {
		fmt.Fprintln(os.Stderr, "error: -docs is required (path to a local clone's json/ directory)")
		flag.Usage()
		os.Exit(2)
	}

	docs, err := loadDocPackets(*docsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}
	srcs, err := loadSourcePackets(*packetsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	rep := Compare(docs, srcs)

	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rep); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(2)
		}
	} else {
		fmt.Print(render(rep, *verbose))
	}

	if *failOnDrift && rep.HasDrift() {
		os.Exit(1)
	}
}
