// Command mininggen emits provisional mining facts from the pinned public input.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bedrock-mc/protocolgen/data/internal/mininggen"
	"github.com/bedrock-mc/protocolgen/data/source"
)

func main() {
	lockPath := flag.String("lock", "", "source lock path; defaults to the bundled release")
	outputDir := flag.String("out", "../generated/data", "generated catalog root")
	cacheRoot, _ := os.UserCacheDir()
	cache := flag.String("cache", filepath.Join(cacheRoot, "protocolgen-data"), "authenticated input cache")
	flag.Parse()
	if err := run(*lockPath, *cache, *outputDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run authenticates the mining input and atomically replaces its one output file.
func run(lockPath, cache, outputDir string) error {
	sources, err := source.Open(lockPath)
	if err != nil {
		return err
	}
	input, ok := sources.Lock.Inputs["mining_tools"]
	if !ok || input.Upstream != "prismarine" || input.File != "data/bedrock/1.26.30/blocks.json" {
		return fmt.Errorf("mining_tools must pin Prismarine Bedrock 1.26.30 blocks.json")
	}
	data, err := sources.Read(context.Background(), "mining_tools", cache)
	if err != nil {
		return err
	}
	lockSHA, err := sources.Lock.SHA256()
	if err != nil {
		return err
	}
	generated, count, err := mininggen.Generate(data, sources.Revision("mining_tools"), input.SHA256, lockSHA)
	if err != nil {
		return err
	}
	path := filepath.Join(outputDir, "block", "mining_generated.go")
	file, err := os.CreateTemp(filepath.Dir(path), ".mining-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(generated); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(file.Name(), path); err != nil {
		return err
	}
	fmt.Printf("generated %d provisional mining classifications\n", count)
	return nil
}
