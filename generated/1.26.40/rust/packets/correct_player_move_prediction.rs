// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CorrectPlayerMovePrediction {
    pub prediction_type: RewindType,
    pub pos: glam::Vec3,
    pub pos_delta: glam::Vec3,
    pub rotation: glam::Vec2,
    pub vehicle_angular_velocity: Option<f32>,
    pub on_ground: bool,
    pub tick: PlayerInputTick,
}

pub const CORRECTPLAYERMOVEPREDICTION_PREDICTION_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"RewindType","type_id":"enums/RewindType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Player","encode":{"kind":"void"}},{"value":1,"name":"Vehicle","encode":{"kind":"void"}}]}"#;
pub const CORRECTPLAYERMOVEPREDICTION_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CORRECTPLAYERMOVEPREDICTION_POS_DELTA_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CORRECTPLAYERMOVEPREDICTION_ROTATION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec2","type_id":"Vec2","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CORRECTPLAYERMOVEPREDICTION_VEHICLE_ANGULAR_VELOCITY_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;
pub const CORRECTPLAYERMOVEPREDICTION_ON_GROUND_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CORRECTPLAYERMOVEPREDICTION_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl CorrectPlayerMovePrediction {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CorrectPlayerMovePredictionPacket.PredictionType", CORRECTPLAYERMOVEPREDICTION_PREDICTION_TYPE_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.Pos", CORRECTPLAYERMOVEPREDICTION_POS_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.Pos Delta", CORRECTPLAYERMOVEPREDICTION_POS_DELTA_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.Rotation", CORRECTPLAYERMOVEPREDICTION_ROTATION_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.VehicleAngularVelocity", CORRECTPLAYERMOVEPREDICTION_VEHICLE_ANGULAR_VELOCITY_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.On Ground", CORRECTPLAYERMOVEPREDICTION_ON_GROUND_SHAPE);
        encoder.field("CorrectPlayerMovePredictionPacket.Tick", CORRECTPLAYERMOVEPREDICTION_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CorrectPlayerMovePredictionPacket.PredictionType", CORRECTPLAYERMOVEPREDICTION_PREDICTION_TYPE_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.Pos", CORRECTPLAYERMOVEPREDICTION_POS_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.Pos Delta", CORRECTPLAYERMOVEPREDICTION_POS_DELTA_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.Rotation", CORRECTPLAYERMOVEPREDICTION_ROTATION_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.VehicleAngularVelocity", CORRECTPLAYERMOVEPREDICTION_VEHICLE_ANGULAR_VELOCITY_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.On Ground", CORRECTPLAYERMOVEPREDICTION_ON_GROUND_SHAPE);
        decoder.field("CorrectPlayerMovePredictionPacket.Tick", CORRECTPLAYERMOVEPREDICTION_TICK_SHAPE);
    }
}
