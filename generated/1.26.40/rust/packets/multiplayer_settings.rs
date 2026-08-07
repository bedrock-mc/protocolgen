// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MultiplayerSettings {
    pub packet_type: MultiplayerSettingsPacketType,
}

pub const MULTIPLAYERSETTINGS_PACKET_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"MultiplayerSettingsPacketType","type_id":"enums/MultiplayerSettingsPacketType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"EnableMultiplayer","encode":{"kind":"void"}},{"value":1,"name":"DisableMultiplayer","encode":{"kind":"void"}},{"value":2,"name":"RefreshJoincode","encode":{"kind":"void"}}]}"#;

impl MultiplayerSettings {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MultiplayerSettingsPacket.PacketType", MULTIPLAYERSETTINGS_PACKET_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MultiplayerSettingsPacket.PacketType", MULTIPLAYERSETTINGS_PACKET_TYPE_SHAPE);
    }
}
