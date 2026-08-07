// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CurrentStructureFeature {
    pub current_structure_feature: String,
}

pub const CURRENTSTRUCTUREFEATURE_CURRENT_STRUCTURE_FEATURE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl CurrentStructureFeature {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CurrentStructureFeaturePacket.Current Structure Feature", CURRENTSTRUCTUREFEATURE_CURRENT_STRUCTURE_FEATURE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CurrentStructureFeaturePacket.Current Structure Feature", CURRENTSTRUCTUREFEATURE_CURRENT_STRUCTURE_FEATURE_SHAPE);
    }
}
