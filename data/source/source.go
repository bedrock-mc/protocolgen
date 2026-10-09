// Package source authenticates the inputs shared by Bedrock data generators.
package source

import (
	"context"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//go:embed lock.json inputs
var bundled embed.FS

// Input identifies an immutable source file. Exactly one of Path or URL is set.
type Input struct {
	Path        string `json:"path,omitempty"`
	URL         string `json:"url,omitempty"`
	SHA256      string `json:"sha256"`
	Revision    string `json:"revision"`
	Description string `json:"description,omitempty"`
}

// Lock records the target and the independent revisions of its inputs.
type Lock struct {
	SchemaVersion    int              `json:"schema_version"`
	MinecraftVersion string           `json:"minecraft_version"`
	ProtocolVersion  int              `json:"protocol_version"`
	Inputs           map[string]Input `json:"inputs"`
	Semantic         SemanticInputs   `json:"semantic"`
}

// Sources resolves files relative to a lock and verifies every input digest.
type Sources struct {
	Lock Lock
	dir  string
}

// Open reads an explicit lock, or the bundled release lock when path is empty.
func Open(path string) (*Sources, error) {
	var data []byte
	var err error
	s := &Sources{}
	if path == "" {
		data, err = bundled.ReadFile("lock.json")
	} else {
		data, err = os.ReadFile(path)
		s.dir = filepath.Dir(path)
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &s.Lock); err != nil {
		return nil, fmt.Errorf("decode source lock: %w", err)
	}
	if s.Lock.SchemaVersion != 1 || s.Lock.MinecraftVersion == "" || s.Lock.ProtocolVersion <= 0 || len(s.Lock.Inputs) == 0 {
		return nil, fmt.Errorf("source lock has an unsupported schema or incomplete target")
	}
	for name, input := range s.Lock.Inputs {
		digest, err := hex.DecodeString(input.SHA256)
		if name == "" || err != nil || len(digest) != sha256.Size || input.Revision == "" || (input.Path == "") == (input.URL == "") {
			return nil, fmt.Errorf("source %q has incomplete identity", name)
		}
		if input.Path != "" && (!filepath.IsLocal(input.Path) || strings.Contains(input.Path, "\\")) {
			return nil, fmt.Errorf("source %q has an invalid local path", name)
		}
		if input.URL != "" && !strings.HasPrefix(input.URL, "https://") {
			return nil, fmt.Errorf("source %q must use HTTPS", name)
		}
	}
	return s, nil
}

// Read loads one locked source. Downloaded files use a digest-addressed cache;
// cached and local files are authenticated just like fresh downloads.
func (s *Sources) Read(ctx context.Context, name, cache string) ([]byte, error) {
	input, ok := s.Lock.Inputs[name]
	if !ok {
		return nil, fmt.Errorf("source %q is not in the lock", name)
	}
	var data []byte
	var err error
	if input.Path != "" {
		if s.dir == "" {
			data, err = bundled.ReadFile(input.Path)
		} else {
			data, err = os.ReadFile(filepath.Join(s.dir, input.Path))
		}
	} else {
		cachePath := ""
		if cache != "" {
			cachePath = filepath.Join(cache, input.SHA256)
			data, err = os.ReadFile(cachePath)
			if err != nil && !os.IsNotExist(err) {
				return nil, fmt.Errorf("read source cache: %w", err)
			}
		}
		if data == nil {
			request, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, input.URL, nil)
			if reqErr != nil {
				return nil, reqErr
			}
			client := &http.Client{Timeout: 2 * time.Minute}
			response, reqErr := client.Do(request)
			if reqErr != nil {
				return nil, reqErr
			}
			defer response.Body.Close()
			if response.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("fetch source %q: %s", name, response.Status)
			}
			const limit = 256 << 20
			data, err = io.ReadAll(io.LimitReader(response.Body, limit+1))
			if len(data) > limit {
				return nil, fmt.Errorf("source %q exceeds %d bytes", name, limit)
			}
			if err == nil && fmt.Sprintf("%x", sha256.Sum256(data)) == input.SHA256 && cachePath != "" {
				if err := os.MkdirAll(cache, 0o755); err != nil {
					return nil, err
				}
				if err := writeCache(cachePath, data); err != nil {
					return nil, err
				}
			}
		}
	}
	if err != nil {
		return nil, fmt.Errorf("read source %q: %w", name, err)
	}
	if got := fmt.Sprintf("%x", sha256.Sum256(data)); got != input.SHA256 {
		return nil, fmt.Errorf("source %q digest is %s, want %s", name, got, input.SHA256)
	}
	return data, nil
}

// writeCache installs a verified input atomically so concurrent generators never
// observe a partially written cache entry.
func writeCache(path string, data []byte) error {
	file, err := os.CreateTemp(filepath.Dir(path), ".source-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(data); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return os.Rename(file.Name(), path)
}
