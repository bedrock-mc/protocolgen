// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDataDrivenUICloseScreen {
    pub form_id: Option<u32>,
}

pub const CLIENTBOUNDDATADRIVENUICLOSESCREEN_FORM_ID_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;

impl ClientboundDataDrivenUICloseScreen {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundDataDrivenUICloseScreenPacket.FormId", CLIENTBOUNDDATADRIVENUICLOSESCREEN_FORM_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundDataDrivenUICloseScreenPacket.FormId", CLIENTBOUNDDATADRIVENUICLOSESCREEN_FORM_ID_SHAPE);
    }
}
