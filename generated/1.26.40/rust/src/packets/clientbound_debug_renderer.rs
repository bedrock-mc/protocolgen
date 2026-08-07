// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDebugRenderer {
    pub r#type: String,
    pub debug_marker_data: Option<ClientboundDebugRendererDebugMarkerData>,
}

impl ClientboundDebugRenderer {
    pub const ID: u32 = 164;
}
