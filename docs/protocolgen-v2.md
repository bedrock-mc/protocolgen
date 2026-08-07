# protocolgen v2

v2 is the canonical, fail-closed Bedrock wire pipeline. A source checkout is
lowered into complete machine-derived claims, claims are reconciled for one
locked protocol snapshot, and emitters consume only the resulting manifest:

```text
source lock + pinned source trees
        │
        ├─ raw Mojang JSON frontend
        └─ Endstone dump frontend
                    │
             complete claims + pins
                    │
        equal claims merge; disagreements stop
        unless an evidenced fingerprinted adjudication selects one
                    │
             canonical manifest schema v2
                 ┌──────┴──────┐
             Go emitter   Rust consumer
```

The manifest is a stable pretty-printed JSON document. It records packet
direction, ordered field ordinals, encode/decode shapes, symmetry, semantic
and named-type identities, exact primitives, explicit string/bytes prefixes,
arrays and fixed arrays, nested optionals, ordered maps, sequences, structs,
unions, conditionals, explicit enum values, reserved/ignored compatibility
nodes, recursive markers, and opaque/unresolved nodes. Reachable opaque or
unresolved nodes make validation and generation fail.

`schema_version: 2` is the wire vocabulary version. It is not a target profile:
Cargo layout, borrowed views, lossy strings, naming, and cross-version
deduplication belong in a downstream profile. The generated Rust and Go output
in this slice contains profile-neutral definitions and packet IDs. The
manifest is the only emitted wire-schema artifact, and the language outputs do
not claim packet codec coverage until they can emit real value serialization.

## Sources and version pins

`testdata/source-lock.json` demonstrates the lock format with synthetic
fixtures only. Real runs must supply local source checkouts and immutable
revision/tree digests; protocolgen never downloads, vendors, or embeds Mojang
schemas, BDS binaries, captures, Prismarine data, or proprietary corpora.
The lock and every frontend reject mixing Minecraft 1.26.40/protocol 2168 with
1.26.50/protocol 2169.

Raw Mojang JSON supplies order and declared semantics, but may have wrong
requiredness, enum values, or types. Endstone/BDS dumps supply an independent,
runtime-shaped skeleton, but C++ `std::string` is not automatically UTF-8.
Neither source silently wins. Known Mojang defects are represented directly as
v2 corrections, where every operation names its source locator, exact
pre-patch node hash, complete pre-patch context hash, post-patch hash, reason,
and evidence. Cloudburst, gophertunnel, PMMP, and corrected third-party views
may identify a discrepancy, but no proposed fix becomes canonical without a
pinned, reviewable wire-layout source. They are evidence or assistive inputs,
not generation dependencies or hidden precedence rules.

## Reconciliation and auditability

There are no confidence scores, majority votes, or implicit source precedence.
Missing evidence is not agreement. Equal complete claims inherit only the
sorted source pins. Different claims fail unless an adjudication contains:

- a context hash over the target and the complete disagreement claim set;
- a SHA-256 fingerprint for every competing source claim;
- a selected source and cited evidence; and
- a hand-authored reason.

Changing any field name, semantic identity, locator, primitive shape, nested
node, or source claim makes the fingerprint stale. The same rule applies to
fingerprinted corrections. This is why source claims include the full node,
not just a structural summary.

## Commands

From the repository root:

```sh
go run ./cmd/protocolgen reconcile \
  -lock testdata/source-lock.json \
  -mojang /path/to/bedrock-protocol-docs \
  -endstone /path/to/endstone-dump \
  -out /tmp/bedrock-2168-v2.json

go run ./cmd/protocolgen validate -manifest /tmp/bedrock-2168-v2.json
go run ./cmd/protocolgen emit-go -manifest /tmp/bedrock-2168-v2.json -out /tmp/wire-go -pkg wiregen
go run ./cmd/protocolgen emit-rust -manifest /tmp/bedrock-2168-v2.json -out /tmp/wire-rust
```

The reproducible synthetic vertical slice is:

```sh
go run ./cmd/protocolgen reconcile \
  -lock testdata/source-lock.json \
  -mojang testdata/sources/mojang-v2168 \
  -endstone testdata/sources/endstone-v2168 \
  -out /tmp/protocolgen-v2.json
go run ./cmd/protocolgen validate -manifest /tmp/protocolgen-v2.json
```

`internal/wire` runs both-direction synthetic goldens over the same node
vocabulary. The golden covers nested/double optionals, explicit array prefixes,
nested fixed arrays, skipped union ordinals, explicit enum ordinals, reserved
fields, text versus arbitrary bytes, and semantic-ID/wire-byte equivalence.

The small positive parity normalization is reproducible without a source
checkout:

```sh
go run ./cmd/protocolgen parity \
  -manifest testdata/parity/v2-small.json \
  -axolotl testdata/parity/axolotl-v1-small.json
```

## Axolotl migration gate

The Axolotl protocol-specific history is retained under
`migration/axolotl-protocol/` from `axolotl-stack` `origin/main`, filtered to
the parser/lowering, `bedrock_core` codec vocabulary, corrections, and
`tools/wire-conformance`. Data generation, gameplay data, Prismarine and
submodules were excluded. The lifted Rust is a reference and parity source;
v2 does not reverse-lower generated Rust.

Axolotl's public schema-v1 wire manifest is compared with an explicit
byte-equivalence normalization layer:

```sh
scripts/axolotl-parity.sh /path/to/axolotl-stack /tmp/protocolgen-v2.json
```

The command regenerates exactly one Axolotl `1.26.40` manifest externally and
then compares it. A mismatch is a mechanical blocker, not an automatic
agreement. On this first slice, full live 2168 parity remains unsupported for
shapes not represented by the adapter (notably raw/encapsulated buffers,
profile-specific recursive codecs, and asymmetric decode shapes); the command
reports the exact normalized difference so the next migration slice can pin
the correction rather than deleting the old generator.

The gate was attempted against Axolotl `origin/main`
(`781dfcb0ab443476b62df3c983750c0c1527a95a`) in a clean detached worktree. It
currently stops before parsing with the exact mechanical error
`crates/valentine_gen/bedrock-protocol-docs/json: No such file or directory`:
the pinned docs submodule is not initialized in that checkout. Initialize the
submodule in a disposable Axolotl checkout, or set `AXOLOTL_MOJANG_DOCS` to an
immutable local docs checkout, then rerun the same script. The older checked-out
Axolotl feature worktree is not a substitute for `origin/main` in this gate.

The preserved conformance machinery remains runnable at
`migration/axolotl-protocol/tools/wire-conformance/`. Its accepted divergence
file is part of the gate: divergence is accepted only when explicitly listed
with a reason, while unresolved and absent oracle packets remain non-agreement.

## Gaps deliberately left fail-closed

- This is not full 2168 packet coverage and no live/proprietary corpus is
  committed.
- NBT requires a bounded/profile codec in the synthetic interpreter; recursive
  nodes require a profile codec; conditional decode needs discriminator
  context.
- The Rust/Go emitters provide canonical shape/API consumers, not the complete
  Axolotl borrowed codec or gophertunnel profile. The latter remains a narrow
  `go/ast` assistive adapter in `internal/gophertunnel`.
- Existing `cmd/raw` and `cmd/gophertunnel` are preserved standalone WIP tools.
  They are not v2 inputs or emitters and do not weaken v2 validation.
