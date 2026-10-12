// Code generated from canonical protocol manifest v2. DO NOT EDIT.

// Variant layouts follow the manifest; the emitter boxes by its own size estimate.
#![allow(clippy::large_enum_variant)]

pub const GAME_VERSION: &str = "1.26.51";
pub const PROTOCOL_VERSION: i32 = 2193;
/// Rust codec contract this crate implements; see protocolgen's
/// `docs/rust-codec-contract.md`.
pub const CODEC_CONTRACT: u32 = 1;
/// SHA-256 of the stable serialization of the manifest this crate was generated from.
pub const MANIFEST_SHA256: &str =
    "dc97931d683795fefffbe41b3be167b4ff8f820628323d18b42f97fb883c84df";

pub mod enums;
pub mod fields;
pub mod packets;
pub mod types;
pub mod wire;
