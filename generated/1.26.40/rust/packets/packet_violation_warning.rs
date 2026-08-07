// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PacketViolationWarning {
    pub violation_type: PacketViolationType,
    pub violation_severity: PacketViolationSeverity,
    pub violation_packet_id: i32,
    pub violation_context: String,
}

pub const PACKETVIOLATIONWARNING_VIOLATION_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"PacketViolationType","type_id":"enums/PacketViolationType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":-1,"name":"Unknown","encode":{"kind":"void"}},{"value":0,"name":"PacketMalformed","encode":{"kind":"void"}}]}"#;
pub const PACKETVIOLATIONWARNING_VIOLATION_SEVERITY_SHAPE: &str = r#"{"kind":"enum","semantic":"PacketViolationSeverity","type_id":"enums/PacketViolationSeverity","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":-1,"name":"Unknown","encode":{"kind":"void"}},{"value":0,"name":"Warning","encode":{"kind":"void"}},{"value":1,"name":"FinalWarning","encode":{"kind":"void"}},{"value":2,"name":"TerminatingConnection","encode":{"kind":"void"}}]}"#;
pub const PACKETVIOLATIONWARNING_VIOLATION_PACKET_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const PACKETVIOLATIONWARNING_VIOLATION_CONTEXT_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl PacketViolationWarning {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PacketViolationWarningPacket.Violation Type", PACKETVIOLATIONWARNING_VIOLATION_TYPE_SHAPE);
        encoder.field("PacketViolationWarningPacket.Violation Severity", PACKETVIOLATIONWARNING_VIOLATION_SEVERITY_SHAPE);
        encoder.field("PacketViolationWarningPacket.Violation PacketId", PACKETVIOLATIONWARNING_VIOLATION_PACKET_ID_SHAPE);
        encoder.field("PacketViolationWarningPacket.Violation Context", PACKETVIOLATIONWARNING_VIOLATION_CONTEXT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PacketViolationWarningPacket.Violation Type", PACKETVIOLATIONWARNING_VIOLATION_TYPE_SHAPE);
        decoder.field("PacketViolationWarningPacket.Violation Severity", PACKETVIOLATIONWARNING_VIOLATION_SEVERITY_SHAPE);
        decoder.field("PacketViolationWarningPacket.Violation PacketId", PACKETVIOLATIONWARNING_VIOLATION_PACKET_ID_SHAPE);
        decoder.field("PacketViolationWarningPacket.Violation Context", PACKETVIOLATIONWARNING_VIOLATION_CONTEXT_SHAPE);
    }
}
