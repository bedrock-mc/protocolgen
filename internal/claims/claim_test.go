package claims

import (
	"testing"

	"protocolgen/internal/manifest"
)

func TestFingerprintIncludesCompleteClaimContext(t *testing.T) {
	claim := Claim{
		SourceID: "endstone",
		Locator:  "packets/Vocabulary.field",
		PacketID: 1, PacketName: "Vocabulary", FieldPath: "Vocabulary.Value", Ordinal: 0, Name: "Value",
		Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric,
	}
	first, err := Fingerprint(claim)
	if err != nil {
		t.Fatalf("Fingerprint: %v", err)
	}
	claim.Name = "ChangedName"
	second, err := Fingerprint(claim)
	if err != nil {
		t.Fatalf("Fingerprint changed claim: %v", err)
	}
	if first == second {
		t.Fatalf("claim fingerprint ignored semantic context")
	}
}
