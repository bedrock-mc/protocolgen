// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CraftingData {
    pub shaped_recipes: Vec<ShapedRecipe>,
    pub shapeless_recipes: Vec<ShapelessRecipe>,
    pub multi_recipes: Vec<MultiRecipe>,
    pub user_data_shapeless_recipes: Vec<ShapelessRecipe>,
    pub shapeless_chemistry_recipes: Vec<ShapelessRecipe>,
    pub shaped_chemistry_recipes: Vec<ShapedRecipe>,
    pub smithing_transform_recipes: Vec<SmithingTransformRecipe>,
    pub smithing_trim_recipes: Vec<SmithingTrimRecipe>,
    pub potion_mixes: Vec<PotionMixDataEntry>,
    pub container_mixes: Vec<ContainerMixDataEntry>,
    pub material_reducers: Vec<MaterialReducerDataEntry>,
    pub clear_recipes: bool,
}

impl CraftingData {
    pub const ID: u32 = 52;
}
