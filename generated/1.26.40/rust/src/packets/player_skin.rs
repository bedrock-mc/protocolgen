// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerSkin {
    pub uuid: uuid::Uuid,
    pub serialized_skin: SerializedSkinRef,
    pub localized_new_skin_name: String,
    pub localized_old_skin_name: String,
}

impl PlayerSkin {
    pub const ID: u32 = 93;
}
