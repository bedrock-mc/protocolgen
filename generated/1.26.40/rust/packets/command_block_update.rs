// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CommandBlockUpdate {
    pub target: CommandBlockUpdateTarget,
    pub command: String,
    pub last_output: String,
    pub name: String,
    pub filtered_name: String,
    pub track_output: bool,
    pub tick_delay: i32,
    pub execute_on_first_tick: bool,
}

pub const COMMANDBLOCKUPDATE_TARGET_SHAPE: &str = r#"{"kind":"union","variants":[{"value":0,"name":"CommandBlockUpdatePacketPayload::EntityCommandTarget","encode":{"kind":"struct","semantic":"CommandBlockUpdatePacketPayload::EntityCommandTarget","type_id":"CommandBlockUpdatePacketPayload::EntityCommandTarget","fields":[{"ordinal":0,"name":"Target Runtime ID","semantic":"Target Runtime ID","encode":{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},{"value":1,"name":"CommandBlockUpdatePacketPayload::BlockCommandData","encode":{"kind":"struct","semantic":"CommandBlockUpdatePacketPayload::BlockCommandData","type_id":"CommandBlockUpdatePacketPayload::BlockCommandData","fields":[{"ordinal":0,"name":"Block Position","semantic":"Block Position","encode":{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Command Block Mode","semantic":"Command Block Mode","encode":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Redstone Mode","semantic":"Redstone Mode","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Is Conditional","semantic":"Is Conditional","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}],"control":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}}"#;
pub const COMMANDBLOCKUPDATE_COMMAND_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const COMMANDBLOCKUPDATE_LAST_OUTPUT_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const COMMANDBLOCKUPDATE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const COMMANDBLOCKUPDATE_FILTERED_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const COMMANDBLOCKUPDATE_TRACK_OUTPUT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const COMMANDBLOCKUPDATE_TICK_DELAY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const COMMANDBLOCKUPDATE_EXECUTE_ON_FIRST_TICK_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl CommandBlockUpdate {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CommandBlockUpdatePacket.Target", COMMANDBLOCKUPDATE_TARGET_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Command", COMMANDBLOCKUPDATE_COMMAND_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Last Output", COMMANDBLOCKUPDATE_LAST_OUTPUT_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Name", COMMANDBLOCKUPDATE_NAME_SHAPE);
        encoder.field("CommandBlockUpdatePacket.FilteredName", COMMANDBLOCKUPDATE_FILTERED_NAME_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Track Output", COMMANDBLOCKUPDATE_TRACK_OUTPUT_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Tick Delay", COMMANDBLOCKUPDATE_TICK_DELAY_SHAPE);
        encoder.field("CommandBlockUpdatePacket.Execute On First Tick", COMMANDBLOCKUPDATE_EXECUTE_ON_FIRST_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CommandBlockUpdatePacket.Target", COMMANDBLOCKUPDATE_TARGET_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Command", COMMANDBLOCKUPDATE_COMMAND_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Last Output", COMMANDBLOCKUPDATE_LAST_OUTPUT_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Name", COMMANDBLOCKUPDATE_NAME_SHAPE);
        decoder.field("CommandBlockUpdatePacket.FilteredName", COMMANDBLOCKUPDATE_FILTERED_NAME_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Track Output", COMMANDBLOCKUPDATE_TRACK_OUTPUT_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Tick Delay", COMMANDBLOCKUPDATE_TICK_DELAY_SHAPE);
        decoder.field("CommandBlockUpdatePacket.Execute On First Tick", COMMANDBLOCKUPDATE_EXECUTE_ON_FIRST_TICK_SHAPE);
    }
}
