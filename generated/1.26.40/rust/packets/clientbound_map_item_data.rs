// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundMapItemData {
    pub map_id: ActorUniqueID,
    pub dimension: u8,
    pub is_locked: bool,
    pub map_origin: BlockPos,
    pub creation_map_i_ds: Option<Vec<ActorUniqueID>>,
    pub scale: Option<i8>,
    pub tracked_actor_i_ds: Option<Vec<MapItemTrackedActorUniqueId>>,
    pub decorations: Option<Vec<MapDecoration>>,
    pub width: Option<i32>,
    pub height: Option<i32>,
    pub start_x: Option<i32>,
    pub start_y: Option<i32>,
    pub pixels: Option<Vec<u32>>,
}

pub const CLIENTBOUNDMAPITEMDATA_MAP_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CLIENTBOUNDMAPITEMDATA_DIMENSION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CLIENTBOUNDMAPITEMDATA_IS_LOCKED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CLIENTBOUNDMAPITEMDATA_MAP_ORIGIN_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CLIENTBOUNDMAPITEMDATA_CREATION_MAP_I_DS_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_SCALE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_TRACKED_ACTOR_I_DS_SHAPE: &str = r##"{"kind":"optional","value":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"MapItemTrackedActor::UniqueId","type_id":"MapItemTrackedActor::UniqueId","fields":[{"ordinal":0,"name":"Type","semantic":"Type","encode":{"kind":"enum","semantic":"MapItemTrackedActor::Type","type_id":"enums/MapItemTrackedActor::Type","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"},"variants":[{"value":0,"name":"Entity","encode":{"kind":"void"}},{"value":1,"name":"BlockEntity","encode":{"kind":"void"}},{"value":2,"name":"Other","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Entity ID","semantic":"Entity ID","type_id":"ActorUniqueID.json#","encode":{"kind":"optional","value":{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"Block Position","semantic":"Block Position","type_id":"BlockPos.json#","encode":{"kind":"optional","value":{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}}}"##;
pub const CLIENTBOUNDMAPITEMDATA_DECORATIONS_SHAPE: &str = r##"{"kind":"optional","value":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"MapDecoration","type_id":"MapDecoration","fields":[{"ordinal":0,"name":"Image Type","semantic":"Image Type","encode":{"kind":"enum","semantic":"MapDecoration::Type","type_id":"enums/MapDecoration::Type","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"MarkerWhite","encode":{"kind":"void"}},{"value":1,"name":"MarkerGreen","encode":{"kind":"void"}},{"value":2,"name":"MarkerRed","encode":{"kind":"void"}},{"value":3,"name":"MarkerBlue","encode":{"kind":"void"}},{"value":4,"name":"XWhite","encode":{"kind":"void"}},{"value":5,"name":"TriangleRed","encode":{"kind":"void"}},{"value":6,"name":"SquareWhite","encode":{"kind":"void"}},{"value":7,"name":"MarkerSign","encode":{"kind":"void"}},{"value":8,"name":"MarkerPink","encode":{"kind":"void"}},{"value":9,"name":"MarkerOrange","encode":{"kind":"void"}},{"value":10,"name":"MarkerYellow","encode":{"kind":"void"}},{"value":11,"name":"MarkerTeal","encode":{"kind":"void"}},{"value":12,"name":"TriangleGreen","encode":{"kind":"void"}},{"value":13,"name":"SmallSquareWhite","encode":{"kind":"void"}},{"value":14,"name":"Mansion","encode":{"kind":"void"}},{"value":15,"name":"Monument","encode":{"kind":"void"}},{"value":16,"name":"NoDraw","encode":{"kind":"void"}},{"value":17,"name":"VillageDesert","encode":{"kind":"void"}},{"value":18,"name":"VillagePlains","encode":{"kind":"void"}},{"value":19,"name":"VillageSavanna","encode":{"kind":"void"}},{"value":20,"name":"VillageSnowy","encode":{"kind":"void"}},{"value":21,"name":"VillageTaiga","encode":{"kind":"void"}},{"value":22,"name":"JungleTemple","encode":{"kind":"void"}},{"value":23,"name":"WitchHut","encode":{"kind":"void"}},{"value":24,"name":"TrialChambers","encode":{"kind":"void"}},{"value":25,"name":"Count","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Rotation","semantic":"Rotation","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":3,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":4,"name":"Label","semantic":"Label","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":5,"name":"Color","semantic":"Color","type_id":"Color.json#","encode":{"kind":"struct","semantic":"mce::Color","type_id":"mce::Color","fields":[{"ordinal":0,"name":"Color","semantic":"Color","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}}}"##;
pub const CLIENTBOUNDMAPITEMDATA_WIDTH_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_HEIGHT_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_START_X_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_START_Y_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}}"#;
pub const CLIENTBOUNDMAPITEMDATA_PIXELS_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}}"#;

impl ClientboundMapItemData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundMapItemDataPacket.Map ID", CLIENTBOUNDMAPITEMDATA_MAP_ID_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Dimension", CLIENTBOUNDMAPITEMDATA_DIMENSION_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Is Locked", CLIENTBOUNDMAPITEMDATA_IS_LOCKED_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Map Origin", CLIENTBOUNDMAPITEMDATA_MAP_ORIGIN_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Creation Map IDs", CLIENTBOUNDMAPITEMDATA_CREATION_MAP_I_DS_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Scale", CLIENTBOUNDMAPITEMDATA_SCALE_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Tracked Actor IDs", CLIENTBOUNDMAPITEMDATA_TRACKED_ACTOR_I_DS_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Decorations", CLIENTBOUNDMAPITEMDATA_DECORATIONS_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Width", CLIENTBOUNDMAPITEMDATA_WIDTH_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Height", CLIENTBOUNDMAPITEMDATA_HEIGHT_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Start X", CLIENTBOUNDMAPITEMDATA_START_X_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Start Y", CLIENTBOUNDMAPITEMDATA_START_Y_SHAPE);
        encoder.field("ClientboundMapItemDataPacket.Pixels", CLIENTBOUNDMAPITEMDATA_PIXELS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundMapItemDataPacket.Map ID", CLIENTBOUNDMAPITEMDATA_MAP_ID_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Dimension", CLIENTBOUNDMAPITEMDATA_DIMENSION_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Is Locked", CLIENTBOUNDMAPITEMDATA_IS_LOCKED_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Map Origin", CLIENTBOUNDMAPITEMDATA_MAP_ORIGIN_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Creation Map IDs", CLIENTBOUNDMAPITEMDATA_CREATION_MAP_I_DS_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Scale", CLIENTBOUNDMAPITEMDATA_SCALE_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Tracked Actor IDs", CLIENTBOUNDMAPITEMDATA_TRACKED_ACTOR_I_DS_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Decorations", CLIENTBOUNDMAPITEMDATA_DECORATIONS_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Width", CLIENTBOUNDMAPITEMDATA_WIDTH_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Height", CLIENTBOUNDMAPITEMDATA_HEIGHT_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Start X", CLIENTBOUNDMAPITEMDATA_START_X_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Start Y", CLIENTBOUNDMAPITEMDATA_START_Y_SHAPE);
        decoder.field("ClientboundMapItemDataPacket.Pixels", CLIENTBOUNDMAPITEMDATA_PIXELS_SHAPE);
    }
}
