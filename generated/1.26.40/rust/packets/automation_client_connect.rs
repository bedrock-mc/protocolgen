// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AutomationClientConnect {
    pub web_socket_data: WebSocketPacketData,
}

pub const AUTOMATIONCLIENTCONNECT_WEB_SOCKET_DATA_SHAPE: &str = r#"{"kind":"struct","semantic":"WebSocketPacketData","type_id":"WebSocketPacketData","fields":[{"ordinal":0,"name":"Websocket Server URI","semantic":"Websocket Server URI","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl AutomationClientConnect {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AutomationClientConnectPacket.Web Socket Data", AUTOMATIONCLIENTCONNECT_WEB_SOCKET_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AutomationClientConnectPacket.Web Socket Data", AUTOMATIONCLIENTCONNECT_WEB_SOCKET_DATA_SHAPE);
    }
}
