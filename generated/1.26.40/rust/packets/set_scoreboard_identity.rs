// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetScoreboardIdentity {
    pub scoreboard_identity_packet_type: ScoreboardIdentityPacketType,
    pub scoreboard_identity_info: Vec<ScoreboardIdentityPacketInfo>,
}

pub const SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_PACKET_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"ScoreboardIdentityPacketType","type_id":"enums/ScoreboardIdentityPacketType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Update","encode":{"kind":"void"}},{"value":1,"name":"Remove","encode":{"kind":"void"}}]}"#;
pub const SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_INFO_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"ScoreboardIdentityPacketInfo","type_id":"ScoreboardIdentityPacketInfo","fields":[{"ordinal":0,"name":"Scoreboard Id","semantic":"Scoreboard Id","encode":{"kind":"struct","semantic":"ScoreboardId","type_id":"ScoreboardId","fields":[{"ordinal":0,"name":"Scoreboard Id","semantic":"Scoreboard Id","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Player Unique Id","semantic":"Player Unique Id","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl SetScoreboardIdentity {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetScoreboardIdentityPacket.Scoreboard Identity Packet Type", SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_PACKET_TYPE_SHAPE);
        encoder.field("SetScoreboardIdentityPacket.Scoreboard Identity Info", SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_INFO_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetScoreboardIdentityPacket.Scoreboard Identity Packet Type", SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_PACKET_TYPE_SHAPE);
        decoder.field("SetScoreboardIdentityPacket.Scoreboard Identity Info", SETSCOREBOARDIDENTITY_SCOREBOARD_IDENTITY_INFO_SHAPE);
    }
}
