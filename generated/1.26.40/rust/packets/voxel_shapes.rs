// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct VoxelShapes {
    pub shapes: Vec<VoxelShapesSerializableVoxelShape>,
    pub name_map: Vec<(String, VoxelShapesRegistryHandle)>,
    pub custom_shape_count: u16,
}

pub const VOXELSHAPES_SHAPES_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"VoxelShapes::SerializableVoxelShape","type_id":"VoxelShapes::SerializableVoxelShape","fields":[{"ordinal":0,"name":"Cells","semantic":"Cells","encode":{"kind":"struct","semantic":"VoxelShapes::SerializableCells","type_id":"VoxelShapes::SerializableCells","fields":[{"ordinal":0,"name":"X Size","semantic":"X Size","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y Size","semantic":"Y Size","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z Size","semantic":"Z Size","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Storage","semantic":"Storage","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"X Coordinates","semantic":"X Coordinates","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Y Coordinates","semantic":"Y Coordinates","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Z Coordinates","semantic":"Z Coordinates","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const VOXELSHAPES_NAME_MAP_SHAPE: &str = r#"{"kind":"map","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"representation":"ordered_entries","value":{"kind":"struct","semantic":"VoxelShapes::RegistryHandle","type_id":"VoxelShapes::RegistryHandle","fields":[{"ordinal":0,"name":"Value","semantic":"Value","encode":{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"key":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;
pub const VOXELSHAPES_CUSTOM_SHAPE_COUNT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl VoxelShapes {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("VoxelShapesPacket.Shapes", VOXELSHAPES_SHAPES_SHAPE);
        encoder.field("VoxelShapesPacket.Name Map", VOXELSHAPES_NAME_MAP_SHAPE);
        encoder.field("VoxelShapesPacket.Custom Shape Count", VOXELSHAPES_CUSTOM_SHAPE_COUNT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("VoxelShapesPacket.Shapes", VOXELSHAPES_SHAPES_SHAPE);
        decoder.field("VoxelShapesPacket.Name Map", VOXELSHAPES_NAME_MAP_SHAPE);
        decoder.field("VoxelShapesPacket.Custom Shape Count", VOXELSHAPES_CUSTOM_SHAPE_COUNT_SHAPE);
    }
}
