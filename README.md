# protocolgen

Tooling that turns Mojang's [bedrock-protocol-docs](https://github.com/Mojang/bedrock-protocol-docs)
JSON into Go packet codecs, and detects drift between the docs and an existing
hand-written implementation.

Extracted from the gophertunnel fork into its own repo. **Work in progress —
handed off; see "Status / what's left" below.**

> The docs are under the Minecraft EULA. Pass a local clone path; never vendor
> their content. The `testdata/` fixtures here are synthetic, not Mojang content.

## Layout

```
gen/      protocolgen — the generator (package main)
drift/    protocoldrift — the drift detector (package main)
```

Both are pure-stdlib (they read the docs as JSON and the target packets as Go
source via go/ast; they do not import gophertunnel).

## gen — the generator

```sh
git clone --depth=1 https://github.com/Mojang/bedrock-protocol-docs /tmp/bpd
git -C /tmp/bpd fetch origin pull/33/head && git -C /tmp/bpd checkout FETCH_HEAD

# gophertunnel format (default): maps onto gophertunnel's protocol.* types
go run ./gen -format=gophertunnel -docs=/tmp/bpd/json -out=out_gt

# raw format: a self-contained package, no gophertunnel dependency
go run ./gen -format=raw -docs=/tmp/bpd/json -out=out_raw   # package name defaults to "raw"
```

It is built around a `Format` interface (`gen/format.go`) with two strategies:

- **`gophertunnel`** (`resolve.go`, `emit.go`, `merge.go`): maps composites onto
  existing `protocol.*` types (verified mapping table in `resolve.go`), preserves
  `[]byte`/wire-equivalent fields, and — with `-overwrite` — merges into the
  hand-written packet files in place, keeping their comments/consts/helpers and
  (with `-preserve-matching`, default on) keeping gopher's field names where the
  wire is unchanged. Intended for `git diff` review of a protocol upgrade.
- **`raw`** (`format_raw.go`): generates a self-contained package — a struct for
  every packet (doc-faithful names, the `Packet` suffix kept to avoid colliding
  with composite types) and every composite definition, marshalling against a
  generated `IO` interface (`protocol.go`). Any referenced-but-undefined type is
  stubbed so the package always compiles.

## drift — the drift detector

Compares the docs against an existing gophertunnel packet directory (by packet
id, aligning the doc fields against the `Marshal` op sequence) and reports
new/removed packets and field/encoding mismatches.

```sh
go run ./drift -docs=/tmp/bpd/json -packets=/path/to/gophertunnel/minecraft/protocol/packet
```

## Status / what's left

Against PR #33 docs (1.26.40-beta.29): raw generates 229 packets + 250 composites
and **compiles**; gophertunnel format leaves ~80 packets needing hand-finishing.

Known gaps (the hard 20% the docs can't express, plus rough edges):

- **oneOf unions** (`Text` body, `InventoryTransaction`, camera instructions, …)
  → emitted as `any` + TODO; need hand-written discriminated marshal.
- **Optionals** are flattened to plain types (the doc's `required` is captured but
  not yet turned into an `Optional`/bool-prefix on either format).
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
