// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MovePlayer {
    pub player_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: f32,
    pub position_mode: PlayerPositionModeComponentPositionMode,
    pub on_ground: bool,
    pub riding_runtime_id: ActorRuntimeID,
    pub teleport_data: Option<MovePlayerTeleportData>,
    pub tick: PlayerInputTick,
}

pub const MOVEPLAYER_PLAYER_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOVEPLAYER_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOVEPLAYER_ROTATION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec2","type_id":"Vec2","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOVEPLAYER_Y_HEAD_ROTATION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const MOVEPLAYER_POSITION_MODE_SHAPE: &str = r#"{"kind":"enum","semantic":"PlayerPositionModeComponent::PositionMode","type_id":"enums/PlayerPositionModeComponent::PositionMode","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Normal","encode":{"kind":"void"}},{"value":1,"name":"Respawn","encode":{"kind":"void"}},{"value":2,"name":"Teleport","encode":{"kind":"void"}},{"value":3,"name":"OnlyHeadRot","encode":{"kind":"void"}}]}"#;
pub const MOVEPLAYER_ON_GROUND_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const MOVEPLAYER_RIDING_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOVEPLAYER_TELEPORT_DATA_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"MovePlayerTeleportData","type_id":"MovePlayerTeleportData","fields":[{"ordinal":0,"name":"Teleportation Cause","semantic":"Teleportation Cause","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Source Actor Type","semantic":"Source Actor Type","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const MOVEPLAYER_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl MovePlayer {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MovePlayerPacket.Player Runtime ID", MOVEPLAYER_PLAYER_RUNTIME_ID_SHAPE);
        encoder.field("MovePlayerPacket.Position", MOVEPLAYER_POSITION_SHAPE);
        encoder.field("MovePlayerPacket.Rotation", MOVEPLAYER_ROTATION_SHAPE);
        encoder.field("MovePlayerPacket.Y-Head Rotation", MOVEPLAYER_Y_HEAD_ROTATION_SHAPE);
        encoder.field("MovePlayerPacket.Position Mode", MOVEPLAYER_POSITION_MODE_SHAPE);
        encoder.field("MovePlayerPacket.On Ground", MOVEPLAYER_ON_GROUND_SHAPE);
        encoder.field("MovePlayerPacket.Riding Runtime ID", MOVEPLAYER_RIDING_RUNTIME_ID_SHAPE);
        encoder.field("MovePlayerPacket.Teleport Data", MOVEPLAYER_TELEPORT_DATA_SHAPE);
        encoder.field("MovePlayerPacket.Tick", MOVEPLAYER_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MovePlayerPacket.Player Runtime ID", MOVEPLAYER_PLAYER_RUNTIME_ID_SHAPE);
        decoder.field("MovePlayerPacket.Position", MOVEPLAYER_POSITION_SHAPE);
        decoder.field("MovePlayerPacket.Rotation", MOVEPLAYER_ROTATION_SHAPE);
        decoder.field("MovePlayerPacket.Y-Head Rotation", MOVEPLAYER_Y_HEAD_ROTATION_SHAPE);
        decoder.field("MovePlayerPacket.Position Mode", MOVEPLAYER_POSITION_MODE_SHAPE);
        decoder.field("MovePlayerPacket.On Ground", MOVEPLAYER_ON_GROUND_SHAPE);
        decoder.field("MovePlayerPacket.Riding Runtime ID", MOVEPLAYER_RIDING_RUNTIME_ID_SHAPE);
        decoder.field("MovePlayerPacket.Teleport Data", MOVEPLAYER_TELEPORT_DATA_SHAPE);
        decoder.field("MovePlayerPacket.Tick", MOVEPLAYER_TICK_SHAPE);
    }
}
