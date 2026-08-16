//go:build protocolgen_12640

package vanilladata

import "testing"

func TestProtocolgen12640SelectsThe12640GeneratedSnapshot(t *testing.T) {
	if err := ValidateGeneratedTarget(Target{MinecraftVersion: "1.26.40", ProtocolVersion: 2168}); err != nil {
		t.Fatalf("ValidateGeneratedTarget for protocolgen_12640: %v", err)
	}
}
