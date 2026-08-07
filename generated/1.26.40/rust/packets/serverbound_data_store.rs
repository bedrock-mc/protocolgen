// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundDataStore {
    pub update: BedrockDDUIDataStoreUpdate,
}

pub const SERVERBOUNDDATASTORE_UPDATE_SHAPE: &str = r#"{"kind":"struct","semantic":"Bedrock::DDUI::DataStoreUpdate","type_id":"Bedrock::DDUI::DataStoreUpdate","fields":[{"ordinal":0,"name":"Data Store Name","semantic":"Data Store Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Property","semantic":"Property","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Path","semantic":"Path","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Data","semantic":"Data","encode":{"kind":"union","variants":[{"value":0,"name":"double","encode":{"kind":"primitive","primitive":{"code":"f64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}},{"value":1,"name":"bool","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}},{"value":2,"name":"string","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}],"control":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Property Update Count","semantic":"Property Update Count","encode":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":5,"name":"Path Update Count","semantic":"Path Update Count","encode":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl ServerboundDataStore {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerboundDataStorePacket.Update", SERVERBOUNDDATASTORE_UPDATE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerboundDataStorePacket.Update", SERVERBOUNDDATASTORE_UPDATE_SHAPE);
    }
}
