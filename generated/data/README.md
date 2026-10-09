# Active Bedrock game catalog

This module provides shared game facts for Bedrock servers and clients. Import
it as `github.com/bedrock-mc/protocolgen/generated/data`. It uses only the Go
standard library and does not import the generators or source loader.

There is one active catalog. Older releases remain available through Git and
Go module pins, rather than separate catalog folders for each version.
`MinecraftVersion`, `ProtocolVersion` and
`SourceLockSHA256` identify its target and complete input lock.
[release.json](release.json) records the same target and every pinned input
revision and digest. Consumers using the separate `data/source` module can
compare the catalog hash against SHA-256 of `json.Marshal(sources.Lock)` to
reject mismatched module revisions even when the game version is unchanged.

```go
import "github.com/bedrock-mc/protocolgen/generated/data/block"

state, ok := block.StateByHash(hash)
if ok {
    properties, _ := block.PropertiesAt(state.Properties)
    boxes, available := block.Shape(properties.CollisionShape)
    // available distinguishes a missing shape from a known empty shape.
    _, _ = boxes, available
}
```

| Package | Contents |
| --- | --- |
| `data` | Target release, independent source versions and generated counts |
| `block` | State hashes, physical properties, light, tint and shape references |
| `biome` | Climate, terrain values, water color, rain and tags |
| `entity` | Base collision, health, movement, scale, projectile and family values |
| `item` | Runtime registry, stack limits, equipment flags and food values |
| `voxelshape` | Named shapes in block-local coordinates |

Named values include `block.Stone`, `biome.Plains`, `entity.Arrow`, `item.Apple`
and `voxelshape.Anvil`. Each package also exposes lookups and iteration.
Treat returned slices and definitions as immutable.

## Independent source versions

The runtime target is Minecraft 1.26.50, protocol 2193, with 2,076 runtime items.
The semantic snapshot uses CloudburstMC/Data commit
`fb969c547236d87a17181941cd585a0eb18f7ceb`, BDS `1.26.32.2`, and Allay commit
`59d4007e322a0acf2b0add59133c81e1f7e4c501` for collision, outline and tint.
These inputs contain 16,913 block states, 88 biomes, 96 entities, 38 foods and
57 named voxel shapes. [semantic_sources.json](semantic_sources.json) and
`version_generated.go` record these independent versions and counts.
`BDSVersion` describes the behavior and named-shape inputs, not the runtime
protocol target.

Entity values describe base components. They do not execute goals, events,
filters or component-group transitions. Food definitions alone are not a
runtime item registry. Neighbor-dependent block shapes describe the extraction
context; live worlds may need their own models to resolve connections.
Unavailable shapes remain distinct from known empty shapes.

## Generation and API ownership

[Authored generators and source locks](../../data/README.md) live in `data/`.
Run both generators there when changing the release. `*_generated.go` files
are owned by the generator named in their header. The small type, lookup and
test files beside those values are authored API code and may be edited.
`release.json` belongs to runtime generation; `semantic_sources.json` belongs
to semantic generation. Each preserves the other's outputs. Any source-lock
change requires both generators: catalog tests compare their complete lock
hashes and reject a partial update.

These values are semantic projections, not complete game packs. Raw BDS packs,
textures, structures, localization files and executables are not included.
See [NOTICE.md](NOTICE.md) for attribution.
