// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CreatePhoto {
    pub raw_id: u64,
    pub photo_name: String,
    pub photo_item_name: String,
}

pub const CREATEPHOTO_RAW_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const CREATEPHOTO_PHOTO_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CREATEPHOTO_PHOTO_ITEM_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl CreatePhoto {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CreatePhotoPacket.Raw ID", CREATEPHOTO_RAW_ID_SHAPE);
        encoder.field("CreatePhotoPacket.Photo Name", CREATEPHOTO_PHOTO_NAME_SHAPE);
        encoder.field("CreatePhotoPacket.Photo Item Name", CREATEPHOTO_PHOTO_ITEM_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CreatePhotoPacket.Raw ID", CREATEPHOTO_RAW_ID_SHAPE);
        decoder.field("CreatePhotoPacket.Photo Name", CREATEPHOTO_PHOTO_NAME_SHAPE);
        decoder.field("CreatePhotoPacket.Photo Item Name", CREATEPHOTO_PHOTO_ITEM_NAME_SHAPE);
    }
}
