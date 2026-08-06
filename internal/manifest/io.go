package manifest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load(path string) (Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("read manifest: %w", err)
	}
	var result Manifest
	if err := json.Unmarshal(data, &result); err != nil {
		return Manifest{}, fmt.Errorf("parse manifest: %w", err)
	}
	if err := Validate(result); err != nil {
		return Manifest{}, fmt.Errorf("validate manifest: %w", err)
	}
	return result, nil
}

func Write(path string, value Manifest) error {
	data, err := MarshalStable(value)
	if err != nil {
		return err
	}
	if parent := filepath.Dir(path); parent != "." && parent != "" {
		if err := os.MkdirAll(parent, 0o755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, data, 0o644)
}

func LoadAdjudications(path string) ([]Adjudication, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read adjudications: %w", err)
	}
	var list []Adjudication
	if err := json.Unmarshal(data, &list); err == nil {
		return list, nil
	}
	var wrapped struct {
		Adjudications []Adjudication `json:"adjudications"`
	}
	if err := json.Unmarshal(data, &wrapped); err != nil {
		return nil, fmt.Errorf("parse adjudications: %w", err)
	}
	return wrapped.Adjudications, nil
}
