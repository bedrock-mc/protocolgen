//! Protocol-schema frontends retained for the v2 migration reference.
//!
//! The old Prismarine/minecraft-data frontend was intentionally removed from
//! this lift. Only the Mojang/bpd-corrected and Endstone frontends remain.

use crate::ir::{Packet, Type};

pub mod endstone;
pub mod mojang;

pub struct ParseResult {
    pub packets: Vec<Packet>,
    pub types: std::collections::HashMap<String, Type>,
}
