// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AgentAnimation {
    pub agent_animation: AgentAnimationType,
    pub runtime_id: ActorRuntimeID,
}

pub const AGENTANIMATION_AGENT_ANIMATION_SHAPE: &str = r#"{"kind":"enum","semantic":"AgentAnimation","type_id":"enums/AgentAnimation","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"ArmSwing","encode":{"kind":"void"}},{"value":1,"name":"Shrug","encode":{"kind":"void"}}]}"#;
pub const AGENTANIMATION_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl AgentAnimation {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AgentAnimationPacket.Agent Animation", AGENTANIMATION_AGENT_ANIMATION_SHAPE);
        encoder.field("AgentAnimationPacket.Runtime Id", AGENTANIMATION_RUNTIME_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AgentAnimationPacket.Agent Animation", AGENTANIMATION_AGENT_ANIMATION_SHAPE);
        decoder.field("AgentAnimationPacket.Runtime Id", AGENTANIMATION_RUNTIME_ID_SHAPE);
    }
}
