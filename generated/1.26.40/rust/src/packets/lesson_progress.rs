// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LessonProgress {
    pub lesson_action: i32,
    pub score: i32,
    pub activity_id: String,
}

impl LessonProgress {
    pub const ID: u32 = 183;
}
