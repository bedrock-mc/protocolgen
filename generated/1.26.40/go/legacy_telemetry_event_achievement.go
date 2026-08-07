// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LegacyTelemetryEventAchievement struct {
	AchievementID MinecraftEventingAchievementIds
}

func (LegacyTelemetryEventAchievement) isLegacyTelemetryEventEventData() {}

// Marshal reads or writes LegacyTelemetryEventAchievement using its canonical wire layout.
func (x *LegacyTelemetryEventAchievement) Marshal(io IO) {
	IntegerFunc(&x.AchievementID, io.Uint8)
}
