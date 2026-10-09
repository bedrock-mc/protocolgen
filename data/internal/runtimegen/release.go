package runtimegen

import (
	"encoding/json"
	"fmt"
	"go/format"

	"github.com/bedrock-mc/protocolgen/data/source"
)

// Release records the target and all independent source pins beside the catalog.
// Its identity is the SHA-256 of the compact JSON encoding of the complete lock.
func Release(lock source.Lock) (map[string][]byte, error) {
	digest, err := lock.SHA256()
	if err != nil {
		return nil, err
	}
	manifest, err := json.MarshalIndent(struct {
		MinecraftVersion string         `json:"minecraft_version"`
		ProtocolVersion  int            `json:"protocol_version"`
		SourceLockSHA256 string         `json:"source_lock_sha256"`
		SourceLock       source.Lock    `json:"source_lock"`
		Target           source.Release `json:"target"`
	}{lock.Target.MinecraftVersion, lock.Target.ProtocolVersion, digest, lock, lock.Target}, "", "  ")
	if err != nil {
		return nil, err
	}
	goSource, err := format.Source(fmt.Appendf(nil, `%s

package data

// MinecraftVersion is the target release of the active catalog.
const MinecraftVersion = %q

// ProtocolVersion is the protocol used by the target release.
const ProtocolVersion = %d

// ProtocolSnapshot identifies the compatible generated codec directory.
const ProtocolSnapshot = %q

// ReleaseChannel distinguishes retail labels from preview labels.
const ReleaseChannel = %q

// SourceLockSHA256 identifies every pinned input in the source lock.
// It includes the resolved release record as well as every input identity.
const SourceLockSHA256 = %q
`, GeneratedMarker, lock.Target.MinecraftVersion, lock.Target.ProtocolVersion, lock.Target.Snapshot, lock.Target.Channel, digest))
	if err != nil {
		return nil, err
	}
	return map[string][]byte{
		"release_generated.go": goSource,
		"release.json":         append(manifest, '\n'),
	}, nil
}
