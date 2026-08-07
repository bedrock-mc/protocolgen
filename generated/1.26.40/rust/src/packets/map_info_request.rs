// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MapInfoRequest {
    pub map_unique_id: ActorUniqueID,
    pub client_pixels_list: Vec<MapInfoRequestPacketAnonClientPixelsProxy>,
}

impl MapInfoRequest {
    pub const ID: u32 = 68;
}
