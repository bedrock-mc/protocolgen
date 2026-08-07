// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct TrimData {
    pub trim_pattern_list: Vec<TrimPattern>,
    pub trim_material_list: Vec<TrimMaterial>,
}

pub const TRIMDATA_TRIM_PATTERN_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"TrimPattern","type_id":"TrimPattern","fields":[{"ordinal":0,"name":"Item Name","semantic":"Item Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Pattern Id","semantic":"Pattern Id","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const TRIMDATA_TRIM_MATERIAL_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"TrimMaterial","type_id":"TrimMaterial","fields":[{"ordinal":0,"name":"Material Id","semantic":"Material Id","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Color","semantic":"Color","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Item Name","semantic":"Item Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl TrimData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("TrimDataPacket.TrimPattern List", TRIMDATA_TRIM_PATTERN_LIST_SHAPE);
        encoder.field("TrimDataPacket.TrimMaterial List", TRIMDATA_TRIM_MATERIAL_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("TrimDataPacket.TrimPattern List", TRIMDATA_TRIM_PATTERN_LIST_SHAPE);
        decoder.field("TrimDataPacket.TrimMaterial List", TRIMDATA_TRIM_MATERIAL_LIST_SHAPE);
    }
}
