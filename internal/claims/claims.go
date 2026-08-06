// Package claims is the source-facing side of reconciliation. A claim is a
// complete machine-derived assertion about one packet field, including its
// semantic identity and locator. It is never a confidence vote.
package claims

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"protocolgen/internal/manifest"
)

type Claim struct {
	SourceID      string             `json:"source_id"`
	Locator       string             `json:"locator"`
	PacketID      uint32             `json:"packet_id"`
	PacketName    string             `json:"packet_name"`
	Direction     manifest.Direction `json:"direction"`
	FieldPath     string             `json:"field_path"`
	Ordinal       int                `json:"ordinal"`
	Name          string             `json:"name"`
	Semantic      string             `json:"semantic,omitempty"`
	TypeID        string             `json:"type_id,omitempty"`
	Encode        manifest.Node      `json:"encode"`
	Decode        *manifest.Node     `json:"decode,omitempty"`
	Symmetry      manifest.Symmetry  `json:"symmetry"`
	Reserved      bool               `json:"reserved,omitempty"`
	Ignored       bool               `json:"ignored,omitempty"`
	Compatibility []string           `json:"compatibility,omitempty"`
}

type Result struct {
	Pin       manifest.SourcePin
	Target    manifest.Target
	Claims    []Claim
	Overrides []manifest.OverrideProof
}

func Fingerprint(claim Claim) (string, error) {
	return digest(claim)
}

func ContextFingerprint(target manifest.Target, input []Claim) (string, error) {
	claims := append([]Claim(nil), input...)
	sort.SliceStable(claims, func(i, j int) bool {
		if claims[i].SourceID != claims[j].SourceID {
			return claims[i].SourceID < claims[j].SourceID
		}
		if claims[i].FieldPath != claims[j].FieldPath {
			return claims[i].FieldPath < claims[j].FieldPath
		}
		return claims[i].Ordinal < claims[j].Ordinal
	})
	return digest(struct {
		Target manifest.Target `json:"target"`
		Claims []Claim         `json:"claims"`
	}{Target: target, Claims: claims})
}

func digest(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("fingerprint JSON: %w", err)
	}
	sum := sha256.Sum256(data)
	return "sha256:" + hex.EncodeToString(sum[:]), nil
}
