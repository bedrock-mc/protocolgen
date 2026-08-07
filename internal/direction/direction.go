// Package direction loads the reviewed packet-direction overlay for one target.
package direction

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"protocolgen/internal/manifest"
)

const SchemaVersion uint32 = 1

type Direction string

const (
	DirectionClient Direction = "client"
	DirectionServer Direction = "server"
	DirectionBoth   Direction = "both"
)

type Table struct {
	SchemaVersion uint32          `json:"schema_version"`
	Target        manifest.Target `json:"target"`
	Source        Source          `json:"source"`
	Packets       []Entry         `json:"packets"`
}

type Source struct {
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Locator    string `json:"locator"`
	SHA256     string `json:"sha256"`
}

type Entry struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Direction Direction `json:"direction"`
	Evidence  Evidence  `json:"evidence"`
}

type Evidence struct {
	Locator string `json:"locator"`
	Summary string `json:"summary"`
}

func Load(path string) (Table, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Table{}, fmt.Errorf("read direction table: %w", err)
	}
	var table Table
	if err := json.Unmarshal(data, &table); err != nil {
		return Table{}, fmt.Errorf("parse direction table: %w", err)
	}
	if err := table.validateShape(); err != nil {
		return Table{}, fmt.Errorf("validate direction table: %w", err)
	}
	return table, nil
}

func (t Table) Validate(target manifest.Target, packets []manifest.Packet) error {
	if err := t.validateShape(); err != nil {
		return err
	}
	if t.Target != target {
		return fmt.Errorf("direction table target does not match manifest target")
	}
	byID := make(map[uint32]manifest.Packet, len(packets))
	for _, packet := range packets {
		byID[packet.ID] = packet
	}
	for _, entry := range t.Packets {
		packet, ok := byID[entry.ID]
		if !ok {
			return fmt.Errorf("direction table names unknown packet id %d", entry.ID)
		}
		if packet.Name != entry.Name {
			return fmt.Errorf("direction table packet id %d names %q, manifest names %q", entry.ID, entry.Name, packet.Name)
		}
		delete(byID, entry.ID)
	}
	for id, packet := range byID {
		return fmt.Errorf("direction table has missing direction for packet id %d (%s)", id, packet.Name)
	}
	return nil
}

func (t Table) Apply(value *manifest.Manifest) error {
	if value == nil {
		return fmt.Errorf("cannot apply direction table to nil manifest")
	}
	if err := t.Validate(value.Target, value.Packets); err != nil {
		return err
	}
	byID := make(map[uint32]Direction, len(t.Packets))
	for _, entry := range t.Packets {
		byID[entry.ID] = entry.Direction
	}
	for index := range value.Packets {
		value.Packets[index].Direction = manifestDirection(byID[value.Packets[index].ID])
	}
	return nil
}

func (t Table) validateShape() error {
	if t.SchemaVersion != SchemaVersion {
		return fmt.Errorf("schema_version %d is not v%d", t.SchemaVersion, SchemaVersion)
	}
	if t.Target.MinecraftVersion == "" || t.Target.ProtocolVersion <= 0 {
		return fmt.Errorf("target must pin minecraft_version and positive protocol_version")
	}
	if t.Source.Repository == "" || t.Source.Revision == "" || t.Source.Locator == "" || t.Source.SHA256 == "" {
		return fmt.Errorf("source is missing repository, revision, locator, or sha256")
	}
	if len(t.Source.Revision) != 40 || !isHex(t.Source.Revision) {
		return fmt.Errorf("source revision is not a full SHA-1")
	}
	if len(t.Source.SHA256) != len("sha256:")+64 || !strings.HasPrefix(t.Source.SHA256, "sha256:") || !isHex(strings.TrimPrefix(t.Source.SHA256, "sha256:")) {
		return fmt.Errorf("source sha256 is not a SHA-256 fingerprint")
	}
	seen := make(map[uint32]bool, len(t.Packets))
	for index, entry := range t.Packets {
		if entry.ID == 0 || entry.Name == "" {
			return fmt.Errorf("packets[%d] is missing id or name", index)
		}
		if seen[entry.ID] {
			return fmt.Errorf("duplicate direction entry for packet id %d", entry.ID)
		}
		seen[entry.ID] = true
		if !validDirection(entry.Direction) {
			return fmt.Errorf("packet id %d has invalid direction %q", entry.ID, entry.Direction)
		}
		if entry.Evidence.Locator == "" || entry.Evidence.Summary == "" {
			return fmt.Errorf("packet id %d has incomplete evidence", entry.ID)
		}
	}
	return nil
}

func validDirection(value Direction) bool {
	switch value {
	case DirectionClient, DirectionServer, DirectionBoth:
		return true
	default:
		return false
	}
}

func manifestDirection(value Direction) manifest.Direction {
	switch value {
	case DirectionClient:
		return manifest.DirectionServerbound
	case DirectionServer:
		return manifest.DirectionClientbound
	case DirectionBoth:
		return manifest.DirectionBidirectional
	default:
		return manifest.DirectionUnknown
	}
}

func isHex(value string) bool {
	_, err := hex.DecodeString(value)
	return err == nil
}
