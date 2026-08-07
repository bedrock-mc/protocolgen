// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundMapItemData {
    pub map_id: ActorUniqueID,
    pub dimension: u8,
    pub is_locked: bool,
    pub map_origin: BlockPos,
    pub creation_map_i_ds: Option<Vec<ActorUniqueID>>,
    pub scale: Option<i8>,
    pub tracked_actor_i_ds: Option<Vec<MapItemTrackedActorUniqueId>>,
    pub decorations: Option<Vec<MapDecoration>>,
    pub width: Option<i32>,
    pub height: Option<i32>,
    pub start_x: Option<i32>,
    pub start_y: Option<i32>,
    pub pixels: Option<Vec<u32>>,
}

impl ClientboundMapItemData {
    pub const ID: u32 = 67;
}
