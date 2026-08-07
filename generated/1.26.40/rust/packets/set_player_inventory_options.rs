// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetPlayerInventoryOptions {
    pub inventory_options: InventoryOptions,
}

pub const SETPLAYERINVENTORYOPTIONS_INVENTORY_OPTIONS_SHAPE: &str = r#"{"kind":"struct","semantic":"InventoryOptions","type_id":"InventoryOptions","fields":[{"ordinal":0,"name":"Left Inventory Tab","semantic":"Left Inventory Tab","encode":{"kind":"enum","semantic":"InventoryLeftTabIndex","type_id":"enums/InventoryLeftTabIndex","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"RecipeConstruction","encode":{"kind":"void"}},{"value":2,"name":"RecipeEquipment","encode":{"kind":"void"}},{"value":3,"name":"RecipeItems","encode":{"kind":"void"}},{"value":4,"name":"RecipeNature","encode":{"kind":"void"}},{"value":5,"name":"RecipeSearch","encode":{"kind":"void"}},{"value":6,"name":"Survival","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Right Inventory Tab","semantic":"Right Inventory Tab","encode":{"kind":"enum","semantic":"InventoryRightTabIndex","type_id":"enums/InventoryRightTabIndex","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"FullScreen","encode":{"kind":"void"}},{"value":2,"name":"Crafting","encode":{"kind":"void"}},{"value":3,"name":"Armor","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Filtering","semantic":"Filtering","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"Layout Inv","semantic":"Layout Inv","encode":{"kind":"enum","semantic":"InventoryLayout","type_id":"enums/InventoryLayout","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"InventoryOnly","encode":{"kind":"void"}},{"value":2,"name":"Default","encode":{"kind":"void"}},{"value":3,"name":"RecipeBookOnly","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Layout Craft","semantic":"Layout Craft","encode":{"kind":"enum","semantic":"InventoryLayout","type_id":"enums/InventoryLayout","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"InventoryOnly","encode":{"kind":"void"}},{"value":2,"name":"Default","encode":{"kind":"void"}},{"value":3,"name":"RecipeBookOnly","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl SetPlayerInventoryOptions {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetPlayerInventoryOptionsPacket.Inventory Options", SETPLAYERINVENTORYOPTIONS_INVENTORY_OPTIONS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetPlayerInventoryOptionsPacket.Inventory Options", SETPLAYERINVENTORYOPTIONS_INVENTORY_OPTIONS_SHAPE);
    }
}
