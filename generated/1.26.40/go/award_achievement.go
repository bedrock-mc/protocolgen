// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AwardAchievement struct {
	AchievementID int32
}

// Marshal reads or writes AwardAchievement using its canonical wire layout.
func (x *AwardAchievement) Marshal(io IO) {
	io.Int32(&x.AchievementID)
}
