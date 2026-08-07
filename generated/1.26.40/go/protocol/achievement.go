// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type Achievement struct {
	AchievementID MinecraftEventingAchievementIds
}

func (*Achievement) isEvent() {}

// Marshal reads or writes Achievement using its canonical wire layout.
func (x *Achievement) Marshal(io IO) {
	IntegerFunc(&x.AchievementID, io.Uint8)
}
