// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LegacyTelemetryEvent {
    pub target_actor_id: ActorUniqueID,
    pub event_type: LegacyTelemetryEventType,
    pub use_player_id: bool,
    pub event_data: LegacyTelemetryEventEventData,
}

impl LegacyTelemetryEvent {
    pub const ID: u32 = 65;
}
