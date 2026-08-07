// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct VoxelShapes {
    pub shapes: Vec<VoxelShapesSerializableVoxelShape>,
    pub name_map: Vec<(String, VoxelShapesRegistryHandle)>,
    pub custom_shape_count: u16,
}

impl VoxelShapes {
    pub const ID: u32 = 337;
}
