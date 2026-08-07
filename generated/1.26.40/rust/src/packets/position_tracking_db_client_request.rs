// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PositionTrackingDBClientRequest {
    pub action: PositionTrackingDBClientRequestAction,
    pub id: PositionTrackingId,
}

impl PositionTrackingDBClientRequest {
    pub const ID: u32 = 154;
}
