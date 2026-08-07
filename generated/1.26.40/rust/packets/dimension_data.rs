// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct DimensionData {
    pub definitions: Vec<(String, DimensionDefinitionGroupDimensionDefinition)>,
}

pub const DIMENSIONDATA_DEFINITIONS_SHAPE: &str = r##"{"kind":"map","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"representation":"ordered_entries","value":{"kind":"struct","semantic":"DimensionDefinitionGroup::DimensionDefinition","type_id":"DimensionDefinitionGroup::DimensionDefinition","fields":[{"ordinal":0,"name":"Height Maximum","semantic":"Height Maximum","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":1,"name":"Height Minimum","semantic":"Height Minimum","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"Generator Type","semantic":"Generator Type","encode":{"kind":"enum","semantic":"GeneratorType","type_id":"enums/GeneratorType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Legacy","encode":{"kind":"void"}},{"value":1,"name":"Overworld","encode":{"kind":"void"}},{"value":2,"name":"Flat","encode":{"kind":"void"}},{"value":3,"name":"Nether","encode":{"kind":"void"}},{"value":4,"name":"TheEnd","encode":{"kind":"void"}},{"value":5,"name":"Void","encode":{"kind":"void"}},{"value":6,"name":"Undefined","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Dimension Type","semantic":"Dimension Type","type_id":"DimensionType.json#","encode":{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":4,"name":"Pack Id","semantic":"Pack Id","type_id":"mce__UUID.json#","encode":{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]},"key":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"##;

impl DimensionData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("DimensionDataPacket.Definitions", DIMENSIONDATA_DEFINITIONS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("DimensionDataPacket.Definitions", DIMENSIONDATA_DEFINITIONS_SHAPE);
    }
}
