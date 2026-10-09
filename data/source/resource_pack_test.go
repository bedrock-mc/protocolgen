package source

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// sampleArchive creates a synthetic sample-pack archive for extraction tests.
func sampleArchive(t *testing.T, entries []tar.Header) []byte {
	t.Helper()
	var output bytes.Buffer
	gz := gzip.NewWriter(&output)
	tarWriter := tar.NewWriter(gz)
	for _, header := range entries {
		if header.Typeflag == tar.TypeReg {
			header.Size = 2
		}
		if err := tarWriter.WriteHeader(&header); err != nil {
			t.Fatal(err)
		}
		if header.Typeflag == tar.TypeReg {
			if _, err := tarWriter.Write([]byte("{}")); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := tarWriter.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	return output.Bytes()
}

func TestResourcePackExtractsOnlyResourcesAfterAuthentication(t *testing.T) {
	archive := sampleArchive(t, []tar.Header{
		{Name: "samples/resource_pack/manifest.json", Typeflag: tar.TypeReg},
		{Name: "samples/resource_pack/textures/test.json", Typeflag: tar.TypeReg},
		{Name: "samples/behavior_pack/entity.json", Typeflag: tar.TypeReg},
	})
	dir, cache := t.TempDir(), t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "pack.tar.gz"), archive, 0o644); err != nil {
		t.Fatal(err)
	}
	sources := &Sources{dir: dir, Lock: Lock{Inputs: map[string]Input{
		"resource_pack": {Path: "pack.tar.gz", SHA256: fmt.Sprintf("%x", sha256.Sum256(archive))},
	}}}
	root, err := sources.ResourcePack(context.Background(), cache)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "textures", "test.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(root), "behavior_pack")); !os.IsNotExist(err) {
		t.Fatalf("non-resource files were extracted: %v", err)
	}
	if reused, err := sources.ResourcePack(context.Background(), cache); err != nil || reused != root {
		t.Fatalf("cache reuse = %q, %v", reused, err)
	}
	// An unauthenticated archive must never be extracted into a fresh cache.
	if err := os.WriteFile(filepath.Join(dir, "pack.tar.gz"), []byte("changed"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := sources.ResourcePack(context.Background(), t.TempDir()); err == nil {
		t.Fatal("changed archive was accepted")
	}
}

func TestResourcePackRejectsEscapesAndLinks(t *testing.T) {
	for _, header := range []tar.Header{
		{Name: "samples/resource_pack/../../outside", Typeflag: tar.TypeReg},
		{Name: "samples/resource_pack/..", Typeflag: tar.TypeDir},
		{Name: "samples/resource_pack/link", Typeflag: tar.TypeSymlink, Linkname: "../../outside"},
		{Name: "samples/resource_pack/link", Typeflag: tar.TypeLink, Linkname: "../../outside"},
	} {
		t.Run(header.Name+fmt.Sprint(header.Typeflag), func(t *testing.T) {
			if _, err := extractResourcePack(bytes.NewReader(sampleArchive(t, []tar.Header{header})), t.TempDir()); err == nil {
				t.Fatal("unsafe archive entry was accepted")
			}
		})
	}
}

func TestResourcePackRejectsChangedCachedTree(t *testing.T) {
	for _, test := range []struct {
		name   string
		change func(string) error
	}{
		{"changed file", func(root string) error {
			return os.WriteFile(filepath.Join(root, "textures", "test.json"), []byte("[]"), 0o644)
		}},
		{"deleted file", func(root string) error { return os.Remove(filepath.Join(root, "textures", "test.json")) }},
		{"extra file", func(root string) error { return os.WriteFile(filepath.Join(root, "extra.json"), []byte("{}"), 0o644) }},
		{"extra directory", func(root string) error { return os.Mkdir(filepath.Join(root, "extra"), 0o755) }},
		{"missing manifest", func(root string) error { return os.Remove(filepath.Join(root, "manifest.json")) }},
		{"symlink", func(root string) error {
			path := filepath.Join(root, "textures", "test.json")
			if err := os.Remove(path); err != nil {
				return err
			}
			return os.Symlink(filepath.Join(root, "manifest.json"), path)
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			sources, cache := resourcePackFixture(t)
			root, err := sources.ResourcePack(context.Background(), cache)
			if err != nil {
				t.Fatal(err)
			}
			if err := test.change(root); err != nil {
				t.Fatal(err)
			}
			for range 2 {
				got, err := sources.ResourcePack(context.Background(), cache)
				if err == nil || got != "" || !strings.Contains(err.Error(), filepath.Dir(root)) || !strings.Contains(err.Error(), "remove") {
					t.Fatalf("invalid cache returned %q, %v; want an error naming the directory to remove", got, err)
				}
			}
		})
	}
}

func TestResourcePackConcurrentInstallation(t *testing.T) {
	sources, cache := resourcePackFixture(t)
	const callers = 12
	results := make(chan string, callers)
	errors := make(chan error, callers)
	start := make(chan struct{})
	var workers sync.WaitGroup
	for range callers {
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			root, err := sources.ResourcePack(context.Background(), cache)
			results <- root
			errors <- err
		}()
	}
	close(start)
	workers.Wait()
	close(results)
	close(errors)
	for err := range errors {
		if err != nil {
			t.Fatal(err)
		}
	}
	var expectedRoot string
	for root := range results {
		if expectedRoot == "" {
			expectedRoot = root
		}
		if root != expectedRoot {
			t.Fatalf("concurrent calls returned different roots: %q and %q", root, expectedRoot)
		}
	}
	if payload, err := os.ReadFile(filepath.Join(expectedRoot, "textures", "test.json")); err != nil || string(payload) != "{}" {
		t.Fatalf("installed texture = %q, %v", payload, err)
	}
}

// resourcePackFixture supplies an authenticated local archive and an empty cache.
func resourcePackFixture(t *testing.T) (*Sources, string) {
	t.Helper()
	archive := sampleArchive(t, []tar.Header{
		{Name: "samples/resource_pack/manifest.json", Typeflag: tar.TypeReg},
		{Name: "samples/resource_pack/textures/test.json", Typeflag: tar.TypeReg},
	})
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "pack.tar.gz"), archive, 0o644); err != nil {
		t.Fatal(err)
	}
	return &Sources{dir: dir, Lock: Lock{Inputs: map[string]Input{
		"resource_pack": {Path: "pack.tar.gz", SHA256: fmt.Sprintf("%x", sha256.Sum256(archive))},
	}}}, t.TempDir()
}
