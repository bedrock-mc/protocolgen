# protocolgen

Generates gophertunnel-format packet codecs from Mojang's
[bedrock-protocol-docs](https://github.com/Mojang/bedrock-protocol-docs) JSON,
writing one `.go` file per packet into an output directory so the result can be
**diffed against the hand-written packets**.

This is the companion to [`protocoldrift`](../protocoldrift): drift *detects*
divergence, gen *produces* the candidate upgrade. The intended workflow is to
generate against a new protocol version and diff generated-vs-manual — the diff
is your upgrade review.

## Usage

```sh
git clone --depth=1 https://github.com/Mojang/bedrock-protocol-docs /path/to/docs
go run ./cmd/protocolgen \
  -docs=/path/to/docs/json \
  -packets=minecraft/protocol/packet \
  -out=generated

# review a packet:
diff generated/animate.go minecraft/protocol/packet/animate.go
```

| flag | default | meaning |
|------|---------|---------|
| `-docs` | (required) | docs `json/` directory of a local clone |
| `-packets` | `minecraft/protocol/packet` | gophertunnel packet dir (read for `id.go`, so `ID()` uses the right `IDXxx` constant) |
| `-out` | `generated` | output directory (non-destructive; never touches the real packets) |

Output is one file per packet (`<snake>.go`), each with the struct, `ID()`, and a
single bidirectional `Marshal(io protocol.IO)` — gophertunnel's exact shape.

## What it generates cleanly vs. flags

Clean (no TODO): scalars (with varint/zigzag/LE/compression encoding), known
composites (`Vec2`/`Vec3`/`BlockPos`/`ChunkPos`/`SubChunkPos`/`UUID`, including
namespaced titles like `mce::UUID`), plain strings/bools, and slices of those.

Emitted with a `// TODO(protocoldrift)` marker (the docs genuinely can't express
these at field level):

- `oneOf` **unions** — discriminated marshal must be hand-written
- **NBT** blobs — docs model them as string blobs; encoding is guessed
- **ItemStack / Recipe / EntityMetadata** and other engine composites — map to the
  existing `protocol` type or hand-write
- **optional** fields (absent from `required`) — may be a `protocol.Optional`
- enums serialized as their **string name** — need value↔string conversion
- bare `integer`/`number` with no `x-underlying-type` — ambiguous width/encoding

Every emitted file is **syntactically valid Go** (gofmt-checked). A clean packet
is effectively a drop-in body for the real packet (modulo field naming, which
follows the doc names, e.g. `PlayerRuntimeId` vs gopher's `EntityRuntimeID`).

The run prints a summary: how many packets are clean, how many need
hand-finishing, and the TODO reasons by frequency.

## Known limitations (v1)

- **Generated field names follow the docs**, not gopher's conventions — so a
  clean packet's diff still shows naming differences. Intentional for v1 (you see
  every difference); a future name-aliasing pass could match gopher's names.
- **Composite sub-types are referenced, not generated** — `[]AttributeData`
  refers to a type you must provide/map. A future pass could emit those structs.
- Array count type is assumed `varuint32` (the Bedrock default).
- Only the current-gopher output format is implemented; a second, more reusable
  format is planned.

## Licensing

Mojang's docs are under the **Minecraft EULA**, not OSS. Pass a local clone path
via `-docs`; this tool never vendors their content. `testdata/` fixtures are
synthetic.
