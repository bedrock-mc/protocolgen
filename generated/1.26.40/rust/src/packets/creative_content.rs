// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CreativeContent {
    pub groups: Vec<CreativeGroupInfo>,
    pub entries: Vec<CreativeItemEntry>,
}

impl CreativeContent {
    pub const ID: u32 = 145;
}
