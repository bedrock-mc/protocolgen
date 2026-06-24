// Command gophertunnel generates gophertunnel-format packet codecs from Mojang's
// bedrock-protocol-docs JSON (mapping onto gophertunnel's protocol.* types), and
// — with -drift — reports where the docs and an existing packet directory differ.
//
// The docs are under the Minecraft EULA; pass a local clone path via -docs and do
// not vendor their content.
package main

import (
	"flag"
	"fmt"
	"os"

	"protocolgen/internal/codegen"
	"protocolgen/internal/driftcheck"
)

func main() {
	docs := flag.String("docs", "", "path to the Mojang bedrock-protocol-docs json/ directory (local clone)")
	packets := flag.String("packets", "minecraft/protocol/packet", "gophertunnel packet directory (id.go + overwrite/drift target)")
	out := flag.String("out", "out_gophertunnel", "directory to write generated files into")
	overwrite := flag.Bool("overwrite", false, "overwrite the packet files in -packets in place (matched by packet id) for git-diff review")
	preserve := flag.Bool("preserve-matching", true, "overwrite: keep the existing field name/type/op when the wire encoding is unchanged (skips pure renames)")
	drift := flag.Bool("drift", false, "report drift between the docs and -packets instead of generating")
	verbose := flag.Bool("verbose", false, "-drift: show every field, not just findings")
	asJSON := flag.Bool("json", false, "-drift: emit the report as JSON")
	flag.Parse()

	if *docs == "" {
		fmt.Fprintln(os.Stderr, "error: -docs is required")
		flag.Usage()
		os.Exit(2)
	}

	if *drift {
		found, err := driftcheck.Check(*docs, *packets, *verbose, *asJSON)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(2)
		}
		if found {
			os.Exit(1)
		}
		return
	}

	if err := codegen.Generate(codegen.Options{
		DocsDir:          *docs,
		PacketsDir:       *packets,
		OutDir:           *out,
		Format:           "gophertunnel",
		Overwrite:        *overwrite,
		PreserveMatching: *preserve,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}
}
