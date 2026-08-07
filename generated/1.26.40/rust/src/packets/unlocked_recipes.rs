// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UnlockedRecipes {
    pub packet_type: UnlockedRecipesPacketType,
    pub unlocked_recipes_list: Vec<String>,
}

impl UnlockedRecipes {
    pub const ID: u32 = 199;
}
