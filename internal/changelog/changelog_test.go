package changelog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateReportsWireConstraintAndEnumChanges(t *testing.T) {
	from := t.TempDir()
	to := t.TempDir()

	writeSchema(t, from, "RequestNetworkSettingsPacket", `{
  "title":"RequestNetworkSettingsPacket","x-minecraft-version":"1.26.40-beta.0","x-protocol-version":2168,
  "$ref":"./RequestNetworkSettingsPacketPayload.json","$metaProperties":{"[cereal:packet]":193}}
`)
	writeSchema(t, to, "RequestNetworkSettingsPacket", `{
  "title":"RequestNetworkSettingsPacket","x-minecraft-version":"1.26.50","x-protocol-version":2169,
  "$ref":"./RequestNetworkSettingsPacketPayload.json","$metaProperties":{"[cereal:packet]":193}}
`)
	writeSchema(t, from, "RequestNetworkSettingsPacketPayload", objectSchema("1.26.40-beta.0", 2168, `
    "ClientNetworkVersion":{"type":"integer","x-underlying-type":"int32","x-serialization-options":["Big Endian"],"x-ordinal-index":0,"minimum":2168,"maximum":2168}`, `"ClientNetworkVersion"`))
	writeSchema(t, to, "RequestNetworkSettingsPacketPayload", objectSchema("1.26.50", 2169, `
    "ClientNetworkVersion":{"type":"integer","x-underlying-type":"int32","x-serialization-options":["Big Endian"],"x-ordinal-index":0,"minimum":2169,"maximum":2169}`, `"ClientNetworkVersion"`))

	writeSchema(t, from, "TextDataPayload", objectSchema("1.26.40-beta.0", 2168, `
    "Text":{"type":"string","x-ordinal-index":0},
    "DepthTest":{"type":"boolean","x-underlying-type":"boolean","x-ordinal-index":1}`, `"Text","DepthTest"`))
	writeSchema(t, to, "TextDataPayload", objectSchema("1.26.50", 2169, `
    "Text":{"type":"string","x-ordinal-index":0},
    "LineGapHeight":{"type":"number","x-underlying-type":"float","x-ordinal-index":1},
    "DepthTest":{"type":"boolean","x-underlying-type":"boolean","x-ordinal-index":2}`, `"Text","DepthTest"`))
	writeSchema(t, from, "PrimitiveShapesPacket", packetSchema("1.26.40-beta.0", 2168, 328, "PrimitiveShapesPacketPayload"))
	writeSchema(t, to, "PrimitiveShapesPacket", packetSchema("1.26.50", 2169, 328, "PrimitiveShapesPacketPayload"))
	writeSchema(t, from, "PrimitiveShapesPacketPayload", objectWithRef("1.26.40-beta.0", 2168, "Shape", "TextDataPayload"))
	writeSchema(t, to, "PrimitiveShapesPacketPayload", objectWithRef("1.26.50", 2169, "Shape", "TextDataPayload"))

	writeSchema(t, from, "persona__AnimatedTextureType", enumSchema("1.26.40-beta.0", 2168, `"Face","Body32x32"`))
	writeSchema(t, to, "persona__AnimatedTextureType", enumSchema("1.26.50", 2169, `"None","Face","Body32x32"`))
	writeSchema(t, from, "PlayerSkinPacket", packetSchema("1.26.40-beta.0", 2168, 93, "PlayerSkinPacketPayload"))
	writeSchema(t, to, "PlayerSkinPacket", packetSchema("1.26.50", 2169, 93, "PlayerSkinPacketPayload"))
	writeSchema(t, from, "PlayerSkinPacketPayload", objectWithRef("1.26.40-beta.0", 2168, "AnimationType", "persona__AnimatedTextureType"))
	writeSchema(t, to, "PlayerSkinPacketPayload", objectWithRef("1.26.50", 2169, "AnimationType", "persona__AnimatedTextureType"))

	got, err := Generate(Config{
		FromDir: from, ToDir: to,
		FromBranch: "automated/1.26.40", ToBranch: "r/26_u4",
		FromUpstream: "0e00fe80f4", ToUpstream: "0f6a0bff19",
		FromFixer: "9fa8eb75f9", ToFixer: "9fa8eb75f9",
	})
	if err != nil {
		t.Fatal(err)
	}
	text := string(got)
	for _, want := range []string{
		"# Bedrock protocol changes — 2168 to 2169",
		"1.26.40-beta.0 → 1.26.50",
		"**0** new, **0** removed, **0** renamed, **3** modified",
		"**1** type(s) change the bytes on the wire",
		"**3** packet(s) touched, counting those that only embed a changed type",
		"constraints changed from maximum=2168, minimum=2168 to maximum=2169, minimum=2169",
		"### TextDataPayload\nstruct · **wire break**",
		"added `LineGapHeight` at ordinal 1 (float, optional)",
		"`DepthTest` moved from ordinal 1 to 2",
		"**Also affects:** `PrimitiveShapesPacket`",
		"### persona__AnimatedTextureType\nenum",
		"added value `None` = 0",
		"**Also affects:** `PlayerSkinPacket`",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\n%s", want, text)
		}
	}
}

func TestGenerateRejectsInconsistentMetadata(t *testing.T) {
	from := t.TempDir()
	to := t.TempDir()
	writeSchema(t, from, "One", objectSchema("1.0", 1, `"A":{"type":"string","x-ordinal-index":0}`, `"A"`))
	writeSchema(t, from, "Two", objectSchema("1.0", 2, `"A":{"type":"string","x-ordinal-index":0}`, `"A"`))
	writeSchema(t, to, "One", objectSchema("2.0", 3, `"A":{"type":"string","x-ordinal-index":0}`, `"A"`))
	if _, err := Generate(Config{FromDir: from, ToDir: to}); err == nil || !strings.Contains(err.Error(), "inconsistent protocol version") {
		t.Fatalf("expected inconsistent protocol version error, got %v", err)
	}
}

func TestGenerateIgnoresDocumentationOnlyChanges(t *testing.T) {
	from := t.TempDir()
	to := t.TempDir()
	writeSchema(t, from, "Shared", `{"title":"Shared","description":"old docs","x-minecraft-version":"1.0","x-protocol-version":1,"type":"object","properties":{"Value":{"description":"old field docs","type":"string","x-ordinal-index":0}},"required":["Value"]}`)
	writeSchema(t, to, "Shared", `{"title":"Shared","description":"new docs","x-minecraft-version":"2.0","x-protocol-version":2,"type":"object","properties":{"Value":{"description":"new field docs","type":"string","x-ordinal-index":0}},"required":["Value"]}`)

	got, err := Generate(Config{FromDir: from, ToDir: to})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(got), "## Modified Types") || !strings.Contains(string(got), "**0** new, **0** removed, **0** renamed, **0** modified") {
		t.Fatalf("documentation-only edits were reported as protocol changes:\n%s", got)
	}
}

func TestGenerateRendersAddedRemovedAndRenamedDefinitions(t *testing.T) {
	from := t.TempDir()
	to := t.TempDir()
	writeSchema(t, from, "OldType", objectSchema("1.0", 1, `"Value":{"type":"string","x-ordinal-index":0}`, `"Value"`))
	writeSchema(t, to, "NewType", objectSchema("2.0", 2, `"Value":{"type":"string","x-ordinal-index":0}`, `"Value"`))
	writeSchema(t, from, "OldPacket", packetSchema("1.0", 1, 7, "OldPacketPayload"))
	writeSchema(t, from, "OldPacketPayload", objectSchema("1.0", 1, `"Value":{"type":"string","x-ordinal-index":0}`, `"Value"`))
	writeSchema(t, to, "NewPacket", packetSchema("2.0", 2, 7, "NewPacketPayload"))
	writeSchema(t, to, "NewPacketPayload", objectSchema("2.0", 2, `"Value":{"type":"string","x-ordinal-index":0}`, `"Value"`))

	got, err := Generate(Config{FromDir: from, ToDir: to})
	if err != nil {
		t.Fatal(err)
	}
	text := string(got)
	for _, want := range []string{
		"**1** new, **1** removed, **1** renamed, **0** modified",
		"## Added Types\n\n### NewType",
		"## Removed Types\n\n### OldType",
		"## Renamed Packets\n\n- `OldPacket` → `NewPacket` (packet id `7`)",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\n%s", want, text)
		}
	}
}

func writeSchema(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name+".json"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func objectSchema(version string, protocol int, properties, required string) string {
	return `{"title":"Payload","x-minecraft-version":"` + version + `","x-protocol-version":` + itoa(protocol) + `,"type":"object","properties":{` + properties + `},"required":[` + required + `]}`
}

func packetSchema(version string, protocol, id int, payload string) string {
	return `{"title":"Packet","x-minecraft-version":"` + version + `","x-protocol-version":` + itoa(protocol) + `,"$ref":"./` + payload + `.json","$metaProperties":{"[cereal:packet]":` + itoa(id) + `}}`
}

func objectWithRef(version string, protocol int, field, ref string) string {
	return objectSchema(version, protocol, `"`+field+`":{"$ref":"./`+ref+`.json","x-ordinal-index":0}`, `"`+field+`"`)
}

func enumSchema(version string, protocol int, values string) string {
	return `{"title":"persona::AnimatedTextureType","x-minecraft-version":"` + version + `","x-protocol-version":` + itoa(protocol) + `,"type":"string","enum":[` + values + `],"x-underlying-type":"uint32"}`
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var b [20]byte
	i := len(b)
	for v > 0 {
		i--
		b[i] = byte('0' + v%10)
		v /= 10
	}
	return string(b[i:])
}
