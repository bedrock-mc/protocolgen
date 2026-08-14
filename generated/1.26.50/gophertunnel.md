# gophertunnel update guide — protocol 2168 to 2169

1.26.40-beta.0 → 1.26.50

| | From | To |
| --- | --- | --- |
| Branch | `automated/1.26.40` | `r/26_u4` |
| Protocol | 2168 | 2169 |
| Game version | 1.26.40-beta.0 | 1.26.50 |
| Upstream commit | `0e00fe80f4` | `0f6a0bff19` |
| Fixer commit | `9fa8eb75f9` | `9fa8eb75f9` |

Read `changelog.md` for the human-readable diff. This guide gives a target-version `Marshal` transcription for each changed definition.

**The Go below is a transcription aid, not a patch.** Names come from the schema, so they may not match gophertunnel's existing names. Field comments are emitted only when Mojang provides a description.

## Modified Packets

### RequestNetworkSettingsPacket
packet id `193` · struct

- `ClientNetworkVersion` constraints changed from maximum=2168, minimum=2168 to maximum=2169, minimum=2169

```go
type RequestNetworkSettingsPacket struct {
	ClientNetworkVersion int32
}

func (*RequestNetworkSettingsPacket) ID() uint32 {
	return IDRequestNetworkSettingsPacket
}

func (pk *RequestNetworkSettingsPacket) Marshal(io protocol.IO) {
	io.BEInt32(&pk.ClientNetworkVersion)
}
```

## Modified Types

### TextDataPayload
struct · **wire break**

- added `LineGapHeight` at ordinal 3 (float, optional)
- `DepthTest` moved from ordinal 3 to 4
- `ShowBackface` moved from ordinal 4 to 5
- `ShowTextBackface` moved from ordinal 5 to 6

**Also update:** `PrimitiveShapesPacket` — it embeds this type.

```go
type TextData struct {
	// Text Text (string) of the debug text shape.
	Text string
	// UseRotation Whether to use specified rotation, otherwise faces camera.
	UseRotation bool
	// BackgroundColor Color to draw background.
	BackgroundColor Optional[Color]
	// LineGapHeight Line gap height for multiline text rendering.
	LineGapHeight Optional[float32]
	// DepthTest Whether other objects block seeing the text (false=like nametags,true=like signs).
	DepthTest bool
	// ShowBackface Whether the backface of the background should be rendered.
	ShowBackface bool
	// ShowTextBackface Whether the backface of the text should be rendered.
	ShowTextBackface bool
}

func (pk *TextData) Marshal(io IO) {
	io.String(&pk.Text)
	io.Bool(&pk.UseRotation)
	OptionalMarshaler(io, &pk.BackgroundColor)
	OptionalFunc(io, &pk.LineGapHeight, io.Float32)
	io.Bool(&pk.DepthTest)
	io.Bool(&pk.ShowBackface)
	io.Bool(&pk.ShowTextBackface)
}
```

## Modified Enums

### persona__AnimatedTextureType
enum

- added value `None` = 0

**Also update:** `PlayerListPacket`, `PlayerSkinPacket` — each embeds this type.

```go
const (
	PersonaAnimatedTextureTypeNone        = 0
	PersonaAnimatedTextureTypeFace        = 1
	PersonaAnimatedTextureTypeBody32x32   = 2
	PersonaAnimatedTextureTypeBody128x128 = 3
)
```
