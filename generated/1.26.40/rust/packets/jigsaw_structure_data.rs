// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct JigsawStructureData {
    pub jigsaw_structure_data_tag: Nbt,
}

pub const JIGSAWSTRUCTUREDATA_JIGSAW_STRUCTURE_DATA_TAG_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl JigsawStructureData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("JigsawStructureDataPacket.Jigsaw Structure Data Tag", JIGSAWSTRUCTUREDATA_JIGSAW_STRUCTURE_DATA_TAG_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("JigsawStructureDataPacket.Jigsaw Structure Data Tag", JIGSAWSTRUCTUREDATA_JIGSAW_STRUCTURE_DATA_TAG_SHAPE);
    }
}
