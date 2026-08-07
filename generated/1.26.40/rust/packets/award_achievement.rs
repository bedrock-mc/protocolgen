// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AwardAchievement {
    pub achievement_id: i32,
}

pub const AWARDACHIEVEMENT_ACHIEVEMENT_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;

impl AwardAchievement {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AwardAchievementPacket.AchievementID", AWARDACHIEVEMENT_ACHIEVEMENT_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AwardAchievementPacket.AchievementID", AWARDACHIEVEMENT_ACHIEVEMENT_ID_SHAPE);
    }
}
