// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayStatus {
    pub status: PlayStatusType,
}

pub const PLAYSTATUS_STATUS_SHAPE: &str = r#"{"kind":"enum","semantic":"PlayStatus","type_id":"enums/PlayStatus","primitive":{"code":"i32be","width":32,"signed":true,"zigzag":false,"endianness":"big"},"variants":[{"value":0,"name":"LoginSuccess","encode":{"kind":"void"}},{"value":1,"name":"LoginFailed_ClientOld","encode":{"kind":"void"}},{"value":2,"name":"LoginFailed_ServerOld","encode":{"kind":"void"}},{"value":3,"name":"PlayerSpawn","encode":{"kind":"void"}},{"value":4,"name":"LoginFailed_InvalidTenant","encode":{"kind":"void"}},{"value":5,"name":"LoginFailed_EditionMismatchEduToVanilla","encode":{"kind":"void"}},{"value":6,"name":"LoginFailed_EditionMismatchVanillaToEdu","encode":{"kind":"void"}},{"value":7,"name":"LoginFailed_ServerFullSubClient","encode":{"kind":"void"}},{"value":8,"name":"LoginFailed_EditorMismatchEditorToVanilla","encode":{"kind":"void"}},{"value":9,"name":"LoginFailed_EditorMismatchVanillaToEditor","encode":{"kind":"void"}}]}"#;

impl PlayStatus {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayStatusPacket.Status", PLAYSTATUS_STATUS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayStatusPacket.Status", PLAYSTATUS_STATUS_SHAPE);
    }
}
