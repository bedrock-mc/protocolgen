// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDataDrivenUIShowScreen {
    pub screen_id: String,
    pub form_id: u32,
    pub data_instance_id: Option<u32>,
}

pub const CLIENTBOUNDDATADRIVENUISHOWSCREEN_SCREEN_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTBOUNDDATADRIVENUISHOWSCREEN_FORM_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const CLIENTBOUNDDATADRIVENUISHOWSCREEN_DATA_INSTANCE_ID_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;

impl ClientboundDataDrivenUIShowScreen {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundDataDrivenUIShowScreenPacket.ScreenId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_SCREEN_ID_SHAPE);
        encoder.field("ClientboundDataDrivenUIShowScreenPacket.FormId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_FORM_ID_SHAPE);
        encoder.field("ClientboundDataDrivenUIShowScreenPacket.DataInstanceId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_DATA_INSTANCE_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundDataDrivenUIShowScreenPacket.ScreenId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_SCREEN_ID_SHAPE);
        decoder.field("ClientboundDataDrivenUIShowScreenPacket.FormId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_FORM_ID_SHAPE);
        decoder.field("ClientboundDataDrivenUIShowScreenPacket.DataInstanceId", CLIENTBOUNDDATADRIVENUISHOWSCREEN_DATA_INSTANCE_ID_SHAPE);
    }
}
