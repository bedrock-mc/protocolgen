// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GraphicsOverrideParameter {
    pub parameter_keyframe_values: Vec<(f32, Vec3)>,
    pub float_value: Option<f32>,
    pub vec3_value: Option<Vec3>,
    pub biome_identifier: String,
    pub player_identifier: Option<String>,
    pub identifier_for_parameter: GraphicsOverrideParameterType,
    pub reset_parameter: bool,
}

pub const GRAPHICSOVERRIDEPARAMETER_PARAMETER_KEYFRAME_VALUES_SHAPE: &str = r#"{"kind":"map","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"representation":"ordered_entries","value":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"key":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;
pub const GRAPHICSOVERRIDEPARAMETER_FLOAT_VALUE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;
pub const GRAPHICSOVERRIDEPARAMETER_VEC3_VALUE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const GRAPHICSOVERRIDEPARAMETER_BIOME_IDENTIFIER_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const GRAPHICSOVERRIDEPARAMETER_PLAYER_IDENTIFIER_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;
pub const GRAPHICSOVERRIDEPARAMETER_IDENTIFIER_FOR_PARAMETER_SHAPE: &str = r#"{"kind":"enum","semantic":"GraphicsOverrideParameterType","type_id":"enums/GraphicsOverrideParameterType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"SkyZenithColor","encode":{"kind":"void"}},{"value":1,"name":"SkyHorizonColor","encode":{"kind":"void"}},{"value":2,"name":"HorizonBlendMin","encode":{"kind":"void"}},{"value":3,"name":"HorizonBlendMax","encode":{"kind":"void"}},{"value":4,"name":"HorizonBlendStart","encode":{"kind":"void"}},{"value":5,"name":"HorizonBlendMieStart","encode":{"kind":"void"}},{"value":6,"name":"RayleighStrength","encode":{"kind":"void"}},{"value":7,"name":"SunMieStrength","encode":{"kind":"void"}},{"value":8,"name":"MoonMieStrength","encode":{"kind":"void"}},{"value":9,"name":"SunGlareShape","encode":{"kind":"void"}},{"value":10,"name":"Chlorophyll","encode":{"kind":"void"}},{"value":11,"name":"CDOM","encode":{"kind":"void"}},{"value":12,"name":"SuspendedSediment","encode":{"kind":"void"}},{"value":13,"name":"WavesDepth","encode":{"kind":"void"}},{"value":14,"name":"WavesFrequency","encode":{"kind":"void"}},{"value":15,"name":"WavesFrequencyScaling","encode":{"kind":"void"}},{"value":16,"name":"WavesSpeed","encode":{"kind":"void"}},{"value":17,"name":"WavesSpeedScaling","encode":{"kind":"void"}},{"value":18,"name":"WavesShape","encode":{"kind":"void"}},{"value":19,"name":"WavesOctaves","encode":{"kind":"void"}},{"value":20,"name":"WavesMix","encode":{"kind":"void"}},{"value":21,"name":"WavesPull","encode":{"kind":"void"}},{"value":22,"name":"WavesDirectionIncrement","encode":{"kind":"void"}},{"value":23,"name":"MidtonesContrast","encode":{"kind":"void"}},{"value":24,"name":"HighlightsContrast","encode":{"kind":"void"}},{"value":25,"name":"ShadowsContrast","encode":{"kind":"void"}},{"value":26,"name":"HighlightsGain","encode":{"kind":"void"}},{"value":27,"name":"HighlightsGamma","encode":{"kind":"void"}},{"value":28,"name":"HighlightsOffset","encode":{"kind":"void"}},{"value":29,"name":"HighlightsSaturation","encode":{"kind":"void"}},{"value":30,"name":"MidtonesGain","encode":{"kind":"void"}},{"value":31,"name":"MidtonesGamma","encode":{"kind":"void"}},{"value":32,"name":"MidtonesOffset","encode":{"kind":"void"}},{"value":33,"name":"MidtonesSaturation","encode":{"kind":"void"}},{"value":34,"name":"ShadowsGain","encode":{"kind":"void"}},{"value":35,"name":"ShadowsGamma","encode":{"kind":"void"}},{"value":36,"name":"ShadowsOffset","encode":{"kind":"void"}},{"value":37,"name":"ShadowsSaturation","encode":{"kind":"void"}},{"value":38,"name":"HighlightsMin","encode":{"kind":"void"}},{"value":39,"name":"ShadowsMax","encode":{"kind":"void"}},{"value":40,"name":"Temperature","encode":{"kind":"void"}},{"value":41,"name":"SunColor","encode":{"kind":"void"}},{"value":42,"name":"SunIlluminance","encode":{"kind":"void"}},{"value":43,"name":"MoonColor","encode":{"kind":"void"}},{"value":44,"name":"MoonIlluminance","encode":{"kind":"void"}},{"value":45,"name":"FlashColor","encode":{"kind":"void"}},{"value":46,"name":"FlashIlluminance","encode":{"kind":"void"}},{"value":47,"name":"AmbientColor","encode":{"kind":"void"}},{"value":48,"name":"AmbientIlluminance","encode":{"kind":"void"}},{"value":49,"name":"EmissiveDesaturation","encode":{"kind":"void"}},{"value":50,"name":"SkyIntensity","encode":{"kind":"void"}},{"value":51,"name":"OrbitalOffsetDegrees","encode":{"kind":"void"}}]}"#;
pub const GRAPHICSOVERRIDEPARAMETER_RESET_PARAMETER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl GraphicsOverrideParameter {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("GraphicsOverrideParameterPacket.Parameter Keyframe Values", GRAPHICSOVERRIDEPARAMETER_PARAMETER_KEYFRAME_VALUES_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Float Value", GRAPHICSOVERRIDEPARAMETER_FLOAT_VALUE_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Vec3 Value", GRAPHICSOVERRIDEPARAMETER_VEC3_VALUE_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Biome Identifier", GRAPHICSOVERRIDEPARAMETER_BIOME_IDENTIFIER_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Player Identifier", GRAPHICSOVERRIDEPARAMETER_PLAYER_IDENTIFIER_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Identifier for Parameter", GRAPHICSOVERRIDEPARAMETER_IDENTIFIER_FOR_PARAMETER_SHAPE);
        encoder.field("GraphicsOverrideParameterPacket.Reset Parameter", GRAPHICSOVERRIDEPARAMETER_RESET_PARAMETER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("GraphicsOverrideParameterPacket.Parameter Keyframe Values", GRAPHICSOVERRIDEPARAMETER_PARAMETER_KEYFRAME_VALUES_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Float Value", GRAPHICSOVERRIDEPARAMETER_FLOAT_VALUE_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Vec3 Value", GRAPHICSOVERRIDEPARAMETER_VEC3_VALUE_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Biome Identifier", GRAPHICSOVERRIDEPARAMETER_BIOME_IDENTIFIER_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Player Identifier", GRAPHICSOVERRIDEPARAMETER_PLAYER_IDENTIFIER_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Identifier for Parameter", GRAPHICSOVERRIDEPARAMETER_IDENTIFIER_FOR_PARAMETER_SHAPE);
        decoder.field("GraphicsOverrideParameterPacket.Reset Parameter", GRAPHICSOVERRIDEPARAMETER_RESET_PARAMETER_SHAPE);
    }
}
