// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerAction {
    pub player_runtime_id: ActorRuntimeID,
    pub action: PlayerActionType,
    pub block_position: BlockPos,
    pub result_pos: BlockPos,
    pub face: i32,
}

pub const PLAYERACTION_PLAYER_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYERACTION_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"PlayerActionType","type_id":"enums/PlayerActionType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":-1,"name":"Unknown","encode":{"kind":"void"}},{"value":0,"name":"StartDestroyBlock","encode":{"kind":"void"}},{"value":1,"name":"AbortDestroyBlock","encode":{"kind":"void"}},{"value":2,"name":"StopDestroyBlock","encode":{"kind":"void"}},{"value":3,"name":"GetUpdatedBlock","encode":{"kind":"void"}},{"value":4,"name":"DropItem","encode":{"kind":"void"}},{"value":5,"name":"StartSleeping","encode":{"kind":"void"}},{"value":6,"name":"StopSleeping","encode":{"kind":"void"}},{"value":7,"name":"Respawn","encode":{"kind":"void"}},{"value":8,"name":"StartJump","encode":{"kind":"void"}},{"value":9,"name":"StartSprinting","encode":{"kind":"void"}},{"value":10,"name":"StopSprinting","encode":{"kind":"void"}},{"value":11,"name":"StartSneaking","encode":{"kind":"void"}},{"value":12,"name":"StopSneaking","encode":{"kind":"void"}},{"value":13,"name":"CreativeDestroyBlock","encode":{"kind":"void"}},{"value":14,"name":"ChangeDimensionAck","encode":{"kind":"void"}},{"value":15,"name":"StartGliding","encode":{"kind":"void"}},{"value":16,"name":"StopGliding","encode":{"kind":"void"}},{"value":17,"name":"DenyDestroyBlock","encode":{"kind":"void"}},{"value":18,"name":"CrackBlock","encode":{"kind":"void"}},{"value":19,"name":"ChangeSkin","encode":{"kind":"void"}},{"value":20,"name":"UpdatedEnchantingSeed","encode":{"kind":"void"}},{"value":21,"name":"StartSwimming","encode":{"kind":"void"}},{"value":22,"name":"StopSwimming","encode":{"kind":"void"}},{"value":23,"name":"StartSpinAttack","encode":{"kind":"void"}},{"value":24,"name":"StopSpinAttack","encode":{"kind":"void"}},{"value":25,"name":"InteractWithBlock","encode":{"kind":"void"}},{"value":26,"name":"PredictDestroyBlock","encode":{"kind":"void"}},{"value":27,"name":"ContinueDestroyBlock","encode":{"kind":"void"}},{"value":28,"name":"StartItemUseOn","encode":{"kind":"void"}},{"value":29,"name":"StopItemUseOn","encode":{"kind":"void"}},{"value":30,"name":"HandledTeleport","encode":{"kind":"void"}},{"value":31,"name":"MissedSwing","encode":{"kind":"void"}},{"value":32,"name":"StartCrawling","encode":{"kind":"void"}},{"value":33,"name":"StopCrawling","encode":{"kind":"void"}},{"value":34,"name":"StartFlying","encode":{"kind":"void"}},{"value":35,"name":"StopFlying","encode":{"kind":"void"}},{"value":36,"name":"ClientAckServerData","encode":{"kind":"void"}},{"value":37,"name":"StartUsingItem","encode":{"kind":"void"}},{"value":38,"name":"InternalUpdate","encode":{"kind":"void"}},{"value":39,"name":"Count","encode":{"kind":"void"}}]}"#;
pub const PLAYERACTION_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYERACTION_RESULT_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYERACTION_FACE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl PlayerAction {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlayerActionPacket.Player Runtime ID", PLAYERACTION_PLAYER_RUNTIME_ID_SHAPE);
        encoder.field("PlayerActionPacket.Action", PLAYERACTION_ACTION_SHAPE);
        encoder.field("PlayerActionPacket.Block Position", PLAYERACTION_BLOCK_POSITION_SHAPE);
        encoder.field("PlayerActionPacket.Result Pos", PLAYERACTION_RESULT_POS_SHAPE);
        encoder.field("PlayerActionPacket.Face", PLAYERACTION_FACE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlayerActionPacket.Player Runtime ID", PLAYERACTION_PLAYER_RUNTIME_ID_SHAPE);
        decoder.field("PlayerActionPacket.Action", PLAYERACTION_ACTION_SHAPE);
        decoder.field("PlayerActionPacket.Block Position", PLAYERACTION_BLOCK_POSITION_SHAPE);
        decoder.field("PlayerActionPacket.Result Pos", PLAYERACTION_RESULT_POS_SHAPE);
        decoder.field("PlayerActionPacket.Face", PLAYERACTION_FACE_SHAPE);
    }
}
