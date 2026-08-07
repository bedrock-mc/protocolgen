// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetDisplayObjective {
    pub display_slot_name: String,
    pub objective_name: String,
    pub objective_display_name: String,
    pub criteria_name: String,
    pub sort_order: i32,
}

pub const SETDISPLAYOBJECTIVE_DISPLAY_SLOT_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETDISPLAYOBJECTIVE_OBJECTIVE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETDISPLAYOBJECTIVE_OBJECTIVE_DISPLAY_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETDISPLAYOBJECTIVE_CRITERIA_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETDISPLAYOBJECTIVE_SORT_ORDER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl SetDisplayObjective {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetDisplayObjectivePacket.DisplaySlotName", SETDISPLAYOBJECTIVE_DISPLAY_SLOT_NAME_SHAPE);
        encoder.field("SetDisplayObjectivePacket.ObjectiveName", SETDISPLAYOBJECTIVE_OBJECTIVE_NAME_SHAPE);
        encoder.field("SetDisplayObjectivePacket.ObjectiveDisplayName", SETDISPLAYOBJECTIVE_OBJECTIVE_DISPLAY_NAME_SHAPE);
        encoder.field("SetDisplayObjectivePacket.CriteriaName", SETDISPLAYOBJECTIVE_CRITERIA_NAME_SHAPE);
        encoder.field("SetDisplayObjectivePacket.SortOrder", SETDISPLAYOBJECTIVE_SORT_ORDER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetDisplayObjectivePacket.DisplaySlotName", SETDISPLAYOBJECTIVE_DISPLAY_SLOT_NAME_SHAPE);
        decoder.field("SetDisplayObjectivePacket.ObjectiveName", SETDISPLAYOBJECTIVE_OBJECTIVE_NAME_SHAPE);
        decoder.field("SetDisplayObjectivePacket.ObjectiveDisplayName", SETDISPLAYOBJECTIVE_OBJECTIVE_DISPLAY_NAME_SHAPE);
        decoder.field("SetDisplayObjectivePacket.CriteriaName", SETDISPLAYOBJECTIVE_CRITERIA_NAME_SHAPE);
        decoder.field("SetDisplayObjectivePacket.SortOrder", SETDISPLAYOBJECTIVE_SORT_ORDER_SHAPE);
    }
}
