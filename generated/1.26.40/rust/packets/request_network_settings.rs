// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestNetworkSettings {
    pub client_network_version: i32,
}

pub const REQUESTNETWORKSETTINGS_CLIENT_NETWORK_VERSION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32be","width":32,"signed":true,"zigzag":false,"endianness":"big"}}"#;

impl RequestNetworkSettings {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RequestNetworkSettingsPacket.ClientNetworkVersion", REQUESTNETWORKSETTINGS_CLIENT_NETWORK_VERSION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RequestNetworkSettingsPacket.ClientNetworkVersion", REQUESTNETWORKSETTINGS_CLIENT_NETWORK_VERSION_SHAPE);
    }
}
