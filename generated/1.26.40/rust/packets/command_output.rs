// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandOutput {
    pub origin_data: CommandOriginData,
    pub output: CommandOutputData,
}

pub const COMMANDOUTPUT_ORIGIN_DATA_SHAPE: &str = r##"{"kind":"struct","semantic":"CommandOriginData","type_id":"CommandOriginData.json#","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"UUID","semantic":"UUID","type_id":"mce__UUID.json#","encode":{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"RequestId","semantic":"RequestId","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"PlayerId","semantic":"PlayerId","encode":{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}"##;
pub const COMMANDOUTPUT_OUTPUT_SHAPE: &str = r#"{"kind":"struct","semantic":"CommandOutput","type_id":"CommandOutput","fields":[{"ordinal":0,"name":"Output Type","semantic":"Output Type","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Success Count","semantic":"Success Count","encode":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Output Messages","semantic":"Output Messages","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"CommandOutputMessage","type_id":"CommandOutputMessage","fields":[{"ordinal":0,"name":"Message ID","semantic":"Message ID","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Successful?","semantic":"Successful?","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Parameters","semantic":"Parameters","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Data Set","semantic":"Data Set","encode":{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl CommandOutput {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CommandOutputPacket.Origin Data", COMMANDOUTPUT_ORIGIN_DATA_SHAPE);
        encoder.field("CommandOutputPacket.Output", COMMANDOUTPUT_OUTPUT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CommandOutputPacket.Origin Data", COMMANDOUTPUT_ORIGIN_DATA_SHAPE);
        decoder.field("CommandOutputPacket.Output", COMMANDOUTPUT_OUTPUT_SHAPE);
    }
}
