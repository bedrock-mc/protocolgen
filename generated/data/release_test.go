package data_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bedrock-mc/protocolgen/generated/data"
)

// TestReleaseMetadata verifies that both projections used the same complete lock.
func TestReleaseMetadata(t *testing.T) {
	release, err := os.ReadFile("release.json")
	if err != nil {
		t.Fatal(err)
	}
	semantic, err := os.ReadFile("semantic_sources.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := validateReleaseMetadata(release, semantic); err != nil {
		t.Fatal(err)
	}
	var changed map[string]any
	if err := json.Unmarshal(semantic, &changed); err != nil {
		t.Fatal(err)
	}
	changed["source_lock_sha256"] = "another lock"
	mismatched, err := json.Marshal(changed)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateReleaseMetadata(release, mismatched); err == nil {
		t.Fatal("accepted semantic output from another source lock")
	}
}

// validateReleaseMetadata compares the manifests with the compiled catalog facts.
func validateReleaseMetadata(releaseJSON, semanticJSON []byte) error {
	var release struct {
		MinecraftVersion string          `json:"minecraft_version"`
		ProtocolVersion  int             `json:"protocol_version"`
		SourceLockSHA256 string          `json:"source_lock_sha256"`
		SourceLock       json.RawMessage `json:"source_lock"`
	}
	if err := json.Unmarshal(releaseJSON, &release); err != nil {
		return err
	}
	if release.MinecraftVersion != data.MinecraftVersion || release.ProtocolVersion != data.ProtocolVersion {
		return fmt.Errorf("release target differs from the compiled catalog")
	}
	var canonical bytes.Buffer
	if err := json.Compact(&canonical, release.SourceLock); err != nil {
		return err
	}
	digest := fmt.Sprintf("%x", sha256.Sum256(canonical.Bytes()))
	for _, identity := range []string{release.SourceLockSHA256, data.SourceLockSHA256, data.SemanticSourceLockSHA256} {
		if digest != identity {
			return fmt.Errorf("catalog source-lock identities differ; regenerate both projections")
		}
	}
	var lock struct {
		MinecraftVersion string `json:"minecraft_version"`
		ProtocolVersion  int    `json:"protocol_version"`
		Inputs           map[string]struct {
			Revision string `json:"revision"`
		} `json:"inputs"`
		Semantic struct {
			CloudburstRef string `json:"cloudburst_ref"`
			BDSVersion    string `json:"bds_version"`
		} `json:"semantic"`
	}
	if err := json.Unmarshal(release.SourceLock, &lock); err != nil {
		return err
	}
	if lock.MinecraftVersion != data.MinecraftVersion || lock.ProtocolVersion != data.ProtocolVersion ||
		lock.Semantic.CloudburstRef != data.CloudburstRef || lock.Semantic.BDSVersion != data.BDSVersion ||
		lock.Inputs["block_shapes"].Revision != data.BlockShapesRef {
		return fmt.Errorf("compiled source pins differ from the release lock")
	}
	var semantic struct {
		SchemaVersion    int         `json:"schema_version"`
		CloudburstRef    string      `json:"cloudburst_ref"`
		BDSVersion       string      `json:"bds_version"`
		BlockShapesRef   string      `json:"block_shapes_ref"`
		SourceLockSHA256 string      `json:"source_lock_sha256"`
		Counts           data.Counts `json:"counts"`
	}
	if err := json.Unmarshal(semanticJSON, &semantic); err != nil {
		return err
	}
	if semantic.SchemaVersion != data.SchemaVersion || semantic.SourceLockSHA256 != digest ||
		semantic.CloudburstRef != data.CloudburstRef || semantic.BDSVersion != data.BDSVersion ||
		semantic.BlockShapesRef != data.BlockShapesRef || semantic.Counts != data.GeneratedCounts {
		return fmt.Errorf("semantic manifest differs from the compiled catalog")
	}
	return nil
}
