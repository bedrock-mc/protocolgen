# protocolgen

Tooling that turns Mojang's [bedrock-protocol-docs](https://github.com/Mojang/bedrock-protocol-docs)
JSON into Go packet codecs, in two formats, and detects drift between the docs
and an existing hand-written implementation.

Extracted from the gophertunnel fork into its own repo. **Work in progress —
handed off; see "Status / what's left" below.**

> The docs are under the Minecraft EULA. Pass a local clone path; never vendor
> their content. The `testdata/` fixtures here are synthetic, not Mojang content.

## Layout

```
cmd/gophertunnel/        gophertunnel-format generator (+ -drift)  (package main)
cmd/raw/                 raw-format generator                       (package main)
internal/codegen/    shared generator library (both formats)
internal/driftcheck/ drift detector
```

Everything is pure-stdlib (it reads the docs as JSON and the target packets as Go
source via go/ast; it does not import gophertunnel).

## gophertunnel — maps onto gophertunnel's protocol.\* types

```sh
git clone --depth=1 https://github.com/Mojang/bedrock-protocol-docs /tmp/bpd
git -C /tmp/bpd fetch origin pull/33/head && git -C /tmp/bpd checkout FETCH_HEAD

# generate into a directory for inspection
go run ./cmd/gophertunnel -docs=/tmp/bpd/json -out=out_gophertunnel

# overwrite an existing gophertunnel packet dir in place for git-diff review
go run ./cmd/gophertunnel -docs=/tmp/bpd/json \
    -packets=/path/to/gophertunnel/minecraft/protocol/packet -overwrite

# report drift between the docs and a packet dir (no generation)
go run ./cmd/gophertunnel -docs=/tmp/bpd/json \
    -packets=/path/to/gophertunnel/minecraft/protocol/packet -drift
```

Maps composites onto existing `protocol.*` types (verified mapping table in
`internal/codegen/resolve.go`), preserves `[]byte`/wire-equivalent fields, and —
with `-overwrite` — merges into the hand-written packet files in place, keeping
their comments/consts/helpers and (with `-preserve-matching`, default on) keeping
gopher's field names where the wire is unchanged.

## raw — a self-contained package

```sh
go run ./cmd/raw -docs=/tmp/bpd/json -out=out_raw   # -pkg sets the package name (default "raw")
```

Generates a struct for every packet (doc-faithful names, the `Packet` suffix kept
to avoid colliding with composite types) and every composite definition,
marshalling against a generated `IO` interface (`protocol.go`). Any
referenced-but-undefined type is stubbed so the package always compiles.

## Status / what's left

Against PR #33 docs (1.26.40-beta.29): raw generates 229 packets + 250 composites
and **compiles**; the gophertunnel format leaves ~80 packets needing
hand-finishing.

Known gaps (the hard 20% the docs can't express, plus rough edges):

- **oneOf unions** (`Text` body, `InventoryTransaction`, camera instructions, …)
  → emitted as `any` + TODO; need hand-written discriminated marshal.
- **Optionals** are flattened to plain types (the doc's `required` is captured but
  not yet turned into an optional/bool-prefix on either format).
- **Composite inlining**: the docs nest groups of fields into composites
  (`StartGame`→`LevelSettings`, the whole `CameraInstruction` body); gophertunnel
  inlines them flat. Raw generates the composite types; gophertunnel leaves them
  as TODO references.
- **Raw Marshal** is best-effort: scalars, single composites, and slices are
  handled; nested length-prefix variants, optionals, and unions are not.
- **Naming**: raw uses Mojang's inconsistent doc names (`ContainerId` vs
  `RuntimeID`, etc.) verbatim. The gophertunnel format preserves gopher's names
  for wire-matched fields only.
- A few composite mappings were rejected by structural verification
  (`CommandOriginData`, `FullContainerName`) and stay TODO.
