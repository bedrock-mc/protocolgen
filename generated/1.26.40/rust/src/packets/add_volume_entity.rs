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

impl AddVolumeEntity {
    pub const ID: u32 = 166;
}
