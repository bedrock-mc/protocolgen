// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct OnScreenTextureAnimation {
    pub effect_id: u32,
}

pub const ONSCREENTEXTUREANIMATION_EFFECT_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl OnScreenTextureAnimation {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("OnScreenTextureAnimationPacket.Effect Id", ONSCREENTEXTUREANIMATION_EFFECT_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("OnScreenTextureAnimationPacket.Effect Id", ONSCREENTEXTUREANIMATION_EFFECT_ID_SHAPE);
    }
}
