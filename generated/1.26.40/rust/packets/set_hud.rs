// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetHud {
    pub hud_element: Vec<HudElement>,
    pub hud_visible: HudVisibility,
}

pub const SETHUD_HUD_ELEMENT_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"enum","semantic":"HudElement","type_id":"enums/HudElement","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"PaperDoll","encode":{"kind":"void"}},{"value":1,"name":"Armor","encode":{"kind":"void"}},{"value":2,"name":"ToolTips","encode":{"kind":"void"}},{"value":3,"name":"TouchControls","encode":{"kind":"void"}},{"value":4,"name":"Crosshair","encode":{"kind":"void"}},{"value":5,"name":"HotBar","encode":{"kind":"void"}},{"value":6,"name":"Health","encode":{"kind":"void"}},{"value":7,"name":"ProgressBar","encode":{"kind":"void"}},{"value":8,"name":"Hunger","encode":{"kind":"void"}},{"value":9,"name":"AirBubbles","encode":{"kind":"void"}},{"value":10,"name":"HorseHealth","encode":{"kind":"void"}},{"value":11,"name":"StatusEffects","encode":{"kind":"void"}},{"value":12,"name":"ItemText","encode":{"kind":"void"}}]}}"#;
pub const SETHUD_HUD_VISIBLE_SHAPE: &str = r#"{"kind":"enum","semantic":"HudVisibility","type_id":"enums/HudVisibility","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"Hide","encode":{"kind":"void"}},{"value":1,"name":"Reset","encode":{"kind":"void"}}]}"#;

impl SetHud {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetHudPacket.Hud Element", SETHUD_HUD_ELEMENT_SHAPE);
        encoder.field("SetHudPacket.Hud Visible", SETHUD_HUD_VISIBLE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetHudPacket.Hud Element", SETHUD_HUD_ELEMENT_SHAPE);
        decoder.field("SetHudPacket.Hud Visible", SETHUD_HUD_VISIBLE_SHAPE);
    }
}
