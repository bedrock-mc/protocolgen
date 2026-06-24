# protocoldrift

A drift detector that compares gophertunnel's hand-written packet codecs against
Mojang's official [bedrock-protocol-docs](https://github.com/Mojang/bedrock-protocol-docs)
JSON schema, and reports where the two have diverged.

It is a **validator, not a generator**. The hand-written packets remain the
source of truth; this tool tells a maintainer, each Minecraft release, *which
packets to look at* — new/removed packets, field-count divergences, and
wire-encoding mismatches.

## Why this exists

Mojang's docs (as of the PR that filled in the remaining complex packets) ship a
machine-readable JSON-Schema file per packet:

- `x-ordinal-index` — the field's position on the wire
- `x-underlying-type` — the concrete scalar (`uint8`, `int32`, `float`, …)
- `x-serialization-options` — e.g. `["Compression"]` (varint/zigzag), `["Enum-as-Value"]`
- `$metaProperties["[cereal:packet]"]` — the numeric packet id

gophertunnel encodes each packet with a single bidirectional `Marshal(io protocol.IO)`
method (the same code reads and writes). protocoldrift parses those methods as
**source via `go/ast`** — it never imports the packet package, so it works across
protocol versions without compiling the target — extracts the ordered wire ops,
and compares them to the docs.

## Usage

```sh
# Clone the docs locally first (do NOT vendor them — see Licensing below).
git clone --depth=1 https://github.com/Mojang/bedrock-protocol-docs /path/to/docs

go run ./cmd/protocoldrift \
  -docs=/path/to/docs/json \
  -packets=minecraft/protocol/packet
```

Flags:

| flag | default | meaning |
|------|---------|---------|
| `-docs` | (required) | path to the docs `json/` directory of a local clone |
| `-packets` | `minecraft/protocol/packet` | gophertunnel's packet directory |
| `-verbose` | `false` | show every field, not just findings |
| `-json` | `false` | emit the full report as JSON |
| `-fail-on-drift` | `true` | exit non-zero when drift is detected (for CI) |

Exit codes: `0` no drift · `1` drift found · `2` error.

## How findings are classified

Per packet, doc fields (sorted by `x-ordinal-index`) are aligned positionally
with the source's wire ops:

- **ok** — the source op is an acceptable encoding for the documented field
  (e.g. `uint64`+`Compression` → `io.Varuint64`, `float` → `io.Float32`,
  composite `Vec3` → `io.Vec3`).
- **mismatch** — the encodings disagree (e.g. doc says `uint32`+`Compression`
  but the source uses fixed `io.Uint32`). **This is the high-confidence drift
  signal** and is always reported, even inside conditional packets.
- **review** — a structural op (`protocol.Slice`/`OptionalFunc`/…) or handcoded
  call that a human must judge, and alignment overflow inside conditional
  packets (where positional alignment is unreliable).
- **missing-in-source / extra-in-source** — field-count divergence in a *linear*
  packet (no conditionals); counts as drift.

A type is treated as a packet only if it declares both `ID()` and `Marshal()`;
composite sub-types (e.g. `PartyInfo`) and runtime-id packets (`Unknown`) are
skipped.

## Known limitations (v1)

- **Discriminated-union / type-switched packets** (e.g. `Text`,
  `InventoryTransaction`) can't be aligned reliably by position. The tool
  deliberately downgrades their alignment findings to `review` rather than
  asserting false drift — encoding mismatches in their linear parts are still
  caught. Read these with `-verbose` and human judgment.
- **NBT, EntityMetadata, ItemStack, Recipe, StackRequestAction** are not
  first-class in the docs (NBT is modelled as a prose length-prefixed string),
  so they surface as `review`.
- Alignment is positional, not name-based (doc/source field names rarely match,
  e.g. "Target Actor Runtime ID" ↔ `EntityRuntimeID`). A future baseline file
  could suppress known-complex packets to keep CI quiet.

## Licensing

Mojang's `bedrock-protocol-docs` is published under the **Minecraft EULA**, not
an OSS license. This tool reads the docs from a **local clone path you pass in**
and never vendors their content into this repository. The fixtures under
`testdata/` are synthetic examples written for this tool, not Mojang content.
