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
  -directions /path/to/reviewed-directions.json \
  -out /tmp/protocol-2168.json

go run ./cmd/protocolgen validate \
  -manifest /tmp/protocol-2168.json

go run ./cmd/protocolgen emit-go \
  -manifest /tmp/protocol-2168.json \
  -naming /path/to/naming.json \
  -out /tmp/protocol-go \
  -protocol-import example.com/project/protocol

go run ./cmd/protocolgen emit-rust \
  -manifest /tmp/protocol-2168.json \
  -naming /path/to/naming.json \
  -out /tmp/protocol-rust

# Verify the canonical manifest against the pinned gophertunnel source oracle.
# Omit -gophertunnel to let protocolgen clone the full-SHA lock into its cache.
go run ./cmd/protocolgen verify-gophertunnel \
  -manifest generated/1.26.40/manifest.json \
  -gophertunnel /path/to/gophertunnel \
  -report /tmp/gophertunnel-2168-report.json

# Diff two corrected schema snapshots into a human-readable changelog.
go run ./cmd/protocolgen changelog \
  -from /path/to/previous/bpd-fixer/output/json \
  -to /path/to/target/bpd-fixer/output/json \
  -from-branch automated/1.26.40 \
  -to-branch r/26_u4 \
  -from-upstream 0e00fe80f4 \
  -to-upstream 0f6a0bff19 \
  -from-fixer 9fa8eb75f9 \
  -to-fixer 9fa8eb75f9 \
  -out /tmp/changelog.md

# Turn that generated changelog into implementation-oriented Go snippets for
# every changed packet, shared type, and enum.
go run ./cmd/protocolgen update-guide \
  -changelog /tmp/changelog.md \
  -schemas /path/to/bpd-fixer/output/json \
  -out /tmp/gophertunnel-update.md
```

This produces:

- a deterministic JSON manifest containing the canonical wire protocol;
- schema-published validation constraints (`min`/`max` lengths, collection and
  map counts, numeric bounds, and string patterns) retained independently of
  wire shape and enforced by generated encoders and decoders;
- typed Go packet structures, one packet per file, with shared semantic types,
  idiomatic enums, closed union interfaces, ordered map entries, and real
  symmetric `Marshal(IO)` methods that operate on packet values. The generated
  package also includes reusable `IntegerFunc`, `OptionalFunc`,
  `DoubleOptionalFunc`, `FuncSlice`, and `OrderedMap` helpers plus bounded
  in-memory `Reader` and `Writer` implementations. Readers cap decoded
  collections by default and expose `NewReaderWithLimit` when an application
  needs a different bound. The packet package also provides a `Packet`
  interface, generated ID methods, and direction-aware constructor pools
  (`NewPacket`, `NewClientPacket`, and `NewServerPacket`);
- typed Rust packet structures, one packet module per file, with shared types,
  native enums whose total `From<integer>` conversion preserves unrecognised
  discriminants, symmetric `Encode`/`Decode` implementations over a slice-based
  runtime, ordered map tuples, and payload-bearing enums for Cereal unions; and
- packet IDs in both language outputs.

Packet APIs use concise target-language names: the schema's transport-oriented
`Packet` and `PacketPayload` suffixes do not leak into public type names. Go
emits a reusable `protocol` package under `protocol/` for runtime codecs and
shared definitions, and a `protocol/packet` subpackage for packet structs and
IDs. Packet files import the protocol package using the path supplied through
`-protocol-import`. Rust keeps its shared definitions in `types.rs` and
`enums.rs` and its packet modules under `src/packets`. The canonical manifest
is the sole wire-schema artifact rather than being duplicated as large runtime
descriptors in every generated language.

Like sqlc's type overrides, each backend maps well-known wire semantics onto
types native to its ecosystem without changing the manifest. Go uses
`uuid.UUID`, `mgl32.Vec2`/`Vec3`, `color.RGBA`, and a value-based
`Optional[T]` that preserves absent versus present-zero state without pointer
nesting. Cereal double optionals retain both markers in the manifest but use
the same single public `Optional[T]` state as gophertunnel. Rust uses
`uuid::Uuid`, `glam::Vec2`/`Vec3`, zero-copy-ready `bytes::Bytes` buffers, and
an explicit `Nbt` byte wrapper. Unknown
or protocol-specific structures remain generated named types rather than being
guessed into an unrelated library type. Go likewise leaves undifferentiated
NBT as `[]byte`, because the manifest's `nbt_le` primitive may be a compound or
an intentionally opaque/loose tag stream.

Native Go mappings are a target profile, not part of the canonical schema. The
CLI enables them by default; pass `-native-types=false` when a consumer wants
only generated named wire structs. Packet runtime helpers and factory pools can
similarly be omitted with `-packet-runtime=false` and `-packet-pools=false`.

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
  -directions generated/1.26.40/directions.json \
  -out generated/1.26.40/manifest.json

go run ./cmd/protocolgen emit-go \
  -manifest generated/1.26.40/manifest.json \
  -naming generated/1.26.40/naming.json \
  -out generated/1.26.40/go \
  -protocol-import protocolgen/generated/1.26.40/go/protocol

go run ./cmd/protocolgen emit-rust \
  -manifest generated/1.26.40/manifest.json \
  -naming generated/1.26.40/naming.json \
  -out generated/1.26.40/rust
```

The checked-in snapshot can be regenerated and checked for drift with the
same commands through the Makefile. Set the local Mojang and Endstone checkout
paths; the Go commands use `/tmp/go-build-cache` locally.

```sh
MOJANG_DIR=/private/tmp/claude-501/-Users-hashim-Coding-Go-Lunar/a6989ec5-1fd8-4136-ad72-5e4a0665aca1/scratchpad/mojang-docs/json \
ENDSTONE_DIR=/private/tmp/endstone-protocol-docs.nTILn9 \
GOCACHE=/tmp/go-build-cache make regen

MOJANG_DIR=/private/tmp/claude-501/-Users-hashim-Coding-Go-Lunar/a6989ec5-1fd8-4136-ad72-5e4a0665aca1/scratchpad/mojang-docs/json \
ENDSTONE_DIR=/private/tmp/endstone-protocol-docs.nTILn9 \
GOCACHE=/tmp/go-build-cache make verify
```

`make regen` omits `GOPHERTUNNEL_DIR` intentionally; the oracle clones the
locked full SHA into its cache. The checked-in `naming.json` is passed to both
emitters.

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
- source validation bounds remain attached to the exact scalar or collection
  they constrain;
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
| `update-guide` | Render target-version gophertunnel transcription snippets for definitions named by a protocol changelog. |

### Adding a language backend

Language backends implement `internal/emitter.Backend`. The shared runner loads
and validates the canonical manifest plus the reviewed naming, domain, and
documentation overlays, then passes them to the backend as one `emitter.Input`.
The backend returns relative filenames and contents; the runner writes them in
deterministic order, rejects paths outside the output directory, removes stale
generated files, and reports documentation coverage.

Keep target-language type selection, naming rules, runtime support, and codec
generation inside the backend. The canonical manifest remains the only shared
wire-schema IR; a backend must not infer wire shape from another language's
generated output.

`update-guide` deliberately takes the human-readable changelog as its change
inventory. It does not guess that every textual schema difference changes the
wire. For each packet, shared type, or enum heading in the changelog, it resolves
the target corrected schema, preserves ordinal field order and optionality, and
emits a complete target-version struct, `Marshal`, packet `ID`, or enum constant
block. Removed definitions receive removal guidance instead of a target snippet.
The command rejects a schema snapshot whose internal protocol metadata does not
match the changelog target, as well as unsupported codecs or ambiguous layouts.
The generated Go is a transcription aid: schema names may differ from the
established gophertunnel API, and maintainers must still reconcile names and
source evidence before applying it.

`verify-gophertunnel` reads `tools/gophertunnel-oracle/lock.json`, whose oracle
commit is a full 40-character SHA. Pass `-gophertunnel` for an existing
checkout at that exact commit, or omit it to clone into the platform user cache
at runtime. The locked repository may be a fork; only the exact locked SHA is
trusted, and the checkout is rejected if `HEAD` is anything else. The command
parses gophertunnel with `go/ast`; it does not import the checkout, download
its modules, or compare generated output. The manifest is the only protocol
shape being verified.

The oracle is not canonical. gophertunnel is a hand-written third-party
implementation used as one independent verification axis, exactly like the
Axolotl parity gate. It never edits the manifest: `verify-gophertunnel` writes
only its report, and no correction, adjudication, or source claim is ever
derived from it automatically. When the oracle and the manifest disagree, the
resolution is a pinned, reviewable wire-layout source recorded under
`generated/<version>/corrections/`, not agreement with gophertunnel.

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
wire-positioned bytes, fixed-array wrapper grouping by its repeated scalar wire
values, and the four named pre-encoded NBT byte fields. Width, endianness,
fixed versus varint, varint versus zigzag, float versus integer, option
presence, array prefixes, scalar counts in fixed arrays, and union discriminants
remain distinct. Reviewed exceptions live in
`tools/gophertunnel-oracle/accepted-divergences.json`; every entry requires a
reason, evidence locator, and a concrete `what_would_settle_it` statement. An
entry only silences the exit code for one packet; it is a record of an open
question, not a decision that gophertunnel is right. Entries whose packets no
longer diverge are reported as `resolved_accepted` so the baseline cannot keep
accepting differences that no longer exist.

## Current scope

The full pipeline—ingestion, reconciliation, validation, and Go/Rust
emission—is implemented.

Both emitters generate packet definitions, semantic types, enums, unions, and
packet IDs. The Go backend also generates a real symmetric `Marshal(IO)` method
for every packet and reachable shared type. Its small `IO` contract names exact
wire operations—including UUID half-endianness, NBT, bounded bitsets, strings,
bytes, collections, and invalid-value handling—so a gophertunnel adapter or a
standalone reader/writer can supply byte transport without shipping a JSON
schema interpreter. Union decoding constructs the selected concrete variant;
repeated payload shapes receive distinct wrappers so their wire tags remain
representable.

The Rust backend emits value-aware `Encode` and fallible `Decode`
implementations over a slice-based runtime, so decoding bounds every declared
length against the bytes actually remaining before it reserves. It carries a
typed `DecodeError`, both NBT scanners, seven-bit bitsets, and direction-aware
packet registration: `Packet::decode_from` rejects an id the sending peer may
not use before reading a field. The generator fails instead of emitting generic
fallback types for unsupported sequence or unresolved nodes.

Collection and byte-buffer bounds are reader settings rather than hard caps, so
a peer that legitimately exceeds the default is recoverable:

```rust
let mut reader = wire::Reader::from_shared(&payload);
reader.set_collection_limit(16_384);
let packet = packets::Packet::decode_from(id, packets::Peer::Server, &mut reader)?;
```

`Reader::from_shared` decodes byte-buffer and NBT fields as refcounted slices of
the source buffer instead of copies, which matters most on the chunk payloads.
Borrowed string views are not emitted yet, so a decoded packet still owns its
strings. Codebase-specific merge rules also remain future backend work.

The superseded v1 experiments are not inputs to the canonical pipeline; the
historical Axolotl protocol work remains under `migration/`.

See [docs/protocolgen-v2.md](docs/protocolgen-v2.md) for the manifest model,
source policy, adjudication format, Axolotl migration gate, and known gaps.

> Mojang protocol documentation is governed by the Minecraft EULA. Supply it
> as a local input; do not commit it to this repository. All committed fixtures
> are synthetic.
