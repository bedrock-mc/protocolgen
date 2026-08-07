// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateAdventureSettings {
    pub adventure_settings: AdventureSettings,
}

pub const UPDATEADVENTURESETTINGS_ADVENTURE_SETTINGS_SHAPE: &str = r#"{"kind":"struct","semantic":"AdventureSettings","type_id":"AdventureSettings","fields":[{"ordinal":0,"name":"no PvM","semantic":"no PvM","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"no MvP","semantic":"no MvP","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Immutable World","semantic":"Immutable World","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Show Name Tags","semantic":"Show Name Tags","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":4,"name":"Auto Jump","semantic":"Auto Jump","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl UpdateAdventureSettings {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateAdventureSettingsPacket.Adventure Settings", UPDATEADVENTURESETTINGS_ADVENTURE_SETTINGS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateAdventureSettingsPacket.Adventure Settings", UPDATEADVENTURESETTINGS_ADVENTURE_SETTINGS_SHAPE);
    }
}
