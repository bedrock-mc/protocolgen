// Command generate converts pinned Cloudburst and vanilla BDS inputs into the
// stable Go packages in this module.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bedrock-mc/protocolgen/data/internal/generator"
	"github.com/bedrock-mc/protocolgen/data/source"
)

func main() {
	lock := flag.String("lock", "", "source lock path; defaults to the bundled release")
	cloudburst := flag.String("cloudburst", "", "path to the locked CloudburstMC/Data files")
	bds := flag.String("bds", "", "path to the locked Bedrock Dedicated Server inputs")
	output := flag.String("out", ".", "data module root to write")
	cacheRoot, _ := os.UserCacheDir()
	cache := flag.String("cache", filepath.Join(cacheRoot, "protocolgen-data"), "authenticated input cache")
	flag.Parse()

	stats, err := run(*lock, *cache, *cloudburst, *bds, *output)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf(
		"generated %d block states, %d biomes, %d voxel shapes, %d entities, and %d foods\n",
		stats.BlockStates, stats.Biomes, stats.VoxelShapes, stats.Entities, stats.Foods,
	)
}

// run authenticates every input and derives generation versions from the lock.
func run(lockPath, cache, cloudburstDir, bdsDir, outputDir string) (generator.Stats, error) {
	sources, err := source.Open(lockPath)
	if err != nil {
		return generator.Stats{}, err
	}
	if err := sources.ValidateSemantic(cloudburstDir, bdsDir); err != nil {
		return generator.Stats{}, err
	}
	shapes, err := sources.Read(context.Background(), "block_shapes", cache)
	if err != nil {
		return generator.Stats{}, err
	}
	file, err := os.CreateTemp("", "protocolgen-block-shapes-*.json")
	if err != nil {
		return generator.Stats{}, err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(shapes); err != nil {
		_ = file.Close()
		return generator.Stats{}, err
	}
	if err := file.Close(); err != nil {
		return generator.Stats{}, err
	}
	files, stats, err := generator.Generate(generator.Config{
		CloudburstDir:   cloudburstDir,
		BDSDir:          bdsDir,
		CloudburstRef:   sources.Lock.Semantic.CloudburstRef,
		BDSVersion:      sources.Lock.Semantic.BDSVersion,
		BlockShapesPath: file.Name(),
		BlockShapesRef:  sources.Lock.Inputs["block_shapes"].Revision,
	})
	if err != nil {
		return generator.Stats{}, err
	}
	return stats, generator.Write(outputDir, files)
}
