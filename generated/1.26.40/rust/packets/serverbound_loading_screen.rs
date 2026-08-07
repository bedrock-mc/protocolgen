// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundLoadingScreen {
    pub loading_screen_packet_type: ServerboundLoadingScreenPacketType,
    pub loading_screen_id: Option<u32>,
}

pub const SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_PACKET_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"ServerboundLoadingScreenPacketType","type_id":"enums/ServerboundLoadingScreenPacketType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":1,"name":"StartLoadingScreen","encode":{"kind":"void"}},{"value":2,"name":"EndLoadingScreen","encode":{"kind":"void"}}]}"#;
pub const SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_ID_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;

impl ServerboundLoadingScreen {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerboundLoadingScreenPacket.Loading Screen Packet Type", SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_PACKET_TYPE_SHAPE);
        encoder.field("ServerboundLoadingScreenPacket.Loading Screen Id", SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerboundLoadingScreenPacket.Loading Screen Packet Type", SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_PACKET_TYPE_SHAPE);
        decoder.field("ServerboundLoadingScreenPacket.Loading Screen Id", SERVERBOUNDLOADINGSCREEN_LOADING_SCREEN_ID_SHAPE);
    }
}
