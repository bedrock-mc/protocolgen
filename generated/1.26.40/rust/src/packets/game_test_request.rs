// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct GameTestRequest {
    pub max_tests_per_batch: i32,
    pub repeat_count: i32,
    pub rotation: Rotation,
    pub stop_on_failure: bool,
    pub test_pos: BlockPos,
    pub tests_per_row: i32,
    pub test_name: String,
}

impl GameTestRequest {
    pub const ID: u32 = 194;
}
