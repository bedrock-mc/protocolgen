// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundControlSchemeSet {
    pub control_scheme: ControlSchemeScheme,
}

pub const CLIENTBOUNDCONTROLSCHEMESET_CONTROL_SCHEME_SHAPE: &str = r#"{"kind":"enum","semantic":"ControlScheme::Scheme","type_id":"enums/ControlScheme::Scheme","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"locked_player_relative_strafe","encode":{"kind":"void"}},{"value":1,"name":"camera_relative","encode":{"kind":"void"}},{"value":2,"name":"camera_relative_strafe","encode":{"kind":"void"}},{"value":3,"name":"player_relative","encode":{"kind":"void"}},{"value":4,"name":"player_relative_strafe","encode":{"kind":"void"}}]}"#;

impl ClientboundControlSchemeSet {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundControlSchemeSetPacket.Control Scheme", CLIENTBOUNDCONTROLSCHEMESET_CONTROL_SCHEME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundControlSchemeSetPacket.Control Scheme", CLIENTBOUNDCONTROLSCHEMESET_CONTROL_SCHEME_SHAPE);
    }
}
