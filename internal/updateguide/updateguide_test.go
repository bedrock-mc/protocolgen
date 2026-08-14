package updateguide

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateRendersChangedPacketsTypesAndEnums(t *testing.T) {
	directory := t.TempDir()
	writeFixture(t, directory, "__protocoldoc.json", `[]`)
	writeFixture(t, directory, "RequestNetworkSettingsPacket.json", `{
		"title":"RequestNetworkSettingsPacket",
		"$ref":"./RequestNetworkSettingsPacketPayload.json",
		"$metaProperties":{"[cereal:packet]":193}
	}`)
	writeFixture(t, directory, "RequestNetworkSettingsPacketPayload.json", `{
		"title":"RequestNetworkSettingsPacketPayload",
		"type":"object",
		"properties":{"ClientNetworkVersion":{"type":"integer","x-underlying-type":"int32","x-serialization-options":["Big Endian"],"x-ordinal-index":0,"minimum":2169,"maximum":2169}},
		"required":["ClientNetworkVersion"]
	}`)
	writeFixture(t, directory, "TextDataPayload.json", `{
		"title":"TextDataPayload",
		"type":"object",
		"properties":{
			"Text":{"description":"Text (string) of the debug text shape.","type":"string","x-ordinal-index":0},
			"UseRotation":{"description":"Whether to use specified rotation, otherwise faces camera.","type":"boolean","x-underlying-type":"boolean","x-ordinal-index":1},
			"BackgroundColor":{"description":"Color to draw background.","$ref":"./Color.json","x-ordinal-index":2},
			"LineGapHeight":{"description":"Line gap height for multiline text rendering.","type":"number","x-underlying-type":"float","x-ordinal-index":3},
			"DepthTest":{"description":"Whether other objects block seeing the text.","type":"boolean","x-underlying-type":"boolean","x-ordinal-index":4},
			"ShowBackface":{"description":"Whether the background backface should be rendered.","type":"boolean","x-underlying-type":"boolean","x-ordinal-index":5},
			"ShowTextBackface":{"description":"Whether the text backface should be rendered.","type":"boolean","x-underlying-type":"boolean","x-ordinal-index":6}
		},
		"required":["Text","UseRotation","DepthTest","ShowBackface","ShowTextBackface"]
	}`)
	writeFixture(t, directory, "Color.json", `{"title":"Color","type":"object","properties":{},"required":[]}`)
	writeFixture(t, directory, "persona__AnimatedTextureType.json", `{
		"title":"persona::AnimatedTextureType",
		"type":"string",
		"enum":["None","Face","Body32x32","Body128x128"],
		"x-underlying-type":"uint32"
	}`)

	changelog := `# Bedrock protocol changes — 2168 to 2169

1.26.40-beta.0 → 1.26.50

| | From | To |
| --- | --- | --- |
| Branch | automated/1.26.40 | r/26_u4 |
| Protocol | 2168 | 2169 |
| Game version | 1.26.40-beta.0 | 1.26.50 |
| Upstream commit | 0e00fe80f4 | 0f6a0bff19 |
| Fixer commit | 9fa8eb75f9 | 9fa8eb75f9 |

## Modified Packets

### RequestNetworkSettingsPacket
packet id 193 · struct

- ClientNetworkVersion constraints changed from maximum=2168, minimum=2168 to maximum=2169, minimum=2169

| # | Before | After |
| --- | --- | --- |
| 0 | int32 | int32 |

## Modified Types

### TextDataPayload
struct · **wire break**

- added LineGapHeight at ordinal 3 (float, optional)
- DepthTest moved from ordinal 3 to 4

**Also affects:** PrimitiveShapesPacket

## Modified Enums

### persona__AnimatedTextureType
enum

- added value None = 0

**Also affects:** PlayerListPacket, PlayerSkinPacket
`

	output, err := Generate([]byte(changelog), directory)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	text := string(output)
	for _, want := range []string{
		"# gophertunnel update guide — protocol 2168 to 2169",
		"Read `changelog.md` for the human-readable diff.",
		"type RequestNetworkSettingsPacket struct {",
		"ClientNetworkVersion int32",
		"return IDRequestNetworkSettingsPacket",
		"io.BEInt32(&pk.ClientNetworkVersion)",
		"type TextData struct {",
		"LineGapHeight protocol.Optional[float32]",
		"protocol.OptionalMarshaler(io, &pk.BackgroundColor)",
		"protocol.OptionalFunc(io, &pk.LineGapHeight, io.Float32)",
		"**Also update:** `PrimitiveShapesPacket` — it embeds this type.",
		"PersonaAnimatedTextureTypeNone",
		"PersonaAnimatedTextureTypeBody128x128 = 3",
		"**Also update:** `PlayerListPacket`, `PlayerSkinPacket` — each embeds this type.",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output does not contain %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, "// ClientNetworkVersion ...") {
		t.Fatalf("output contains a placeholder field comment:\n%s", text)
	}
	if strings.Contains(text, "| # | Before | After |") {
		t.Fatalf("output copied a per-definition changelog table into its metadata header:\n%s", text)
	}
}

func TestGenerateRejectsUnknownChangedSchema(t *testing.T) {
	_, err := Generate([]byte("# Bedrock protocol changes — 1 to 2\n\n## Modified Types\n\n### MissingType\n"), t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "MissingType") {
		t.Fatalf("Generate error = %v, want missing-schema error", err)
	}
}

func TestGenerateRendersPrimitiveSlicesWithoutPlaceholders(t *testing.T) {
	directory := t.TempDir()
	writeFixture(t, directory, "Samples.json", `{
		"title":"Samples",
		"type":"object",
		"properties":{"Values":{"type":"array","items":{"type":"number","x-underlying-type":"float"},"x-ordinal-index":0}},
		"required":["Values"]
	}`)
	changelog := "# Bedrock protocol changes — 1 to 2\n\n## Modified Types\n\n### Samples\nstruct\n"
	output, err := Generate([]byte(changelog), directory)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	text := string(output)
	if !strings.Contains(text, "protocol.FuncSlice(io, &pk.Values, io.Float32)") {
		t.Fatalf("output does not marshal the primitive slice:\n%s", text)
	}
	if strings.Contains(text, "marshal value") {
		t.Fatalf("output contains a marshal placeholder:\n%s", text)
	}
}

func TestGenerateRendersOptionalPrimitiveSlices(t *testing.T) {
	directory := t.TempDir()
	writeFixture(t, directory, "Samples.json", `{
		"title":"Samples",
		"type":"object",
		"properties":{"Values":{"type":"array","items":{"type":"number","x-underlying-type":"float"},"x-ordinal-index":0}},
		"required":[]
	}`)
	changelog := "# Bedrock protocol changes — 1 to 2\n\n## Modified Types\n\n### Samples\nstruct\n"
	output, err := Generate([]byte(changelog), directory)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	text := string(output)
	for _, want := range []string{
		"Values protocol.Optional[[]float32]",
		"protocol.OptionalFunc(io, &pk.Values, func(value *[]float32)",
		"protocol.FuncSlice(io, value, io.Float32)",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("output does not contain %q:\n%s", want, text)
		}
	}
}

func writeFixture(t *testing.T, directory, name, contents string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(directory, name), []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
}
