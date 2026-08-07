// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Transfer {
    pub server_address: String,
    pub server_port: u16,
    pub reload_world: bool,
    pub gatherings_configuration: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
}

pub const TRANSFER_SERVER_ADDRESS_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const TRANSFER_SERVER_PORT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const TRANSFER_RELOAD_WORLD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const TRANSFER_GATHERINGS_CONFIGURATION_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"ServerConfiguration::GatheringsConfigurationJoinInfo","type_id":"ServerConfiguration::GatheringsConfigurationJoinInfo","fields":[{"ordinal":0,"name":"experienceId","semantic":"experienceId","encode":{"kind":"primitive","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"experienceName","semantic":"experienceName","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"worldId","semantic":"worldId","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"worldName","semantic":"worldName","encode":{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"creatorId","semantic":"creatorId","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":5,"name":"targetId","semantic":"targetId","encode":{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":6,"name":"scenarioId","semantic":"scenarioId","encode":{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":7,"name":"serverId","semantic":"serverId","encode":{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl Transfer {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("TransferPacket.Server Address", TRANSFER_SERVER_ADDRESS_SHAPE);
        encoder.field("TransferPacket.Server Port", TRANSFER_SERVER_PORT_SHAPE);
        encoder.field("TransferPacket.Reload World", TRANSFER_RELOAD_WORLD_SHAPE);
        encoder.field("TransferPacket.Gatherings Configuration", TRANSFER_GATHERINGS_CONFIGURATION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("TransferPacket.Server Address", TRANSFER_SERVER_ADDRESS_SHAPE);
        decoder.field("TransferPacket.Server Port", TRANSFER_SERVER_PORT_SHAPE);
        decoder.field("TransferPacket.Reload World", TRANSFER_RELOAD_WORLD_SHAPE);
        decoder.field("TransferPacket.Gatherings Configuration", TRANSFER_GATHERINGS_CONFIGURATION_SHAPE);
    }
}
