// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistActorPriority {
    pub camera_aim_assist_actor_priority_list: Vec<CameraAimAssistActorPriorityPriorityData>,
}

pub const CAMERAAIMASSISTACTORPRIORITY_CAMERA_AIM_ASSIST_ACTOR_PRIORITY_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"CameraAimAssistActorPriority::PriorityData","type_id":"CameraAimAssistActorPriority::PriorityData","fields":[{"ordinal":0,"name":"Preset Index","semantic":"Preset Index","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Category Index","semantic":"Category Index","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Actor Index","semantic":"Actor Index","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Priority Value","semantic":"Priority Value","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl CameraAimAssistActorPriority {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CameraAimAssistActorPriorityPacket.Camera Aim-Assist Actor Priority List", CAMERAAIMASSISTACTORPRIORITY_CAMERA_AIM_ASSIST_ACTOR_PRIORITY_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CameraAimAssistActorPriorityPacket.Camera Aim-Assist Actor Priority List", CAMERAAIMASSISTACTORPRIORITY_CAMERA_AIM_ASSIST_ACTOR_PRIORITY_LIST_SHAPE);
    }
}
