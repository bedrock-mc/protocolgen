// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RemoveVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub dimension_type: DimensionType,
}

pub const REMOVEVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"EntityNetId","type_id":"EntityNetId","fields":[{"ordinal":0,"name":"Raw Id","semantic":"Raw Id","encode":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const REMOVEVOLUMEENTITY_DIMENSION_TYPE_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl RemoveVolumeEntity {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RemoveVolumeEntityPacket.Entity Network Id", REMOVEVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE);
        encoder.field("RemoveVolumeEntityPacket.Dimension Type", REMOVEVOLUMEENTITY_DIMENSION_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RemoveVolumeEntityPacket.Entity Network Id", REMOVEVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE);
        decoder.field("RemoveVolumeEntityPacket.Dimension Type", REMOVEVOLUMEENTITY_DIMENSION_TYPE_SHAPE);
    }
}
