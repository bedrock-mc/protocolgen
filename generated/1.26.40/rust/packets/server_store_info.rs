// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerStoreInfo {
    pub client_store_entry_point_configuration: Option<ServerConfigurationClientStoreEntryPointConfiguration>,
}

pub const SERVERSTOREINFO_CLIENT_STORE_ENTRY_POINT_CONFIGURATION_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"ServerConfiguration::ClientStoreEntryPointConfiguration","type_id":"ServerConfiguration::ClientStoreEntryPointConfiguration","fields":[{"ordinal":0,"name":"storeId","semantic":"storeId","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"storeName","semantic":"storeName","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl ServerStoreInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerStoreInfoPacket.client_store_entry_point_configuration", SERVERSTOREINFO_CLIENT_STORE_ENTRY_POINT_CONFIGURATION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerStoreInfoPacket.client_store_entry_point_configuration", SERVERSTOREINFO_CLIENT_STORE_ENTRY_POINT_CONFIGURATION_SHAPE);
    }
}
