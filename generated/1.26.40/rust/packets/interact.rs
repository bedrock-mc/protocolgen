// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Interact {
    pub action: InteractAction,
    pub target_runtime_id: ActorRuntimeID,
    pub position: Option<Vec3>,
}

pub const INTERACT_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"InteractPacketPayload::Action","type_id":"enums/InteractPacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Invalid","encode":{"kind":"void"}},{"value":3,"name":"StopRiding","encode":{"kind":"void"}},{"value":4,"name":"InteractUpdate","encode":{"kind":"void"}},{"value":5,"name":"NpcOpen","encode":{"kind":"void"}},{"value":6,"name":"OpenInventory","encode":{"kind":"void"}}]}"#;
pub const INTERACT_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const INTERACT_POSITION_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl Interact {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("InteractPacket.Action", INTERACT_ACTION_SHAPE);
        encoder.field("InteractPacket.Target Runtime ID", INTERACT_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("InteractPacket.Position", INTERACT_POSITION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("InteractPacket.Action", INTERACT_ACTION_SHAPE);
        decoder.field("InteractPacket.Target Runtime ID", INTERACT_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("InteractPacket.Position", INTERACT_POSITION_SHAPE);
    }
}
