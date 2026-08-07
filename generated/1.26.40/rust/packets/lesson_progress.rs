// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LessonProgress {
    pub lesson_action: i32,
    pub score: i32,
    pub activity_id: String,
}

pub const LESSONPROGRESS_LESSON_ACTION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const LESSONPROGRESS_SCORE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const LESSONPROGRESS_ACTIVITY_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl LessonProgress {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LessonProgressPacket.Lesson Action", LESSONPROGRESS_LESSON_ACTION_SHAPE);
        encoder.field("LessonProgressPacket.Score", LESSONPROGRESS_SCORE_SHAPE);
        encoder.field("LessonProgressPacket.Activity Id", LESSONPROGRESS_ACTIVITY_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LessonProgressPacket.Lesson Action", LESSONPROGRESS_LESSON_ACTION_SHAPE);
        decoder.field("LessonProgressPacket.Score", LESSONPROGRESS_SCORE_SHAPE);
        decoder.field("LessonProgressPacket.Activity Id", LESSONPROGRESS_ACTIVITY_ID_SHAPE);
    }
}
