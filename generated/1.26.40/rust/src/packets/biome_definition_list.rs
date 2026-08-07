// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BiomeDefinitionList {
    pub map_of_biome_names_to_data: Vec<(u16, BiomeDefinitionData)>,
    pub string_list: BiomeStringList,
}

impl BiomeDefinitionList {
    pub const ID: u32 = 122;
}
