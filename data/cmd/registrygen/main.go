// Command registrygen generates consumer registry payloads from authenticated public inputs.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bedrock-mc/protocolgen/data/internal/generator"
	"github.com/bedrock-mc/protocolgen/data/internal/registrygen"
	"github.com/bedrock-mc/protocolgen/data/source"
)

// main resolves the source lock and writes all registry outputs together.
func main() {
	lock := flag.String("lock", "", "source lock path; defaults to the bundled release")
	output := flag.String("out", "../generated/data", "generated catalog root")
	cacheRoot, _ := os.UserCacheDir()
	cache := flag.String("cache", filepath.Join(cacheRoot, "protocolgen-data"), "authenticated input cache")
	flag.Parse()
	if err := run(*lock, *output, *cache); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run authenticates the release and writes the normalized registry projection.
func run(lock, output, cache string) error {
	sources, err := source.Open(lock)
	if err != nil {
		return err
	}
	files, err := registrygen.FromSources(context.Background(), sources, cache)
	if err != nil {
		return err
	}
	return generator.Write(output, registrygen.GeneratedMarker, files)
}
