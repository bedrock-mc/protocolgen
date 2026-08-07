// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameTestResults {
    pub succeeded: bool,
    pub error: String,
    pub test_name: String,
}

impl GameTestResults {
    pub const ID: u32 = 195;
}
