// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerPresenceInfo {
    pub presence_configuration: Option<ServerConfigurationPresenceConfiguration>,
}

pub const SERVERPRESENCEINFO_PRESENCE_CONFIGURATION_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"ServerConfiguration::PresenceConfiguration","type_id":"ServerConfiguration::PresenceConfiguration","fields":[{"ordinal":0,"name":"richPresenceId","semantic":"richPresenceId","encode":{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl ServerPresenceInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerPresenceInfoPacket.presence_configuration", SERVERPRESENCEINFO_PRESENCE_CONFIGURATION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerPresenceInfoPacket.presence_configuration", SERVERPRESENCEINFO_PRESENCE_CONFIGURATION_SHAPE);
    }
}
