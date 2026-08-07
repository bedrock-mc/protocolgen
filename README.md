# protocolgen

protocolgen turns multiple Minecraft Bedrock protocol descriptions into one
reviewable source of truth, then generates Go and Rust packet types from it. It
is built for developers working on Bedrock servers, proxies, and tooling.

Protocol sources are imperfect. Mojang's documentation may be mislabeled or
incomplete, while projects such as [Endstone](https://github.com/EndstoneMC/endstone)
may describe the same field differently. A normal generator must silently
trust one of them. protocolgen does not: it pins every input, compares their
wire layouts, and stops when they disagree. A human can resolve a conflict only
with a documented, fingerprinted decision that becomes invalid if the source
changes.

```text
Pinned Mojang + Endstone inputs
               |
        independent claims
               |
      compare every wire field
       /                   \
  agreement          conflict: stop
       |                    |
       |        evidenced adjudication
       \____________________/
               |
       canonical manifest
           /          \
     Go output      Rust output
```

The result is reproducible and auditable: both language backends receive the
same packet order, integer encoding, optionals, arrays, unions, enums, reserved
fields, and encode/decode layouts. Language-specific code cannot quietly
reinterpret the protocol.

## Try it

Requires Go 1.26. The repository includes a small synthetic protocol fixture,
so the complete flow can be run without downloading Mojang data or installing
anything else:

```sh
go run ./cmd/protocolgen reconcile \
  -lock testdata/source-lock.json \
  -mojang testdata/sources/mojang-v2168 \
  -endstone testdata/sources/endstone-v2168 \
  -out /tmp/protocol-2168.json

go run ./cmd/protocolgen validate \
  -manifest /tmp/protocol-2168.json

go run ./cmd/protocolgen emit-go \
  -manifest /tmp/protocol-2168.json \
  -out /tmp/protocol-go \
  -pkg wiregen

go run ./cmd/protocolgen emit-rust \
  -manifest /tmp/protocol-2168.json \
  -out /tmp/protocol-rust

# Verify the canonical manifest against the pinned gophertunnel source oracle.
# Omit -gophertunnel to let protocolgen clone the full-SHA lock into its cache.
go run ./cmd/protocolgen verify-gophertunnel \
  -manifest generated/1.26.40/manifest.json \
  -gophertunnel /path/to/gophertunnel \
  -report /tmp/gophertunnel-2168-report.json
```

This produces:

- a deterministic JSON manifest containing the canonical wire protocol;
- typed Go packet structures, one packet per file, with shared semantic types,
  enums, closed union interfaces, and ordered map entries;
- typed Rust packet structures, one packet module per file, with shared types,
  native enums, checked `TryFrom<integer>` decoding, ordered map tuples, and
  payload-bearing enums for Cereal unions; and
- packet IDs in both language outputs.

Packet APIs use concise target-language names: the schema's transport-oriented
`Packet` and `PacketPayload` suffixes do not leak into public type names. Both
backends keep common definitions in `types` and `enums`; packet files contain
only the packet definition. The canonical manifest is the sole wire-schema
artifact rather than being duplicated as large runtime descriptors in every
generated language.

Like sqlc's type overrides, each backend maps well-known wire semantics onto
types native to its ecosystem without changing the manifest. Go uses
`uuid.UUID`, `mgl32.Vec2`/`Vec3`, and `color.RGBA`. Rust uses
`uuid::Uuid`, `glam::Vec2`/`Vec3`, and an explicit `Nbt` byte wrapper. Unknown
or protocol-specific structures remain generated named types rather than being
guessed into an unrelated library type. Go likewise leaves undifferentiated
NBT as `[]byte`, because the manifest's `nbt_le` primitive may be a compound or
an intentionally opaque/loose tag stream.

`generated/1.26.40/` contains the checked-in protocol 2168 source lock,
canonical 229-packet Cereal manifest, and matching Go and Rust outputs generated
by reconciling the pinned raw Mojang and Endstone dumps. Its `corrections/`
directory contains
fingerprinted, evidence-backed fixes for source defects; each correction stops
applying if the pinned source changes. The Mojang pin points at the raw official
JSON. protocolgen applies its own corrections while ingesting that source; it
does not require a preprocessed checkout or bpd-fixer. The checked-in
`adjudications.json` records the small set of disagreements that require exact
serializer evidence instead of pretending one documentation source always wins.

For real generation, replace the fixture paths with immutable local Mojang and
Endstone checkouts and record their revisions and directory hashes in the
source lock. Protocol data is never downloaded or vendored by this repository.

The checked 1.26.40 snapshot is regenerated with:

```sh
go run ./cmd/protocolgen reconcile \
  -lock generated/1.26.40/source-lock.json \
  -mojang /path/to/bedrock-protocol-docs/json \
  -mojang-corrections generated/1.26.40/corrections/mojang \
  -endstone /path/to/endstone-protocol-docs \
  -endstone-corrections generated/1.26.40/corrections/endstone \
  -adjudications generated/1.26.40/adjudications.json \
  -out generated/1.26.40/manifest.json

go run ./cmd/protocolgen emit-go \
  -manifest generated/1.26.40/manifest.json \
  -out generated/1.26.40/go \
  -pkg protocol2168

go run ./cmd/protocolgen emit-rust \
  -manifest generated/1.26.40/manifest.json \
  -out generated/1.26.40/rust
```

The raw Mojang side can be inspected independently with:

```sh
go run ./cmd/protocolgen ingest \
  -kind mojang \
  -root /path/to/bedrock-protocol-docs/json \
  -lock generated/1.26.40/source-lock.json \
  -id mojang \
  -corrections generated/1.26.40/corrections/mojang \
  -out /tmp/mojang-2168-claims.json
```

## Why the manifest matters

The manifest is the boundary between protocol research and code generation.
Sources may be wrong or incomplete; the manifest must be explicit,
version-locked, validated, and reviewable.

Before a manifest is written, protocolgen checks:

- Minecraft and network protocol versions match every source pin;
- field order, widths, signedness, compression, and length prefixes agree;
- strings and arbitrary bytes remain distinct;
- optional nesting and union selectors are explicit;
- enum values are complete and numeric;
- encode and decode layouts are represented independently when necessary;
- unresolved wire nodes cannot reach generated packets; and
- corrections and adjudications still match the exact source content they
  were written for.

Generated code consumes only this validated manifest. It cannot consult the
original docs, apply hidden source precedence, or maintain a second copy of the
wire schema.

## Commands

`cmd/protocolgen` provides the canonical pipeline:

| Command | Purpose |
| --- | --- |
| `hash-source` | Hash a local source tree for an immutable source lock. |
| `ingest` | Lower one source into inspectable claims. |
| `reconcile` | Compare claims and write a canonical manifest. |
| `validate` | Validate a manifest and its fingerprints. |
| `emit-go` | Generate Go packet types from a manifest. |
| `emit-rust` | Generate Rust packet types from a manifest. |
| `parity` | Compare the manifest with an independently generated Axolotl layout. |
| `verify-gophertunnel` | Compare the manifest with the pinned gophertunnel `Marshal` source oracle. |

`verify-gophertunnel` reads `tools/gophertunnel-oracle/lock.json`, whose oracle
commit is a full 40-character SHA. Pass `-gophertunnel` for an existing
checkout at that exact commit, or omit it to clone into the platform user cache
at runtime. The command parses gophertunnel with `go/ast`; it does not import
the checkout, download its modules, or compare generated output. The manifest
is the only protocol shape being verified.

It writes a machine-readable report with one result per manifest packet and
prints the same run's human summary:

```text
AGREEMENT / DIVERGENCE / UNRESOLVED / NO_ORACLE_PACKET
```

Only an unaccepted `DIVERGENCE` makes the command exit non-zero. Packets whose
marshal hides bytes behind runtime branches or opaque interface helpers remain
`UNRESOLVED`; they are never treated as agreement. The comparison normalizes
only the documented byte-equivalences: signed/unsigned fixed-width integers of
the same width and endianness, strings versus byte slices with the same length
prefix, prefixed arrays of one-byte elements versus byte slices, UUIDs as 16
wire-positioned bytes, and the four named pre-encoded NBT byte fields. Width,
endianness, fixed versus varint, varint versus zigzag, float versus integer,
option presence, array prefixes, fixed-array lengths, and union discriminants
remain distinct. Reviewed exceptions live in
`tools/gophertunnel-oracle/accepted-divergences.json`; every entry requires a
reason, evidence locator, and a concrete `what_would_settle_it` statement.

## Current scope

The full pipeline—ingestion, reconciliation, validation, and Go/Rust
emission—is implemented.

The current emitters generate packet definitions, semantic types, enums,
unions, and packet IDs. They deliberately do **not** expose placeholder codec
methods. Complete, drop-in
[gophertunnel](https://github.com/Sandertv/gophertunnel)
`Marshal(protocol.IO)` implementations and the full borrowed
[Axolotl](https://github.com/axolotl-stack/axolotl-stack) codec API remain
future backend work, along with NBT and
recursive target-specific codecs, conditional decode context, packet
registration, and codebase-specific merge rules.

The older `cmd/gophertunnel` and `cmd/raw` experiments remain available, but
they are not inputs to the canonical pipeline and cannot weaken its validation.

See [docs/protocolgen-v2.md](docs/protocolgen-v2.md) for the manifest model,
source policy, adjudication format, Axolotl migration gate, and known gaps.

> Mojang protocol documentation is governed by the Minecraft EULA. Supply it
> as a local input; do not commit it to this repository. All committed fixtures
> are synthetic.
