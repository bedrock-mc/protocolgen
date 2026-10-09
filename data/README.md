# Shared game-data generation

This authored module owns the source lock, input verification and generators
for the active [catalog](../generated/data/README.md). Its module path is
`github.com/bedrock-mc/protocolgen/data`. Runtime consumers import
`github.com/bedrock-mc/protocolgen/generated/data` instead. Both modules use
only the standard library; neither imports the other.

## Layout and ownership

- `source/` holds the release lock, verified input loader and small reviewed
  registry/correction inputs. External packs remain outside the repository.
- `internal/generator` and `cmd/generate` produce semantic values.
- `internal/runtimegen` and `cmd/runtimegen` produce runtime item values and
  release metadata.
- `../generated/data/` is the single active output catalog. It has its own
  `go.mod`; it is not a versioned packet-output directory.

The catalog's small API and lookup files are authored beside their data.
Files ending in `*_generated.go` are owned by their named generator and must
not be edited by hand. Both generators default to `../generated/data` when
run from this directory. `-out` selects another catalog root.

## Regeneration

Regenerate runtime items and release metadata from this directory:

```sh
go run ./cmd/runtimegen
```

The runtime catalog targets Minecraft 1.26.50, protocol 2193, with 2,076 items.
Sources include the locked Cloudburst registry, Allay item properties and tags, and
reviewed corrections. Generation rejects missing values and stale corrections.
It writes `item/runtime_generated.go`, `release_generated.go` and `release.json`
in one rollback-protected batch. The source-lock identity is SHA-256 over
`json.Marshal(source.Lock)`, so indentation in the input lock does not affect it.

Regenerate semantic values using local input trees:

```sh
go run ./cmd/generate \
  -cloudburst /path/to/CloudburstMC-Data \
  -bds /path/to/bedrock-server-1.26.32.2
```

Versions and digests come from [source/lock.json](source/lock.json). The
semantic generator verifies all three Cloudburst semantic files and every
consumed BDS JSON file before writing. Added, missing or changed input files fail validation.
The independent BDS 1.26.32.2 semantic provenance is preserved; it is not a
claim that all inputs were captured from the target runtime release.
Cloudburst `blocks.json` supplies all block fields, including collision, outline
and tint; there is no separate Allay block overlay. Its pinned revision is
`659ce1e2eee3a67045693f4fd5515c6ccf571953`, shared by the biome, named
voxel-shape and runtime item registry inputs. The catalog contains 22,091 block states, 89 biomes and
220 named voxel shapes. Named shapes are decoded from occupied grid cells into
block-local boxes; no BDS shape-pack overlay is required.

Thirty liquid-clip boxes in this extract have reversed bounds. The reviewed
`liquid_clip_omissions` input identifies each exact name, hash and source box.
Generation marks those fields unavailable and records them in
`semantic_sources.json`; it never clamps them or declares them empty. New,
changed and obsolete exceptions fail generation. Required collision and outline
shapes remain strictly validated. Zero-volume outlines remain known-empty.

Both commands accept `-cache /path/to/cache` and `-lock /path/to/lock.json`.
Use a reviewed alternate lock to change a release; there are no independent
version flags. Any source-lock change requires both generators. Catalog tests
require `SemanticSourceLockSHA256` to match `SourceLockSHA256`, so a partial
update fails even when the game version or source revision stays the same.
Semantic generation preserves runtime release metadata. Each generator removes
stale Go files only when they carry its exact ownership header, preserving the other generator's
outputs and authored files.

From the repository root:

```sh
make test-data
make verify-data CLOUDBURST_DIR=/path/to/CloudburstMC-Data BDS_DIR=/path/to/bds
```

`test-data` checks both modules independently. Generation tests use synthetic
fixtures. A repository-level AST check reads the sibling catalog's `Runtime`
struct and verifies every field is emitted without adding a module dependency
or maintaining a second schema list. Run the generator tests from a full
repository checkout. CI downloads and authenticates the locked Cloudburst
block/biome/shape files and compares the full generated projection with the
committed catalog. It also regenerates runtime items and metadata. To run the Cloudburst check locally, set
`PROTOCOLGEN_CLOUDBURST_DIR` when running `make test-data`. Full semantic
regeneration additionally requires the local BDS inputs.

## Shared input access

`source.Open("")` loads the embedded lock. `Sources.Read` authenticates each
input with SHA-256. `Sources.ResourcePack` downloads and extracts the locked
sample resource pack into the user's cache. Texture, glyph and block-metadata
capture tools share those pins while retaining their own processing logic.
Raw packs, textures and executables are not published here.

See [NOTICE.md](NOTICE.md) for source attribution.

## Remaining sources

Cloudburst's item registry covers all 2,076 items, but its component extract does
not supply complete stack sizes or inventory capabilities. Allay item properties
and supplemental tags remain explicit locked inputs, with reviewed corrections.
Removing them requires a complete version-matched item extract containing stack
limits, effective offhand acceptance and equipment-slot eligibility. Missing
component data is not evidence for a default stack size or an unavailable slot.

BDS behavior-pack inputs still supply base entity components and food at their
independently recorded source version. Cloudburst entity IDs and synchronized
properties do not contain those behavior components. The existing
`vanilla-data/endstone` exporter is a separate route for collecting live registry
facts from a matching BDS build; a packet registry alone does not contain all
item capabilities. This update does not change those behavior-pack inputs.
