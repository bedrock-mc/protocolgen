// Command raw generates a self-contained Go package from Mojang's
// bedrock-protocol-docs JSON: a struct for every packet and every composite
// definition, marshalling against a generated IO interface, with no gophertunnel
// dependency.
//
// The docs are under the Minecraft EULA; pass a local clone path via -docs and do
// not vendor their content.
package main

import (
	"flag"
	"fmt"
	"os"

	"protocolgen/internal/codegen"
)

func main() {
	docs := flag.String("docs", "", "path to the Mojang bedrock-protocol-docs json/ directory (local clone)")
	out := flag.String("out", "out_raw", "directory to write the generated package into")
	pkg := flag.String("pkg", "raw", "package name for the generated output")
	flag.Parse()

	if *docs == "" {
		fmt.Fprintln(os.Stderr, "error: -docs is required")
		flag.Usage()
		os.Exit(2)
	}

	if err := codegen.Generate(codegen.Options{
		DocsDir: *docs,
		OutDir:  *out,
		Format:  "raw",
		RawPkg:  *pkg,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}
}
