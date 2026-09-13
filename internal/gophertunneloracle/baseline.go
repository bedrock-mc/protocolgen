package gophertunneloracle

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"protocolgen/internal/manifest"
)

// divergenceFingerprint binds a reviewed conflict to both complete input shapes and the oracle revision.
func divergenceFingerprint(packet manifest.Packet, oracle sourcePacket, paths []PathResult, lock Lock) (string, error) {
	data, err := json.Marshal(struct {
		Packet     manifest.Packet
		Operations []sourceOperation
		Paths      []PathResult
		Oracle     OracleSource
	}{packet, oracle.Operations, paths, OracleSource{Repo: lock.Gophertunnel.Repo, Commit: lock.Gophertunnel.Commit}})
	if err != nil {
		return "", err
	}
	var value any
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&value); err != nil {
		return "", err
	}
	removeDiagnosticSites(value)
	data, err = json.Marshal(value)
	if err != nil {
		return "", err
	}
	digest := sha256.Sum256(data)
	return "sha256:" + hex.EncodeToString(digest[:]), nil
}

// removeDiagnosticSites excludes checkout-dependent locations while retaining every semantic input field.
func removeDiagnosticSites(value any) {
	switch value := value.(type) {
	case map[string]any:
		delete(value, "Site")
		delete(value, "gophertunnel_site")
		for _, child := range value {
			removeDiagnosticSites(child)
		}
	case []any:
		for _, child := range value {
			removeDiagnosticSites(child)
		}
	}
}
