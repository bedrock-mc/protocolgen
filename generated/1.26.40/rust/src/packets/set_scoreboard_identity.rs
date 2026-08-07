// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetScoreboardIdentity {
    pub scoreboard_identity_packet_type: ScoreboardIdentityPacketType,
    pub scoreboard_identity_info: Vec<ScoreboardIdentityPacketInfo>,
}

impl SetScoreboardIdentity {
    pub const ID: u32 = 112;
}
