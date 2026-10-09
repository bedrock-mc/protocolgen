# Shared Bedrock game data

This nested Go module provides the shared semantic game data used by Bedrock
servers and clients. Import it as `github.com/bedrock-mc/protocolgen/data`.
The data packages depend only on the Go standard library.

```go
import "github.com/bedrock-mc/protocolgen/data/block"

state, ok := block.StateByHash(hash)
if ok {
	properties, _ := block.PropertiesAt(state.Properties)
	boxes, available := block.Shape(properties.CollisionShape)
	// Use available to distinguish a missing shape from a known empty shape.
	_, _ = boxes, available
}
```

| Package | Contents |
| --- | --- |
| `data` | Source versions and generated record counts |
| `data/block` | State hashes, physical properties, light, tint and shape references |
| `data/biome` | Climate, terrain values, water color, rain and tags |
| `data/entity` | Base collision, health, movement, scale, projectile and family values |
| `data/item` | Runtime item registry, stack limits, equipment flags and food values |
| `data/voxelshape` | Named shapes in block-local coordinates |

Each named definition has an exported value, such as `block.Stone`,
`biome.Plains`, `entity.Arrow`, `item.Apple` or `voxelshape.Anvil`.
Lookup and iteration functions are available in each package. Treat returned
slices and exported definitions as immutable.

## Current semantic snapshot

The scalar and BDS semantic snapshot is preserved from `bedrock-mc/bedrock-data`
commit `a852f424`. Collision, outline and tint values retain the Allay snapshot
already used by runtime consumers:

- CloudburstMC/Data commit `fb969c547236d87a17181941cd585a0eb18f7ceb` supplies
  block and biome data.
- Bedrock Dedicated Server `1.26.32.2` supplies base entity components, foods
  and named voxel shapes.
- Allay commit `59d4007e322a0acf2b0add59133c81e1f7e4c501` supplies collision,
  outline and tint values, joined to every Cloudburst state by hash and name.
- The snapshot contains 16,913 block states, 88 biomes, 96 entities, 38 foods
  and 57 named voxel shapes.

[`semantic_sources.json`](semantic_sources.json) records those source pins and
counts. `version_generated.go` exposes them to Go callers. The BDS version
describes the behavior and shape inputs; it must not be interpreted as a
verified protocol version for every independently sourced table.

## Regeneration

The runtime catalog targets Minecraft 1.26.50, protocol 2193. It contains 2,076
item records with IDs, component versions, stack sizes and equipment flags.
Its sources are the locked registry input and Allay item properties and tags.
Explicit corrections cover gaps in those sources. Corrections need evidence,
and generation fails when a correction becomes stale or a required value is
missing. `item.LookupRuntime` and `item.RuntimeItems` expose this catalog.

Regenerate it from this directory with:

```sh
go run ./cmd/runtimegen
```

Run the semantic generator from this `data` directory. Input trees remain
outside the repository:

```sh
go run ./cmd/generate \
  -cloudburst /path/to/CloudburstMC-Data \
  -bds /path/to/bedrock-server-1.26.32.2 \
  -out .
go test ./...
```

Versions and file digests come from [`source/lock.json`](source/lock.json).
The generator verifies both Cloudburst files and every consumed BDS JSON file
before writing output. Changed, missing or added input files fail validation.
Use the Cloudburst revision named by the lock; newer checkouts can use different
file names and schemas. BDS input manifests contain only file names and hashes.
The locked Allay input is downloaded through the authenticated shared cache.
Use `-cache /path/to/cache` to choose its location and `-lock /path/to/lock.json`
to generate a reviewed alternate release. There are no separate version flags.
The generator writes deterministic Go values and its semantic source record.
It validates duplicate identifiers and state hashes, rejects malformed or
non-finite boxes, handles JSON comments, normalizes named shapes, and shares
repeated block properties and shapes through compact indexes.

Writes are staged and rolled back if installation fails. Cleanup removes
only stale files carrying this generator's exact ownership header. Handwritten
files and outputs from other generators in the same package are preserved.

Generation tests use small synthetic fixtures. Runtime tests check the
committed semantic counts, lookups and cross-table references.

## Resource inputs

`source.Open("")` loads the embedded lock. `Sources.Read` verifies each input
against its SHA-256 digest. `Sources.ResourcePack` downloads and extracts the
locked official sample resource pack into the user's cache. Texture and glyph
tools share that version and downloader while retaining their own rendering
logic. Raw packs and images stay outside this repository.

To update a release, review the source revisions and digests together in the
lock, refresh the registry and any necessary corrections, then run both
generators. `make verify-data CLOUDBURST_DIR=... BDS_DIR=...` from the repository
root verifies the complete generated snapshot. CI regenerates the runtime item
catalog and runs all data-module tests; the full semantic regeneration also
needs the external BDS inputs.

## Boundaries

These are semantic projections, not copies of complete game packs. The module
does not publish raw BDS behavior/resource catalogs, textures, structures,
localization files or executables, and its generator does not emit them.
See [NOTICE.md](NOTICE.md) for source attribution.

Entity values describe base components. They do not execute goals, events,
filters or component-group transitions. Food definitions alone are not a
complete runtime item registry. Neighbor-dependent block shapes describe the
source extraction context; a live world may need its own model to resolve
connections. Unavailable shape data must remain distinct from a known empty
shape rather than silently becoming a full cube.

Protocol capture and validation remain in the surrounding protocolgen
repository. This module supplies game facts to consumers without importing
the packet generator or a server implementation.
