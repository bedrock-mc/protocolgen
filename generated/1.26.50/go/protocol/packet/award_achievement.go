// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// AwardAchievement is sent by the server to award an achievement to a player.
type AwardAchievement struct {
	// AchievementID is the ID of the achievement that should be awarded to the player. The values for
	// these IDs are currently unknown.
	AchievementID int32
}

// Marshal reads or writes AwardAchievement using its canonical wire layout.
func (x *AwardAchievement) Marshal(io protocol.IO) {
	io.Int32(&x.AchievementID)
}

// ID returns the protocol ID for AwardAchievement.
func (*AwardAchievement) ID() uint32 { return IDAwardAchievement }
