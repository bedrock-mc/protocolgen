// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub components: Nbt,
    pub json_identifier: String,
    pub instance_name: String,
    pub min_bounds: BlockPos,
    pub max_bounds: BlockPos,
    pub dimension_type: DimensionType,
    pub engine_version: String,
}

pub const ADDVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"EntityNetId","type_id":"EntityNetId","fields":[{"ordinal":0,"name":"Raw Id","semantic":"Raw Id","encode":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDVOLUMEENTITY_COMPONENTS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const ADDVOLUMEENTITY_JSON_IDENTIFIER_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ADDVOLUMEENTITY_INSTANCE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ADDVOLUMEENTITY_MIN_BOUNDS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDVOLUMEENTITY_MAX_BOUNDS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDVOLUMEENTITY_DIMENSION_TYPE_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const ADDVOLUMEENTITY_ENGINE_VERSION_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl AddVolumeEntity {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AddVolumeEntityPacket.Entity Network Id", ADDVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE);
        encoder.field("AddVolumeEntityPacket.Components", ADDVOLUMEENTITY_COMPONENTS_SHAPE);
        encoder.field("AddVolumeEntityPacket.JSON Identifier", ADDVOLUMEENTITY_JSON_IDENTIFIER_SHAPE);
        encoder.field("AddVolumeEntityPacket.Instance Name", ADDVOLUMEENTITY_INSTANCE_NAME_SHAPE);
        encoder.field("AddVolumeEntityPacket.Min Bounds", ADDVOLUMEENTITY_MIN_BOUNDS_SHAPE);
        encoder.field("AddVolumeEntityPacket.Max Bounds", ADDVOLUMEENTITY_MAX_BOUNDS_SHAPE);
        encoder.field("AddVolumeEntityPacket.Dimension Type", ADDVOLUMEENTITY_DIMENSION_TYPE_SHAPE);
        encoder.field("AddVolumeEntityPacket.Engine Version", ADDVOLUMEENTITY_ENGINE_VERSION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AddVolumeEntityPacket.Entity Network Id", ADDVOLUMEENTITY_ENTITY_NETWORK_ID_SHAPE);
        decoder.field("AddVolumeEntityPacket.Components", ADDVOLUMEENTITY_COMPONENTS_SHAPE);
        decoder.field("AddVolumeEntityPacket.JSON Identifier", ADDVOLUMEENTITY_JSON_IDENTIFIER_SHAPE);
        decoder.field("AddVolumeEntityPacket.Instance Name", ADDVOLUMEENTITY_INSTANCE_NAME_SHAPE);
        decoder.field("AddVolumeEntityPacket.Min Bounds", ADDVOLUMEENTITY_MIN_BOUNDS_SHAPE);
        decoder.field("AddVolumeEntityPacket.Max Bounds", ADDVOLUMEENTITY_MAX_BOUNDS_SHAPE);
        decoder.field("AddVolumeEntityPacket.Dimension Type", ADDVOLUMEENTITY_DIMENSION_TYPE_SHAPE);
        decoder.field("AddVolumeEntityPacket.Engine Version", ADDVOLUMEENTITY_ENGINE_VERSION_SHAPE);
    }
}
