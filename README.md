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
```

This produces:

- a deterministic JSON manifest containing the canonical wire protocol;
- typed Go packet structures with wire-layout descriptors generated code
  cannot alter; and
- typed Rust packet structures with the same wire-layout descriptors.

`generated/1.26.40/` contains the checked-in protocol 2168 source lock,
canonical 223-packet manifest, and matching Go and Rust outputs generated from
the pinned Endstone Cereal dump.

For real generation, replace the fixture paths with immutable local Mojang and
Endstone checkouts and record their revisions and directory hashes in the
source lock. Protocol data is never downloaded or vendored by this repository.

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

## Current scope

The full pipeline—ingestion, reconciliation, validation, and Go/Rust
emission—is implemented.

The current emitters generate typed packet APIs that delegate wire work to an
encoder/decoder supplied by the target codebase. They do **not yet** generate
complete, drop-in [gophertunnel](https://github.com/Sandertv/gophertunnel)
`Marshal(protocol.IO)` implementations or the full borrowed
[Axolotl](https://github.com/axolotl-stack/axolotl-stack) codec API. NBT,
recursive target-specific codecs, conditional decode context, packet
registration, and codebase-specific merge rules are the remaining backend
work.

The older `cmd/gophertunnel` and `cmd/raw` experiments remain available, but
they are not inputs to the canonical pipeline and cannot weaken its validation.

See [docs/protocolgen-v2.md](docs/protocolgen-v2.md) for the manifest model,
source policy, adjudication format, Axolotl migration gate, and known gaps.

> Mojang protocol documentation is governed by the Minecraft EULA. Supply it
> as a local input; do not commit it to this repository. All committed fixtures
> are synthetic.
