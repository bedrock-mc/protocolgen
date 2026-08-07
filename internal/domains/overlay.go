// Package domains contains the reviewed shared-type-to-file mapping used by
// the source emitters.
package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"protocolgen/internal/manifest"
	"protocolgen/internal/naming"
)

type Entry struct {
	TypeID    string `json:"type_id"`
	Domain    string `json:"domain"`
	Rationale string `json:"rationale"`
}

type Document struct {
	SchemaVersion uint32          `json:"schema_version"`
	Target        manifest.Target `json:"target"`
	Entries       []Entry         `json:"entries"`
}

type Overlay struct {
	Domains map[string]string
}

func LoadOverlay(path string, m manifest.Manifest) (Overlay, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Overlay{}, fmt.Errorf("read domains overlay: %w", err)
	}
	var document Document
	if err := json.Unmarshal(data, &document); err != nil {
		return Overlay{}, fmt.Errorf("parse domains overlay: %w", err)
	}
	if err := ValidateOverlay(m, document); err != nil {
		return Overlay{}, err
	}
	overlay := Overlay{Domains: make(map[string]string, len(document.Entries))}
	for _, entry := range document.Entries {
		overlay.Domains[entry.TypeID] = entry.Domain
	}
	return overlay, nil
}

func ValidateOverlay(m manifest.Manifest, document Document) error {
	if document.SchemaVersion != 1 {
		return fmt.Errorf("domains overlay schema_version %d is not v1", document.SchemaVersion)
	}
	if document.Target.MinecraftVersion != m.Target.MinecraftVersion || document.Target.ProtocolVersion != m.Target.ProtocolVersion {
		return fmt.Errorf("domains overlay target does not match manifest target")
	}
	existing := naming.TypeIDs(m)
	seen := make(map[string]bool, len(document.Entries))
	validDomain := regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	for index, entry := range document.Entries {
		if entry.TypeID == "" || entry.Domain == "" || entry.Rationale == "" {
			return fmt.Errorf("domains overlay entry[%d] requires type_id, domain, and rationale", index)
		}
		if !existing[entry.TypeID] {
			return fmt.Errorf("domains overlay entry[%d] refers to stale TypeID %q", index, entry.TypeID)
		}
		if seen[entry.TypeID] {
			return fmt.Errorf("domains overlay contains duplicate TypeID %q", entry.TypeID)
		}
		seen[entry.TypeID] = true
		if !validDomain.MatchString(entry.Domain) {
			return fmt.Errorf("domains overlay entry %q has invalid domain %q", entry.TypeID, entry.Domain)
		}
	}
	overlay := Overlay{Domains: make(map[string]string, len(seen))}
	for typeID := range seen {
		overlay.Domains[typeID] = "assigned"
	}
	if err := ValidateAssignments(m, overlay); err != nil {
		return err
	}
	return nil
}

func ValidateAssignments(m manifest.Manifest, overlay Overlay) error {
	existing := naming.TypeIDs(m)
	var missing []string
	for typeID := range existing {
		if overlay.Domains[typeID] == "" {
			missing = append(missing, typeID)
		}
	}
	if len(missing) != 0 {
		sort.Strings(missing)
		return fmt.Errorf("domains overlay missing assignments for shared TypeIDs: %s", strings.Join(missing, ", "))
	}
	return nil
}

func (o Overlay) Domain(typeID string) string {
	if o.Domains == nil {
		return ""
	}
	return o.Domains[typeID]
}
