package naming

import (
	"testing"

	"protocolgen/internal/manifest"
)

func TestPublicTypeNameNormalizesQualifiedAndFilenamePacketNames(t *testing.T) {
	want := "CommandData"
	for _, input := range []string{
		"AvailableCommandsPacket::CommandData",
		"AvailableCommandsPacketCommandData.json#",
	} {
		if got := PublicTypeName(input); got != want {
			t.Errorf("PublicTypeName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestResolverUsesOneIdentityAndRejectsCollisions(t *testing.T) {
	r := NewResolver(Overlay{})
	first := manifest.Node{Kind: manifest.KindStruct, TypeID: "Thing"}
	second := manifest.Node{Kind: manifest.KindStruct, TypeID: "Thing.json#"}
	casing := func(value string) string { return value }
	name, err := r.Resolve(first, "ThingStruct", casing)
	if err != nil {
		t.Fatal(err)
	}
	again, err := r.Resolve(first, "DifferentHint", casing)
	if err != nil {
		t.Fatal(err)
	}
	if again != name {
		t.Fatalf("same identity resolved to %q after %q", again, name)
	}
	if _, err := r.Resolve(second, "ThingStruct", casing); err == nil {
		t.Fatal("Resolve accepted a public-name collision")
	}
}
