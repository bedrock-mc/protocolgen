// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LecternUpdate {
    pub new_page_to_show: u8,
    pub total_pages: u8,
    pub position_of_lectern_to_update: BlockPos,
}

pub const LECTERNUPDATE_NEW_PAGE_TO_SHOW_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LECTERNUPDATE_TOTAL_PAGES_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LECTERNUPDATE_POSITION_OF_LECTERN_TO_UPDATE_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl LecternUpdate {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LecternUpdatePacket.New page to show", LECTERNUPDATE_NEW_PAGE_TO_SHOW_SHAPE);
        encoder.field("LecternUpdatePacket.Total Pages", LECTERNUPDATE_TOTAL_PAGES_SHAPE);
        encoder.field("LecternUpdatePacket.Position of Lectern to update", LECTERNUPDATE_POSITION_OF_LECTERN_TO_UPDATE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LecternUpdatePacket.New page to show", LECTERNUPDATE_NEW_PAGE_TO_SHOW_SHAPE);
        decoder.field("LecternUpdatePacket.Total Pages", LECTERNUPDATE_TOTAL_PAGES_SHAPE);
        decoder.field("LecternUpdatePacket.Position of Lectern to update", LECTERNUPDATE_POSITION_OF_LECTERN_TO_UPDATE_SHAPE);
    }
}
