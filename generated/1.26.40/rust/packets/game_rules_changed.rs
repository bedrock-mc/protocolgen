// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameRulesChanged {
    pub rule_data: GameRulesChangedPacketData,
}

pub const GAMERULESCHANGED_RULE_DATA_SHAPE: &str = r#"{"kind":"struct","semantic":"GameRulesChangedPacketData","type_id":"GameRulesChangedPacketData","fields":[{"ordinal":0,"name":"Rules List","semantic":"Rules List","encode":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"GameRule","type_id":"GameRule","fields":[{"ordinal":0,"name":"Rule Name","semantic":"Rule Name","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Rule Can Be Modified","semantic":"Rule Can Be Modified","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Rule Value","semantic":"Rule Value","encode":{"kind":"union","variants":[{"value":0,"name":"Empty0","encode":{"kind":"void"}},{"value":1,"name":"bool","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}},{"value":2,"name":"int32","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}},{"value":3,"name":"float","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}],"control":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl GameRulesChanged {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("GameRulesChangedPacket.Rule Data", GAMERULESCHANGED_RULE_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("GameRulesChangedPacket.Rule Data", GAMERULESCHANGED_RULE_DATA_SHAPE);
    }
}
