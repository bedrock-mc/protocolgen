// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateTrade {
    pub container_id: u8,
    pub r#type: u8,
    pub size: i32,
    pub trader_tier: i32,
    pub entity_unique_id: ActorUniqueID,
    pub last_trading_player: ActorUniqueID,
    pub display_name: String,
    pub use_new_trade_screen: bool,
    pub using_economy_trade: bool,
    pub data: Vec<u8>,
}

pub const UPDATETRADE_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATETRADE_R_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATETRADE_SIZE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const UPDATETRADE_TRADER_TIER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const UPDATETRADE_ENTITY_UNIQUE_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATETRADE_LAST_TRADING_PLAYER_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATETRADE_DISPLAY_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const UPDATETRADE_USE_NEW_TRADE_SCREEN_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATETRADE_USING_ECONOMY_TRADE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATETRADE_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl UpdateTrade {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateTradePacket.Container Id", UPDATETRADE_CONTAINER_ID_SHAPE);
        encoder.field("UpdateTradePacket.Type", UPDATETRADE_R_TYPE_SHAPE);
        encoder.field("UpdateTradePacket.Size", UPDATETRADE_SIZE_SHAPE);
        encoder.field("UpdateTradePacket.Trader Tier", UPDATETRADE_TRADER_TIER_SHAPE);
        encoder.field("UpdateTradePacket.Entity Unique Id", UPDATETRADE_ENTITY_UNIQUE_ID_SHAPE);
        encoder.field("UpdateTradePacket.Last Trading Player", UPDATETRADE_LAST_TRADING_PLAYER_SHAPE);
        encoder.field("UpdateTradePacket.Display Name", UPDATETRADE_DISPLAY_NAME_SHAPE);
        encoder.field("UpdateTradePacket.Use New Trade Screen", UPDATETRADE_USE_NEW_TRADE_SCREEN_SHAPE);
        encoder.field("UpdateTradePacket.Using Economy Trade", UPDATETRADE_USING_ECONOMY_TRADE_SHAPE);
        encoder.field("UpdateTradePacket.Data", UPDATETRADE_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateTradePacket.Container Id", UPDATETRADE_CONTAINER_ID_SHAPE);
        decoder.field("UpdateTradePacket.Type", UPDATETRADE_R_TYPE_SHAPE);
        decoder.field("UpdateTradePacket.Size", UPDATETRADE_SIZE_SHAPE);
        decoder.field("UpdateTradePacket.Trader Tier", UPDATETRADE_TRADER_TIER_SHAPE);
        decoder.field("UpdateTradePacket.Entity Unique Id", UPDATETRADE_ENTITY_UNIQUE_ID_SHAPE);
        decoder.field("UpdateTradePacket.Last Trading Player", UPDATETRADE_LAST_TRADING_PLAYER_SHAPE);
        decoder.field("UpdateTradePacket.Display Name", UPDATETRADE_DISPLAY_NAME_SHAPE);
        decoder.field("UpdateTradePacket.Use New Trade Screen", UPDATETRADE_USE_NEW_TRADE_SCREEN_SHAPE);
        decoder.field("UpdateTradePacket.Using Economy Trade", UPDATETRADE_USING_ECONOMY_TRADE_SHAPE);
        decoder.field("UpdateTradePacket.Data", UPDATETRADE_DATA_SHAPE);
    }
}
