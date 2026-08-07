// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackClientResponse {
    pub response: ResourcePackClientResponseResponse,
}

pub const RESOURCEPACKCLIENTRESPONSE_RESPONSE_SHAPE: &str = r#"{"kind":"union","variants":[{"value":1,"name":"ResourcePackClientResponsePacketPayload::Cancel","encode":{"kind":"struct","semantic":"ResourcePackClientResponsePacketPayload::Cancel","type_id":"ResourcePackClientResponsePacketPayload::Cancel","fields":[{"ordinal":0,"name":"Response Type","semantic":"Response Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":2,"name":"ResourcePackClientResponsePacketPayload::Downloading","encode":{"kind":"struct","semantic":"ResourcePackClientResponsePacketPayload::Downloading","type_id":"ResourcePackClientResponsePacketPayload::Downloading","fields":[{"ordinal":0,"name":"Response Type","semantic":"Response Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Downloading Packs","semantic":"Downloading Packs","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":3,"name":"ResourcePackClientResponsePacketPayload::DownloadingFinished","encode":{"kind":"struct","semantic":"ResourcePackClientResponsePacketPayload::DownloadingFinished","type_id":"ResourcePackClientResponsePacketPayload::DownloadingFinished","fields":[{"ordinal":0,"name":"Response Type","semantic":"Response Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":4,"name":"ResourcePackClientResponsePacketPayload::ResourcePackStackFinished","encode":{"kind":"struct","semantic":"ResourcePackClientResponsePacketPayload::ResourcePackStackFinished","type_id":"ResourcePackClientResponsePacketPayload::ResourcePackStackFinished","fields":[{"ordinal":0,"name":"Response Type","semantic":"Response Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}],"control":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}}}"#;

impl ResourcePackClientResponse {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePackClientResponsePacket.Response", RESOURCEPACKCLIENTRESPONSE_RESPONSE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePackClientResponsePacket.Response", RESOURCEPACKCLIENTRESPONSE_RESPONSE_SHAPE);
    }
}
