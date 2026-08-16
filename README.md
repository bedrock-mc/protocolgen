# protocolgen

Turns Mojang and Endstone's Bedrock protocol docs into one version-locked
manifest, then generates typed Go and Rust packet code from it.

The two doc sources disagree with each other and are each individually wrong
sometimes. protocolgen pins both, diffs every wire field, and refuses to
guess: a disagreement is resolved only by a fingerprinted correction or
adjudication, which stops applying the moment the source it was written
against changes. Emitters read the resulting manifest only — never the raw
docs.

## Commands

`cmd/protocolgen`:

| Command | Purpose |
| --- | --- |
| `hash-source` | Hash a local source tree for an immutable source lock. |
| `ingest` | Lower one source into inspectable claims. |
| `reconcile` | Compare claims and write a canonical manifest. |
| `validate` | Validate a manifest and its fingerprints. |
| `emit-go` | Generate Go packet types from a manifest. |
| `emit-rust` | Generate Rust packet types from a manifest. |
| `parity` | Compare the manifest with an independently generated Axolotl layout. |
| `verify-gophertunnel` | Compare the manifest with the pinned gophertunnel source oracle. |
| `changelog` | Diff two corrected schema snapshots into a Markdown changelog. |
| `update-guide` | Turn a changelog into gophertunnel transcription snippets. |
| `hotfix` | Derive a fingerprinted same-protocol patch from a reconciled manifest. |

Run any command with `-h` for its flags.

`vanilla-data/cmd/vanilla-data` is the companion BDS capture bot. It follows
the same offline gophertunnel login flow as `df-mc/datagen`, but stores raw,
version-locked packet evidence rather than Dragonfly- or PocketMine-specific
output. It is a separate Go module so its game-client dependencies do not leak
into the protocol generator.

## Regenerating the 1.26.40 snapshot

`generated/1.26.40/` is the checked-in manifest and matching Go/Rust output
for protocol 2168, built from pinned Mojang and Endstone checkouts. Mojang's
docs are EULA-restricted and not vendored here, so regenerating needs local
checkouts of both (requires Go 1.26):

```sh
MOJANG_DIR=/path/to/bedrock-protocol-docs/json \
ENDSTONE_DIR=/path/to/protocol-docs \
make regen
```

`make verify` runs the same pipeline and fails if it produces drift from
what's checked in — this is what CI enforces.

## Regenerating the 1.26.44 same-ID hotfix

Minecraft 1.26.44 retained protocol 2168 but added one outer presence marker
around `RemoveScore.ObjectiveName`. Because no second complete 1.26.44 source
snapshot exists, `generated/1.26.44/hotfix.json` derives the release from the
fully reconciled 1.26.40 manifest. The spec pins the complete base-manifest
hash, the exact node hash, the target codec evidence, and one constrained
`wrap_optional` operation. Run:

```sh
make hotfix
```

The derivation fails closed if the base manifest or target node changes. It
does not relax normal reconciliation or allow arbitrary manifest replacement.

## Capturing vanilla BDS data

Vanilla data is evidence alongside a generated protocol, not an input to wire
reconciliation. The capture includes actor identifiers, biomes, recipes,
creative content, items, dimensions, features, camera presets, trims, voxel
shapes. Dimensions and features are optional because vanilla BDS does not send
them for every world; `capture.json` explicitly records whether they were
captured or absent. It also records the target, BDS archive and executable
SHA-256 values, server settings, exact gophertunnel build, and every output
digest. `StartGame.CustomBlocks` is intentionally not exported: on an
addon-free server it is empty and is not the vanilla runtime block palette.

For the pinned local BDS configured as described by the version's source lock:

```sh
make vanilla-data \
  BDS_BINARY=/absolute/path/to/bedrock_server
```

The bot verifies `server.properties` beside that executable before connecting.

The `Protocol update BDS data` workflow is deliberately a post-correction
stage, not an unconditional continuation of ingestion. After source
reconciliation and any manual corrections or adjudications are complete, its
first job validates and builds the corrected generated manifest and exact
updated gophertunnel codec. Only then may the capture job start the matching
BDS. Once this workflow exists on the default branch, manual dispatch can
select an update branch; `workflow_call` also lets a protocol updater invoke
capture once its manual gate is resolved.

Each generated version pins its BDS download, archive checksum, build, and
gophertunnel module version in `generated/<version>/vanilla-source.json`. The
captured artifact should be reviewed and checked in beside that generated
protocol before the update is considered complete. The workflow requires
explicit EULA acceptance and never commits or pushes automatically.

## The manifest

Emitters consume only the canonical manifest, never the source docs. Before
`reconcile` writes one, it checks: source versions match across every pin;
field order, width, signedness, and length prefixes agree between sources;
validation bounds stay attached to the exact field they constrain; optionals
and union selectors are unambiguous; and every correction or adjudication
still matches the source content it was written against.

## Generated code

- **Go** — one file per packet, shared semantic types, closed union
  interfaces, symmetric `Marshal(IO)` methods, and a `Packet` interface with
  ID methods plus direction-aware constructor pools (`NewPacket`,
  `NewClientPacket`, `NewServerPacket`). Native type mappings (`uuid.UUID`,
  `mgl32.Vec2`/`Vec3`, `color.RGBA`, a value-based `Optional[T]`) are on by
  default; disable with `-native-types=false`.
- **Rust** — one module per packet, native enums, fallible slice-based
  `Encode`/`Decode`, direction-checked `Packet::decode_from`, and refcounted
  `Bytes` for buffer/NBT fields. Collection limits are a `Reader` setting
  (`set_collection_limit`), not a hard cap.

Both backends add target-language ergonomics on top of the manifest; neither
infers wire shape from anything but it.

## Cross-checking against independent implementations

- `parity` compares the manifest against an independently generated Axolotl
  layout.
- `verify-gophertunnel` parses a pinned gophertunnel commit
  (`tools/gophertunnel-oracle/lock.json`, a full SHA — checkout is rejected if
  `HEAD` differs) with `go/ast` and reports each packet as `AGREEMENT`,
  `DIVERGENCE`, `UNRESOLVED`, or `NO_ORACLE_PACKET`. Only an unaccepted
  `DIVERGENCE` fails the command. Reviewed exceptions live in
  `tools/gophertunnel-oracle/accepted-divergences.json`, each with a reason,
  evidence locator, and what would settle it.

Neither check edits the manifest. A real disagreement is resolved by a
correction under `generated/<version>/corrections/`, not by trusting the
oracle.

## Scope

Ingestion, reconciliation, validation, and Go/Rust emission are implemented.
The Rust backend doesn't yet emit borrowed string views. `migration/` holds
the superseded v1 (Axolotl-based) work and isn't an input to the current
pipeline.

See [docs/protocolgen-v2.md](docs/protocolgen-v2.md) for the manifest model,
source policy, adjudication format, and known gaps.

> Mojang protocol docs are EULA-restricted. Supply them as a local input;
> don't commit them. All checked-in fixtures are synthetic.
