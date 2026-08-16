package vanilladata

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Target identifies one generated protocol snapshot.
type Target struct {
	MinecraftVersion string `json:"minecraft_version"`
	ProtocolVersion  int    `json:"protocol_version"`
	Build            string `json:"build"`
}

// Manifest is the capture-relevant subset of a protocolgen manifest.
type Manifest struct {
	Target  Target      `json:"target"`
	Sources []SourcePin `json:"sources"`
	Packets []Packet    `json:"packets"`
}

// SourcePin identifies an upstream manifest input.
type SourcePin struct {
	Kind     string `json:"kind"`
	Revision string `json:"revision"`
}

// Packet is the capture-relevant subset of a manifest packet entry.
type Packet struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
}

// SourceConfig pins the BDS and gophertunnel inputs for one capture.
type SourceConfig struct {
	SchemaVersion    int                `json:"schema_version"`
	MinecraftVersion string             `json:"minecraft_version"`
	ProtocolVersion  int                `json:"protocol_version"`
	BDS              BDSSource          `json:"bds"`
	Gophertunnel     GophertunnelSource `json:"gophertunnel"`
	ServerProperties map[string]string  `json:"server_properties"`
}

// BDSSource pins the downloadable Linux BDS build.
type BDSSource struct {
	Version string      `json:"version"`
	Linux   LinuxSource `json:"linux"`
}

// LinuxSource identifies and authenticates a BDS archive.
type LinuxSource struct {
	URL           string `json:"url"`
	ArchiveSHA256 string `json:"archive_sha256"`
}

// GophertunnelSource pins the codec used by the bot.
type GophertunnelSource struct {
	Repository    string `json:"repository"`
	ModulePath    string `json:"module_path"`
	Revision      string `json:"revision"`
	ModuleVersion string `json:"module_version"`
}

// LoadManifest loads the capture-relevant portion of a generated manifest.
func LoadManifest(path string) (Manifest, error) {
	var value Manifest
	if err := loadJSON(path, &value); err != nil {
		return Manifest{}, fmt.Errorf("load manifest: %w", err)
	}
	return value, nil
}

// LoadSourceConfig loads and validates a version-locked vanilla capture source.
func LoadSourceConfig(path string) (SourceConfig, error) {
	var value SourceConfig
	if err := loadJSON(path, &value); err != nil {
		return SourceConfig{}, fmt.Errorf("load vanilla source: %w", err)
	}
	if value.SchemaVersion != 1 {
		return SourceConfig{}, fmt.Errorf("unsupported vanilla source schema %d", value.SchemaVersion)
	}
	if value.BDS.Version == "" || value.BDS.Linux.URL == "" || !validHex(value.BDS.Linux.ArchiveSHA256, 64) {
		return SourceConfig{}, fmt.Errorf("vanilla source has incomplete BDS provenance")
	}
	if !strings.Contains(value.BDS.Linux.URL, "bedrock-server-"+value.BDS.Version+".zip") {
		return SourceConfig{}, fmt.Errorf("BDS version %s does not match archive URL", value.BDS.Version)
	}
	if value.Gophertunnel.Repository == "" || value.Gophertunnel.ModulePath == "" || value.Gophertunnel.ModuleVersion == "" || !validHex(value.Gophertunnel.Revision, 40) {
		return SourceConfig{}, fmt.Errorf("vanilla source has incomplete gophertunnel provenance")
	}
	if value.ServerProperties["online-mode"] != "false" {
		return SourceConfig{}, fmt.Errorf("vanilla source must configure BDS with online-mode=false")
	}
	if value.ServerProperties["server-port"] == "" || value.ServerProperties["level-seed"] == "" || value.ServerProperties["level-type"] == "" {
		return SourceConfig{}, fmt.Errorf("vanilla source must pin server-port, level-seed, and level-type")
	}
	return value, nil
}

func loadJSON(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}

func validHex(value string, length int) bool {
	if len(value) != length {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}
