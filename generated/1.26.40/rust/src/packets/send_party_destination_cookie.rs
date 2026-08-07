// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SendPartyDestinationCookie {
    pub cookie: String,
    pub intent: String,
    pub destination_name: String,
}

impl SendPartyDestinationCookie {
    pub const ID: u32 = 349;
}
