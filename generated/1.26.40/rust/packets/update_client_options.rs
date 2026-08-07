// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateClientOptions {
    pub graphics_mode_change: Option<GraphicsMode>,
    pub filter_profanity_change: Option<bool>,
}

pub const UPDATECLIENTOPTIONS_GRAPHICS_MODE_CHANGE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"enum","semantic":"GraphicsMode","type_id":"enums/GraphicsMode","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Simple","encode":{"kind":"void"}},{"value":1,"name":"Fancy","encode":{"kind":"void"}},{"value":2,"name":"Advanced","encode":{"kind":"void"}},{"value":3,"name":"RayTraced","encode":{"kind":"void"}}]}}"#;
pub const UPDATECLIENTOPTIONS_FILTER_PROFANITY_CHANGE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}}"#;

impl UpdateClientOptions {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateClientOptionsPacket.Graphics Mode Change", UPDATECLIENTOPTIONS_GRAPHICS_MODE_CHANGE_SHAPE);
        encoder.field("UpdateClientOptionsPacket.Filter Profanity Change", UPDATECLIENTOPTIONS_FILTER_PROFANITY_CHANGE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateClientOptionsPacket.Graphics Mode Change", UPDATECLIENTOPTIONS_GRAPHICS_MODE_CHANGE_SHAPE);
        decoder.field("UpdateClientOptionsPacket.Filter Profanity Change", UPDATECLIENTOPTIONS_FILTER_PROFANITY_CHANGE_SHAPE);
    }
}
