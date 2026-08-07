// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct FeatureRegistry {
    pub features_data_list: Vec<FeatureRegistryFeatureBinaryJsonFormat>,
}

pub const FEATUREREGISTRY_FEATURES_DATA_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"FeatureRegistry::FeatureBinaryJsonFormat","type_id":"FeatureRegistry::FeatureBinaryJsonFormat","fields":[{"ordinal":0,"name":"Feature Name","semantic":"Feature Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Binary Json Output","semantic":"Binary Json Output","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl FeatureRegistry {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("FeatureRegistryPacket.FeaturesDataList", FEATUREREGISTRY_FEATURES_DATA_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("FeatureRegistryPacket.FeaturesDataList", FEATUREREGISTRY_FEATURES_DATA_LIST_SHAPE);
    }
}
