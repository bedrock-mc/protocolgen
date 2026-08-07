// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SimulationType {
    pub sim_type: SimulationTypeType,
}

pub const SIMULATIONTYPE_SIM_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"SimulationType","type_id":"enums/SimulationType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Game","encode":{"kind":"void"}},{"value":1,"name":"Editor","encode":{"kind":"void"}},{"value":2,"name":"Test","encode":{"kind":"void"}},{"value":3,"name":"INVALID","encode":{"kind":"void"}}]}"#;

impl SimulationType {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SimulationTypePacket.Sim Type", SIMULATIONTYPE_SIM_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SimulationTypePacket.Sim Type", SIMULATIONTYPE_SIM_TYPE_SHAPE);
    }
}
