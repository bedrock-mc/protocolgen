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

impl PacketViolationWarning {
    pub const ID: u32 = 156;
}
