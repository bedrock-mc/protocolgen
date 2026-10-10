# Data provenance

The semantic generator and initial data snapshot were moved from
[bedrock-mc/bedrock-data](https://github.com/bedrock-mc/bedrock-data) at commit
`a852f424`.

The semantic values are derived from:

- Block properties, geometry, named voxel shapes, biome values, block palettes, data-driven block components,
  and complete item registry identities/components from
  [CloudburstMC/Data](https://github.com/CloudburstMC/Data)
- Numeric biome IDs for the retail 1.26.50 catalog from the public,
  BDS-verified [Cinnabar biome registry](https://github.com/bedrock-mc/cinnabar/blob/9fc5ac80a8b4a9f897dc992ad55ccc30a2afaf16/crates/assets/data/biome-registry-v2193.bin)
- Base entity components and food definitions distributed
  with Minecraft Bedrock Dedicated Server
- Item properties and supplemental tags from
  [AllayMC/Allay](https://github.com/AllayMC/Allay), plus the documented corrections
  in the source inputs
- Resource inputs from [Mojang/bedrock-samples](https://github.com/Mojang/bedrock-samples)
- Provisional block tool and harvest classifications from
  [PrismarineJS/minecraft-data](https://github.com/PrismarineJS/minecraft-data),
  pinned to its Bedrock 1.26.30 `blocks.json` extract. Its README declares MIT;
  these facts have not been verified for the catalog's target Bedrock release.

The source revisions and versions are recorded in
[`semantic_sources.json`](../generated/data/semantic_sources.json), with input digests in
[`source/lock.json`](source/lock.json). Upstream data and Minecraft
materials remain subject to their respective terms. No full BDS source pack,
raw generated source catalog or executable is included in this module.

The pinned Cloudburst extract contains 30 invalid liquid-clip boxes. Their exact
source values are recorded as omissions; the catalog reports those shapes as
unavailable. The public [Endstone data exporter](https://github.com/EndstoneMC/endstone/blob/3491c609ddfde392cee2b062e3063e39aae87274/src/endstone/core/devtools/vanilla_data.cpp)
illustrates why an output box without the query's success flag cannot establish
valid geometry. This is not a claim about the exact tool revision that produced
the Cloudburst files.
