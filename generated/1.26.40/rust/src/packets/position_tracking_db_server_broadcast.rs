// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PositionTrackingDBServerBroadcast {
    pub action: PositionTrackingDBServerBroadcastAction,
    pub id: PositionTrackingId,
    pub position_tracking_data: Nbt,
}

impl PositionTrackingDBServerBroadcast {
    pub const ID: u32 = 153;
}
