// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePacksInfo {
    pub resource_pack_required: bool,
    pub has_addon_packs: bool,
    pub has_scripts: bool,
    pub force_disable_vibrant_visuals: bool,
    pub world_template_id_and_version: PackIdVersion,
    pub resource_packs: Vec<PackInfoData>,
}

impl ResourcePacksInfo {
    pub const ID: u32 = 6;
}
