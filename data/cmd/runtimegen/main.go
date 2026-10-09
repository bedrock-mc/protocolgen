// Command runtimegen generates inventory data from the shared source lock.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bedrock-mc/protocolgen/data/internal/generator"
	"github.com/bedrock-mc/protocolgen/data/internal/runtimegen"
	"github.com/bedrock-mc/protocolgen/data/source"
)

func main() {
	lock := flag.String("lock", "", "source lock path; defaults to the bundled release")
	out := flag.String("out", "../generated/data", "generated catalog root")
	cacheRoot, _ := os.UserCacheDir()
	cache := flag.String("cache", filepath.Join(cacheRoot, "protocolgen-data"), "authenticated input cache")
	flag.Parse()
	if err := run(*lock, *out, *cache); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run verifies the release inputs and writes a fully validated catalog.
func run(lock, output, cache string) error {
	sources, err := source.Open(lock)
	if err != nil {
		return err
	}
	data, err := runtimegen.FromSources(context.Background(), sources, cache)
	if err != nil {
		return err
	}
	files, err := runtimegen.Release(sources.Lock)
	if err != nil {
		return err
	}
	files["item/runtime_generated.go"] = data
	return generator.Write(output, runtimegen.GeneratedMarker, files)
}
