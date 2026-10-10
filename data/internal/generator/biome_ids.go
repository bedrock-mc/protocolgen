package generator

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type biomeIDCatalog struct {
	SchemaVersion    int    `json:"schema_version"`
	MinecraftVersion string `json:"minecraft_version"`
	ProtocolVersion  int    `json:"protocol_version"`
	Source           string `json:"source"`
	SourceSHA256     string `json:"source_sha256"`
	Biomes           []struct {
		Name string `json:"name"`
		ID   int32  `json:"id"`
	} `json:"biomes"`
}

// DecodeBiomeIDs validates the versioned biome-ID input before it joins the catalog.
func DecodeBiomeIDs(data []byte, minecraftVersion string, protocolVersion int) (map[string]int32, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var catalog biomeIDCatalog
	if err := decoder.Decode(&catalog); err != nil {
		return nil, fmt.Errorf("decode biome IDs: %w", err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return nil, fmt.Errorf("biome IDs have trailing JSON")
	}
	sourceDigest, digestErr := hex.DecodeString(catalog.SourceSHA256)
	if catalog.SchemaVersion != 1 || catalog.MinecraftVersion != minecraftVersion || catalog.ProtocolVersion != protocolVersion || catalog.Source == "" || digestErr != nil || len(sourceDigest) != 32 {
		return nil, fmt.Errorf("biome IDs do not match Minecraft %s / protocol %d or have incomplete provenance", minecraftVersion, protocolVersion)
	}
	if len(catalog.Biomes) == 0 {
		return nil, fmt.Errorf("biome ID input is empty")
	}
	ids := make(map[string]int32, len(catalog.Biomes))
	seenIDs := make(map[int32]string, len(catalog.Biomes))
	for _, biome := range catalog.Biomes {
		if !strings.HasPrefix(biome.Name, "minecraft:") || biome.ID < 0 || biome.ID > 65535 {
			return nil, fmt.Errorf("invalid biome ID record %q/%d", biome.Name, biome.ID)
		}
		if _, exists := ids[biome.Name]; exists {
			return nil, fmt.Errorf("duplicate biome name %q", biome.Name)
		}
		if previous, exists := seenIDs[biome.ID]; exists {
			return nil, fmt.Errorf("duplicate biome ID %d for %q and %q", biome.ID, previous, biome.Name)
		}
		ids[biome.Name], seenIDs[biome.ID] = biome.ID, biome.Name
	}
	return ids, nil
}
