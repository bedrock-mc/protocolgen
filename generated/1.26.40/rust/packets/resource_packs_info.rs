// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePacksInfo {
    pub resource_pack_required: bool,
    pub has_addon_packs: bool,
    pub has_scripts: bool,
    pub force_disable_vibrant_visuals: bool,
    pub world_template_id_and_version: PackIdVersion,
    pub resource_packs: Vec<PackInfoData>,
}

pub const RESOURCEPACKSINFO_RESOURCE_PACK_REQUIRED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKSINFO_HAS_ADDON_PACKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKSINFO_HAS_SCRIPTS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKSINFO_FORCE_DISABLE_VIBRANT_VISUALS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKSINFO_WORLD_TEMPLATE_ID_AND_VERSION_SHAPE: &str = r##"{"kind":"struct","semantic":"PackIdVersion","type_id":"PackIdVersion.json#","fields":[{"ordinal":0,"name":"Pack UUID","semantic":"Pack UUID","type_id":"mce__UUID.json#","encode":{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["mojang"]}},{"ordinal":1,"name":"Pack Version","semantic":"Pack Version","type_id":"SemVersion.json#","encode":{"kind":"struct","semantic":"SemVersion","type_id":"SemVersion.json#","fields":[{"ordinal":0,"name":"Version","semantic":"Version","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["mojang"]}}]},"symmetry":"symmetric","provenance":{"pins":["mojang"]}}]}"##;
pub const RESOURCEPACKSINFO_RESOURCE_PACKS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"PackInfoData","type_id":"PackInfoData","fields":[{"ordinal":0,"name":"Pack Id Version","semantic":"Pack Id Version","encode":{"kind":"struct","semantic":"PackIdVersion","type_id":"PackIdVersion","fields":[{"ordinal":0,"name":"Pack UUID","semantic":"Pack UUID","encode":{"kind":"primitive","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Pack Version","semantic":"Pack Version","encode":{"kind":"struct","semantic":"SemVersion","type_id":"SemVersion","fields":[{"ordinal":0,"name":"Version","semantic":"Version","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Pack Size","semantic":"Pack Size","encode":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Content Key","semantic":"Content Key","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Subpack Name","semantic":"Subpack Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Content Identity","semantic":"Content Identity","encode":{"kind":"struct","semantic":"ContentIdentity","type_id":"ContentIdentity","fields":[{"ordinal":0,"name":"Identity","semantic":"Identity","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":5,"name":"Has Scripts","semantic":"Has Scripts","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":6,"name":"Is Addon Pack","semantic":"Is Addon Pack","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":7,"name":"Is Ray Tracing Capable","semantic":"Is Ray Tracing Capable","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":8,"name":"CDN URL","semantic":"CDN URL","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl ResourcePacksInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePacksInfoPacket.Resource Pack Required", RESOURCEPACKSINFO_RESOURCE_PACK_REQUIRED_SHAPE);
        encoder.field("ResourcePacksInfoPacket.Has Addon Packs", RESOURCEPACKSINFO_HAS_ADDON_PACKS_SHAPE);
        encoder.field("ResourcePacksInfoPacket.Has Scripts", RESOURCEPACKSINFO_HAS_SCRIPTS_SHAPE);
        encoder.field("ResourcePacksInfoPacket.Force Disable Vibrant Visuals", RESOURCEPACKSINFO_FORCE_DISABLE_VIBRANT_VISUALS_SHAPE);
        encoder.field("ResourcePacksInfoPacket.World Template Id And Version", RESOURCEPACKSINFO_WORLD_TEMPLATE_ID_AND_VERSION_SHAPE);
        encoder.field("ResourcePacksInfoPacket.Resource Packs", RESOURCEPACKSINFO_RESOURCE_PACKS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePacksInfoPacket.Resource Pack Required", RESOURCEPACKSINFO_RESOURCE_PACK_REQUIRED_SHAPE);
        decoder.field("ResourcePacksInfoPacket.Has Addon Packs", RESOURCEPACKSINFO_HAS_ADDON_PACKS_SHAPE);
        decoder.field("ResourcePacksInfoPacket.Has Scripts", RESOURCEPACKSINFO_HAS_SCRIPTS_SHAPE);
        decoder.field("ResourcePacksInfoPacket.Force Disable Vibrant Visuals", RESOURCEPACKSINFO_FORCE_DISABLE_VIBRANT_VISUALS_SHAPE);
        decoder.field("ResourcePacksInfoPacket.World Template Id And Version", RESOURCEPACKSINFO_WORLD_TEMPLATE_ID_AND_VERSION_SHAPE);
        decoder.field("ResourcePacksInfoPacket.Resource Packs", RESOURCEPACKSINFO_RESOURCE_PACKS_SHAPE);
    }
}
