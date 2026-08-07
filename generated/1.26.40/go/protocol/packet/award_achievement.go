// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AwardAchievement struct {
	AchievementID int32
}

// Marshal reads or writes AwardAchievement using its canonical wire layout.
func (x *AwardAchievement) Marshal(io protocol.IO) {
	io.Int32(&x.AchievementID)
}
