// Package emitter provides the common orchestration contract for language
// backends. Backends receive the same validated manifest and reviewed overlays,
// then return a deterministic set of relative output files.
package emitter

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"protocolgen/internal/docs"
	"protocolgen/internal/domains"
	"protocolgen/internal/manifest"
	"protocolgen/internal/naming"
)

type Input struct {
	Manifest manifest.Manifest
	Naming   naming.Overlay
	Domains  domains.Overlay
	Docs     docs.Overlay
}

type Backend interface {
	Generate(Input) (map[string]string, error)
}

type Func func(Input) (map[string]string, error)

func (f Func) Generate(input Input) (map[string]string, error) { return f(input) }

type Config struct {
	ManifestPath string
	NamingPath   string
	DomainsPath  string
	DocsPath     string
	OutputDir    string
}

type Result struct {
	FileCount int
	Coverage  docs.Coverage
}

// Run loads and validates the common generator inputs, invokes one language
// backend, and safely replaces its generated output files.
func Run(config Config, backend Backend) (Result, error) {
	if config.ManifestPath == "" || config.OutputDir == "" {
		return Result{}, fmt.Errorf("manifest path and output directory are required")
	}
	if backend == nil {
		return Result{}, fmt.Errorf("emitter backend is required")
	}
	m, err := manifest.Load(config.ManifestPath)
	if err != nil {
		return Result{}, err
	}
	namingOverlay, err := loadOptionalOverlay(config.ManifestPath, config.NamingPath, "naming.json", "naming", func(path string) (naming.Overlay, error) {
		return naming.LoadOverlay(path, m)
	})
	if err != nil {
		return Result{}, err
	}
	domainOverlay, err := loadOptionalOverlay(config.ManifestPath, config.DomainsPath, "domains.json", "domains", func(path string) (domains.Overlay, error) {
		return domains.LoadOverlay(path, m)
	})
	if err != nil {
		return Result{}, err
	}
	docOverlay, err := loadOptionalOverlay(config.ManifestPath, config.DocsPath, "docs.json", "docs", func(path string) (docs.Overlay, error) {
		return docs.LoadOverlay(path, m)
	})
	if err != nil {
		return Result{}, err
	}
	files, err := backend.Generate(Input{Manifest: m, Naming: namingOverlay, Domains: domainOverlay, Docs: docOverlay})
	if err != nil {
		return Result{}, err
	}
	if err := WriteFiles(config.OutputDir, files); err != nil {
		return Result{}, err
	}
	return Result{FileCount: len(files), Coverage: docs.CoverageOf(m, docOverlay)}, nil
}

func loadOptionalOverlay[T any](manifestPath, explicitPath, defaultName, label string, load func(string) (T, error)) (T, error) {
	path := explicitPath
	if path == "" {
		path = filepath.Join(filepath.Dir(manifestPath), defaultName)
	}
	if _, err := os.Stat(path); err != nil {
		var zero T
		if os.IsNotExist(err) && explicitPath == "" {
			return zero, nil
		}
		return zero, fmt.Errorf("stat %s overlay: %w", label, err)
	}
	return load(path)
}

func WriteFiles(directory string, files map[string]string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return err
	}
	desired := make(map[string]bool, len(files))
	names := make([]string, 0, len(files))
	for name := range files {
		clean := filepath.Clean(name)
		if clean == "." || filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
			return fmt.Errorf("emitter returned unsafe filename %q", name)
		}
		desired[clean] = true
		names = append(names, name)
	}
	if err := removeStaleGeneratedFiles(directory, desired); err != nil {
		return err
	}
	sort.Strings(names)
	for _, name := range names {
		path := filepath.Join(directory, filepath.Clean(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(files[name]), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func removeStaleGeneratedFiles(directory string, desired map[string]bool) error {
	const generatedHeader = "Code generated from canonical protocol manifest v2. DO NOT EDIT."
	return filepath.WalkDir(directory, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(directory, path)
		if err != nil || desired[relative] {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		contents := strings.TrimLeft(string(data), "/# ")
		if strings.HasPrefix(contents, generatedHeader) {
			return os.Remove(path)
		}
		return nil
	})
}
