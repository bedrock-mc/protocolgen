# Data provenance

The semantic generator and initial data snapshot were moved from
[bedrock-mc/bedrock-data](https://github.com/bedrock-mc/bedrock-data) at commit
`a852f424`.

The semantic values are derived from:

- [CloudburstMC/Data](https://github.com/CloudburstMC/Data)
- Block collision, outline and tint values from
  [AllayMC/Allay](https://github.com/AllayMC/Allay)
- Base entity components, food definitions and named voxel shapes distributed
  with Minecraft Bedrock Dedicated Server
- Runtime registry values from [Dragonfly](https://github.com/HashimTheArab/dragonfly),
  item properties and tags from Allay, and the documented corrections in the
  source inputs
- Resource inputs from [Mojang/bedrock-samples](https://github.com/Mojang/bedrock-samples)

The source revisions and versions are recorded in
[`semantic_sources.json`](semantic_sources.json), with input digests in
[`source/lock.json`](source/lock.json). Upstream data and Minecraft
materials remain subject to their respective terms. No full BDS source pack,
raw generated source catalog or executable is included in this module.
