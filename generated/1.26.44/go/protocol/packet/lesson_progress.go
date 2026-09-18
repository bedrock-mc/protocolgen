// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// LessonProgress is a packet sent by the server to the client to inform the client of updated progress on a
// lesson. This packet only functions on the Minecraft: Education Edition version of the game.
type LessonProgress struct {
	LessonAction int32
	// Score is the score the client should use when displaying the progress.
	Score      int32
	ActivityID string
}

// ID ...
func (*LessonProgress) ID() uint32 {
	return IDLessonProgress
}

func (pk *LessonProgress) Marshal(io protocol.IO) {
	io.Varint32(&pk.LessonAction)
	io.Varint32(&pk.Score)
	io.String(&pk.ActivityID)
}
