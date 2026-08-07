// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientMovementPredictionSync {
    pub actor_data_flag: ActorDataFlagComponent,
    pub actor_bounding_box: ActorDataBoundingBoxComponent,
    pub movement_attributes: [f32; 9],
    pub actor_unique_id: ActorUniqueID,
    pub actor_flying_state: bool,
}

impl ClientMovementPredictionSync {
    pub const ID: u32 = 322;
}
