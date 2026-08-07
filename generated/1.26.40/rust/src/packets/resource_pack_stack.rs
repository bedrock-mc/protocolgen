// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackStack {
    pub texture_pack_required: bool,
    pub texture_pack_list: Vec<PackInstanceId>,
    pub base_game_version: String,
    pub experiments: Experiments,
    pub include_editor_packs: bool,
}

impl ResourcePackStack {
    pub const ID: u32 = 7;
}
