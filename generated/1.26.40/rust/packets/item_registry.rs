// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ItemRegistry {
    pub item_data: Vec<ItemData>,
}

pub const ITEMREGISTRY_ITEM_DATA_SHAPE: &str = r##"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"ItemData","type_id":"ItemData","fields":[{"ordinal":0,"name":"Item Name","semantic":"Item Name","type_id":"hashed_string.json#","encode":{"kind":"string","semantic":"hashed_string","type_id":"hashed_string.json#","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":1,"name":"Item Id","semantic":"Item Id","encode":{"kind":"primitive","primitive":{"code":"i16le","width":16,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"Is Component Based","semantic":"Is Component Based","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"Item Version","semantic":"Item Version","encode":{"kind":"enum","semantic":"ItemVersion","type_id":"enums/ItemVersion","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Legacy","encode":{"kind":"void"}},{"value":1,"name":"DataDriven","encode":{"kind":"void"}},{"value":2,"name":"None","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Item Component Data","semantic":"Item Component Data","encode":{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"##;

impl ItemRegistry {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ItemRegistryPacket.Item Data", ITEMREGISTRY_ITEM_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ItemRegistryPacket.Item Data", ITEMREGISTRY_ITEM_DATA_SHAPE);
    }
}
