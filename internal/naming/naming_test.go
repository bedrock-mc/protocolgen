package naming

import (
	"strings"
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

func TestPublicTypeNameCollapsesImmediateDuplicateTokens(t *testing.T) {
	if got := PublicTypeName("MemoryMemoryCategory"); got != "MemoryCategory" {
		t.Fatalf("PublicTypeName duplicate token = %q, want MemoryCategory", got)
	}
}

func TestValidateRequiredEntriesListsArtifactTypeIDs(t *testing.T) {
	m := manifest.Manifest{
		Packets: []manifest.Packet{{Fields: []manifest.Field{
			{Encode: manifest.Node{Kind: manifest.KindStruct, TypeID: "GameRuleRuleValueEmpty0"}},
			{Encode: manifest.Node{Kind: manifest.KindStruct, TypeID: "TypedClientNetId<struct ItemStackRequestIdTag, int32_t, 0>"}},
		}}},
	}
	err := ValidateRequiredEntries(m, Overlay{})
	if err == nil || !strings.Contains(err.Error(), "GameRuleRuleValueEmpty0") || !strings.Contains(err.Error(), "TypedClientNetId") {
		t.Fatalf("ValidateRequiredEntries error = %v, want all artifact TypeIDs", err)
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

func TestResolverDerivesReviewedUnionFamilyName(t *testing.T) {
	node := manifest.Union(manifest.Primitive("u8"),
		manifest.Variant{Value: 0, Name: "ItemStackRequestCereal::TakeActionData", Encode: manifest.Node{Kind: manifest.KindStruct, TypeID: "ItemStackRequestCereal::TakeActionData"}},
		manifest.Variant{Value: 1, Name: "ItemStackRequestCereal::PlaceActionData", Encode: manifest.Node{Kind: manifest.KindStruct, TypeID: "ItemStackRequestCereal::PlaceActionData"}},
	)
	r := NewResolver(Overlay{Names: map[string]string{
		"ItemStackRequestCereal::TakeActionData":  "TakeStackRequestAction",
		"ItemStackRequestCereal::PlaceActionData": "PlaceStackRequestAction",
	}})
	got, err := r.Resolve(node, "ItemStackRequestCereal", func(value string) string { return value })
	if err != nil {
		t.Fatal(err)
	}
	if got != "StackRequestAction" {
		t.Fatalf("resolved union family = %q, want StackRequestAction", got)
	}
}

func TestValidateOverlayRejectsStaleTypeID(t *testing.T) {
	m := manifest.Manifest{
		Target:  manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Packets: []manifest.Packet{{Fields: []manifest.Field{{Encode: manifest.Node{Kind: manifest.KindStruct, TypeID: "Current"}}}}},
	}
	document := Document{
		SchemaVersion: 1,
		Target:        m.Target,
		Entries:       []Entry{{TypeID: "Stale", Name: "Stale", Rationale: "test"}},
	}
	if err := ValidateOverlay(m, document); err == nil || !strings.Contains(err.Error(), "stale TypeID") {
		t.Fatalf("ValidateOverlay error = %v, want stale TypeID failure", err)
	}
}
