// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct EmoteList {
    pub runtime_id: ActorRuntimeID,
    pub emote_piece_ids: Vec<uuid::Uuid>,
}

impl EmoteList {
    pub const ID: u32 = 152;
}
