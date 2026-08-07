// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PhotoTransfer {
    pub photo_name: String,
    pub photo_data: String,
    pub book_id: String,
    pub r#type: PhotoType,
    pub source_type: PhotoType,
    pub owner_id: i64,
    pub new_photo_name: String,
}

pub const PHOTOTRANSFER_PHOTO_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PHOTOTRANSFER_PHOTO_DATA_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PHOTOTRANSFER_BOOK_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PHOTOTRANSFER_R_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"PhotoType","type_id":"enums/PhotoType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Portfolio","encode":{"kind":"void"}},{"value":1,"name":"PhotoItem","encode":{"kind":"void"}},{"value":2,"name":"Book","encode":{"kind":"void"}}]}"#;
pub const PHOTOTRANSFER_SOURCE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"PhotoType","type_id":"enums/PhotoType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Portfolio","encode":{"kind":"void"}},{"value":1,"name":"PhotoItem","encode":{"kind":"void"}},{"value":2,"name":"Book","encode":{"kind":"void"}}]}"#;
pub const PHOTOTRANSFER_OWNER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const PHOTOTRANSFER_NEW_PHOTO_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl PhotoTransfer {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PhotoTransferPacket.Photo Name", PHOTOTRANSFER_PHOTO_NAME_SHAPE);
        encoder.field("PhotoTransferPacket.Photo Data", PHOTOTRANSFER_PHOTO_DATA_SHAPE);
        encoder.field("PhotoTransferPacket.Book ID", PHOTOTRANSFER_BOOK_ID_SHAPE);
        encoder.field("PhotoTransferPacket.Type", PHOTOTRANSFER_R_TYPE_SHAPE);
        encoder.field("PhotoTransferPacket.Source Type", PHOTOTRANSFER_SOURCE_TYPE_SHAPE);
        encoder.field("PhotoTransferPacket.Owner ID", PHOTOTRANSFER_OWNER_ID_SHAPE);
        encoder.field("PhotoTransferPacket.New Photo Name", PHOTOTRANSFER_NEW_PHOTO_NAME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PhotoTransferPacket.Photo Name", PHOTOTRANSFER_PHOTO_NAME_SHAPE);
        decoder.field("PhotoTransferPacket.Photo Data", PHOTOTRANSFER_PHOTO_DATA_SHAPE);
        decoder.field("PhotoTransferPacket.Book ID", PHOTOTRANSFER_BOOK_ID_SHAPE);
        decoder.field("PhotoTransferPacket.Type", PHOTOTRANSFER_R_TYPE_SHAPE);
        decoder.field("PhotoTransferPacket.Source Type", PHOTOTRANSFER_SOURCE_TYPE_SHAPE);
        decoder.field("PhotoTransferPacket.Owner ID", PHOTOTRANSFER_OWNER_ID_SHAPE);
        decoder.field("PhotoTransferPacket.New Photo Name", PHOTOTRANSFER_NEW_PHOTO_NAME_SHAPE);
    }
}
