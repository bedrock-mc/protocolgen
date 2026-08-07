// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandRequest {
    pub command: String,
    pub origin: CommandOriginData,
    pub is_internal: bool,
    pub version: String,
}

pub const COMMANDREQUEST_COMMAND_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const COMMANDREQUEST_ORIGIN_SHAPE: &str = r##"{"kind":"struct","semantic":"CommandOriginData","type_id":"CommandOriginData.json#","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"UUID","semantic":"UUID","type_id":"mce__UUID.json#","encode":{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"RequestId","semantic":"RequestId","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"PlayerId","semantic":"PlayerId","encode":{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}"##;
pub const COMMANDREQUEST_IS_INTERNAL_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const COMMANDREQUEST_VERSION_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl CommandRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CommandRequestPacket.Command", COMMANDREQUEST_COMMAND_SHAPE);
        encoder.field("CommandRequestPacket.Origin", COMMANDREQUEST_ORIGIN_SHAPE);
        encoder.field("CommandRequestPacket.IsInternal", COMMANDREQUEST_IS_INTERNAL_SHAPE);
        encoder.field("CommandRequestPacket.Version", COMMANDREQUEST_VERSION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CommandRequestPacket.Command", COMMANDREQUEST_COMMAND_SHAPE);
        decoder.field("CommandRequestPacket.Origin", COMMANDREQUEST_ORIGIN_SHAPE);
        decoder.field("CommandRequestPacket.IsInternal", COMMANDREQUEST_IS_INTERNAL_SHAPE);
        decoder.field("CommandRequestPacket.Version", COMMANDREQUEST_VERSION_SHAPE);
    }
}
