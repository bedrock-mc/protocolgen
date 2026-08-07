// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundDebugRenderer {
    pub r#type: String,
    pub debug_marker_data: Option<ClientboundDebugRendererDebugMarkerData>,
}

pub const CLIENTBOUNDDEBUGRENDERER_R_TYPE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTBOUNDDEBUGRENDERER_DEBUG_MARKER_DATA_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"ClientboundDebugRendererPacketPayload::DebugMarkerData","type_id":"ClientboundDebugRendererPacketPayload::DebugMarkerData","fields":[{"ordinal":0,"name":"Text","semantic":"Text","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Position","semantic":"Position","encode":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Color","semantic":"Color","encode":{"kind":"struct","semantic":"mce::Color","type_id":"mce::Color","fields":[{"ordinal":0,"name":"Color","semantic":"Color","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"duration","semantic":"duration","encode":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl ClientboundDebugRenderer {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundDebugRendererPacket.Type", CLIENTBOUNDDEBUGRENDERER_R_TYPE_SHAPE);
        encoder.field("ClientboundDebugRendererPacket.DebugMarkerData", CLIENTBOUNDDEBUGRENDERER_DEBUG_MARKER_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundDebugRendererPacket.Type", CLIENTBOUNDDEBUGRENDERER_R_TYPE_SHAPE);
        decoder.field("ClientboundDebugRendererPacket.DebugMarkerData", CLIENTBOUNDDEBUGRENDERER_DEBUG_MARKER_DATA_SHAPE);
    }
}
