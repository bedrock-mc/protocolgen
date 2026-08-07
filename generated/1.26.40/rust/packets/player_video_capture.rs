// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerVideoCapture {
    pub action: PlayerVideoCaptureAction,
}

pub const PLAYERVIDEOCAPTURE_ACTION_SHAPE: &str = r#"{"kind":"union","variants":[{"value":0,"name":"PlayerVideoCapturePacketPayload::StopVideoCapture","encode":{"kind":"struct","semantic":"PlayerVideoCapturePacketPayload::StopVideoCapture","type_id":"PlayerVideoCapturePacketPayload::StopVideoCapture"}},{"value":1,"name":"PlayerVideoCapturePacketPayload::StartVideoCapture","encode":{"kind":"struct","semantic":"PlayerVideoCapturePacketPayload::StartVideoCapture","type_id":"PlayerVideoCapturePacketPayload::StartVideoCapture","fields":[{"ordinal":0,"name":"FrameRate","semantic":"FrameRate","encode":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"FilePrefix","semantic":"FilePrefix","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}],"control":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl PlayerVideoCapture {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerVideoCapturePacket.Action", PLAYERVIDEOCAPTURE_ACTION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerVideoCapturePacket.Action", PLAYERVIDEOCAPTURE_ACTION_SHAPE);
    }
}
