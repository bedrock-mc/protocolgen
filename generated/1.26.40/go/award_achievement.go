// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AwardAchievement struct {
	AchievementID int32
}

func (p *AwardAchievement) Encode(w Encoder) error {
	if err := w.Write("AwardAchievementPacket.AchievementID", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.AchievementID); err != nil {
		return err
	}
	return nil
}

func DecodeAwardAchievement(r Decoder) (AwardAchievement, error) {
	var p AwardAchievement
	{
		raw, err := r.Read("AwardAchievementPacket.AchievementID", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field AwardAchievementPacket.AchievementID has unexpected decoded type %T", raw)
		}
		p.AchievementID = value
	}
	return p, nil
}
