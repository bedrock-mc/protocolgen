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
    pub data: Nbt,
}

impl UpdateTrade {
    pub const ID: u32 = 80;
}
