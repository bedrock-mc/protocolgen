// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UnlockedRecipes {
    pub packet_type: UnlockedRecipesPacketType,
    pub unlocked_recipes_list: Vec<String>,
}

pub const UNLOCKEDRECIPES_PACKET_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"UnlockedRecipesPacketPayload::PacketType","type_id":"enums/UnlockedRecipesPacketPayload::PacketType","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"},"variants":[{"value":0,"name":"Empty","encode":{"kind":"void"}},{"value":1,"name":"InitiallyUnlockedRecipes","encode":{"kind":"void"}},{"value":2,"name":"NewlyUnlockedRecipes","encode":{"kind":"void"}},{"value":3,"name":"RemoveUnlockedRecipes","encode":{"kind":"void"}},{"value":4,"name":"RemoveAllUnlockedRecipes","encode":{"kind":"void"}}]}"#;
pub const UNLOCKEDRECIPES_UNLOCKED_RECIPES_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;

impl UnlockedRecipes {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UnlockedRecipesPacket.Packet Type", UNLOCKEDRECIPES_PACKET_TYPE_SHAPE);
        encoder.field("UnlockedRecipesPacket.Unlocked Recipes List", UNLOCKEDRECIPES_UNLOCKED_RECIPES_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UnlockedRecipesPacket.Packet Type", UNLOCKEDRECIPES_PACKET_TYPE_SHAPE);
        decoder.field("UnlockedRecipesPacket.Unlocked Recipes List", UNLOCKEDRECIPES_UNLOCKED_RECIPES_LIST_SHAPE);
    }
}
