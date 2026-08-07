// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetScore {
    pub score_info: Vec<SetScoreScoreInfoItem>,
}

impl SetScore {
    pub const ID: u32 = 108;
}
