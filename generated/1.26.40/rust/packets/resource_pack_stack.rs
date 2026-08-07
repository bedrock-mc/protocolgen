// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackStack {
    pub texture_pack_required: bool,
    pub texture_pack_list: Vec<PackInstanceId>,
    pub base_game_version: String,
    pub experiments: Experiments,
    pub include_editor_packs: bool,
}

pub const RESOURCEPACKSTACK_TEXTURE_PACK_REQUIRED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKSTACK_TEXTURE_PACK_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"PackInstanceId","type_id":"PackInstanceId","fields":[{"ordinal":0,"name":"Pack ID","semantic":"Pack ID","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Version","semantic":"Version","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Sub Pack Name","semantic":"Sub Pack Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const RESOURCEPACKSTACK_BASE_GAME_VERSION_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const RESOURCEPACKSTACK_EXPERIMENTS_SHAPE: &str = r#"{"kind":"struct","semantic":"Experiments","type_id":"Experiments","fields":[{"ordinal":0,"name":"Toggles","semantic":"Toggles","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"element":{"kind":"struct","semantic":"cerealizer_ExperimentsAnon::ExperimentToggle","type_id":"cerealizer_ExperimentsAnon::ExperimentToggle","fields":[{"ordinal":0,"name":"Name","semantic":"Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Enabled","semantic":"Enabled","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"ExperimentsEverToggled","semantic":"ExperimentsEverToggled","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const RESOURCEPACKSTACK_INCLUDE_EDITOR_PACKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ResourcePackStack {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePackStackPacket.Texture Pack Required", RESOURCEPACKSTACK_TEXTURE_PACK_REQUIRED_SHAPE);
        encoder.field("ResourcePackStackPacket.Texture Pack List", RESOURCEPACKSTACK_TEXTURE_PACK_LIST_SHAPE);
        encoder.field("ResourcePackStackPacket.Base Game Version", RESOURCEPACKSTACK_BASE_GAME_VERSION_SHAPE);
        encoder.field("ResourcePackStackPacket.Experiments", RESOURCEPACKSTACK_EXPERIMENTS_SHAPE);
        encoder.field("ResourcePackStackPacket.Include Editor Packs", RESOURCEPACKSTACK_INCLUDE_EDITOR_PACKS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePackStackPacket.Texture Pack Required", RESOURCEPACKSTACK_TEXTURE_PACK_REQUIRED_SHAPE);
        decoder.field("ResourcePackStackPacket.Texture Pack List", RESOURCEPACKSTACK_TEXTURE_PACK_LIST_SHAPE);
        decoder.field("ResourcePackStackPacket.Base Game Version", RESOURCEPACKSTACK_BASE_GAME_VERSION_SHAPE);
        decoder.field("ResourcePackStackPacket.Experiments", RESOURCEPACKSTACK_EXPERIMENTS_SHAPE);
        decoder.field("ResourcePackStackPacket.Include Editor Packs", RESOURCEPACKSTACK_INCLUDE_EDITOR_PACKS_SHAPE);
    }
}
