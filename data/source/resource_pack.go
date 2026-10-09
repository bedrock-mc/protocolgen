package source

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// ResourcePack returns the resource subtree of the locked official sample pack.
// Every cached entry is checked against the authenticated archive before reuse.
// Passing an empty cache selects the user's protocolgen-data cache directory.
func (s *Sources) ResourcePack(ctx context.Context, cache string) (string, error) {
	input, ok := s.Lock.Inputs["resource_pack"]
	if !ok {
		return "", fmt.Errorf("source lock has no resource_pack input")
	}
	if cache == "" {
		base, err := os.UserCacheDir()
		if err != nil {
			return "", err
		}
		cache = filepath.Join(base, "protocolgen-data")
	}
	archive, err := s.Read(ctx, "resource_pack", cache)
	if err != nil {
		return "", err
	}
	expected, err := extractResourcePack(bytes.NewReader(archive), "")
	if err != nil {
		return "", err
	}
	manifest, ok := expected[filepath.Join("resource_pack", "manifest.json")]
	if !ok || manifest.directory {
		return "", fmt.Errorf("locked pack has no resource_pack/manifest.json")
	}
	dir := filepath.Join(cache, "packs", input.SHA256)
	root := filepath.Join(dir, "resource_pack")
	if _, err := os.Lstat(dir); err == nil {
		if err := verifyResourcePack(dir, expected); err != nil {
			return "", fmt.Errorf("cached resource pack is invalid: %w; remove %q and retry", err, dir)
		}
		return root, nil
	} else if !os.IsNotExist(err) {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(dir), 0o755); err != nil {
		return "", err
	}
	staging, err := os.MkdirTemp(filepath.Dir(dir), ".pack-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(staging)
	if _, err := extractResourcePack(bytes.NewReader(archive), staging); err != nil {
		return "", err
	}
	if err := os.Rename(staging, dir); err != nil {
		// A concurrent caller may have installed the same archive. Accept its tree
		// only after checking the same complete manifest used for ordinary reuse.
		if validationErr := verifyResourcePack(dir, expected); validationErr != nil {
			return "", fmt.Errorf("install resource pack: %w; cached tree is invalid: %v; remove %q and retry", err, validationErr, dir)
		}
	}
	return root, nil
}

// resourcePackEntry records the exact type and contents of an extracted entry.
type resourcePackEntry struct {
	directory bool
	size      int64
	digest    [sha256.Size]byte
}

// extractResourcePack reads and hashes the resource subtree, rejecting links,
// traversal and oversized archives. An empty destination only builds its manifest.
func extractResourcePack(input io.Reader, destination string) (map[string]resourcePackEntry, error) {
	gz, err := gzip.NewReader(input)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	reader := tar.NewReader(gz)
	expected := make(map[string]resourcePackEntry)
	var total int64
	for {
		header, err := reader.Next()
		if err == io.EOF {
			return expected, nil
		}
		if err != nil {
			return nil, err
		}
		_, relative, found := strings.Cut(header.Name, "/")
		if !found || !strings.HasPrefix(relative, "resource_pack/") {
			continue
		}
		if !filepath.IsLocal(relative) || strings.Contains(relative, "\\") || strings.Contains(relative, "/../") || strings.HasSuffix(relative, "/..") {
			return nil, fmt.Errorf("resource entry %q escapes the destination", header.Name)
		}
		relative = filepath.Clean(filepath.FromSlash(relative))
		target := filepath.Join(destination, relative)
		for parent := filepath.Dir(relative); parent != "."; parent = filepath.Dir(parent) {
			if previous, exists := expected[parent]; exists && !previous.directory {
				return nil, fmt.Errorf("resource entry %q has a file as its parent", header.Name)
			}
			expected[parent] = resourcePackEntry{directory: true}
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if previous, exists := expected[relative]; exists && !previous.directory {
				return nil, fmt.Errorf("resource directory %q conflicts with a file", header.Name)
			}
			expected[relative] = resourcePackEntry{directory: true}
			if destination != "" {
				if err := os.MkdirAll(target, 0o755); err != nil {
					return nil, err
				}
			}
		case tar.TypeReg:
			if _, exists := expected[relative]; exists {
				return nil, fmt.Errorf("duplicate resource file %q", header.Name)
			}
			if header.Size < 0 || header.Size > (2<<30)-total {
				return nil, fmt.Errorf("resource pack exceeds the expanded size limit")
			}
			total += header.Size
			hash := sha256.New()
			var file *os.File
			var output io.Writer = hash
			if destination != "" {
				if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
					return nil, err
				}
				file, err = os.Create(target)
				if err != nil {
					return nil, err
				}
				output = io.MultiWriter(file, hash)
			}
			_, copyErr := io.Copy(output, reader)
			var closeErr error
			if file != nil {
				closeErr = file.Close()
			}
			if copyErr != nil {
				return nil, copyErr
			}
			if closeErr != nil {
				return nil, closeErr
			}
			expected[relative] = resourcePackEntry{size: header.Size, digest: [sha256.Size]byte(hash.Sum(nil))}
		default:
			return nil, fmt.Errorf("unsupported resource entry type %d for %q", header.Typeflag, header.Name)
		}
	}
}

// verifyResourcePack rejects missing, additional, changed and non-regular entries.
func verifyResourcePack(dir string, expected map[string]resourcePackEntry) error {
	seen := 0
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if relative == "." {
			if !entry.IsDir() {
				return fmt.Errorf("cache root is not a directory")
			}
			return nil
		}
		want, ok := expected[relative]
		if !ok {
			return fmt.Errorf("unexpected entry %q", relative)
		}
		if entry.IsDir() != want.directory || (!entry.IsDir() && !entry.Type().IsRegular()) {
			return fmt.Errorf("entry %q has the wrong file type", relative)
		}
		seen++
		if want.directory {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		hash := sha256.New()
		size, readErr := io.Copy(hash, io.LimitReader(file, want.size+1))
		closeErr := file.Close()
		if readErr != nil {
			return readErr
		}
		if closeErr != nil {
			return closeErr
		}
		if size != want.size || [sha256.Size]byte(hash.Sum(nil)) != want.digest {
			return fmt.Errorf("entry %q differs from the locked archive", relative)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if seen != len(expected) {
		return fmt.Errorf("cache has %d entries, want %d", seen, len(expected))
	}
	return nil
}
