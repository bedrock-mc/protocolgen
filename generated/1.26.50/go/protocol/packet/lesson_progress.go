// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// LessonProgress is a packet sent by the server to the client to inform the client of updated
// progress on a lesson. This packet only functions on the Minecraft: Education Edition version of
// the game.
type LessonProgress struct {
	LessonAction int32
	// Score is the score the client should use when displaying the progress.
	Score      int32
	ActivityID string
}

// Marshal reads or writes LessonProgress using its canonical wire layout.
func (x *LessonProgress) Marshal(io protocol.IO) {
	io.Varint32(&x.LessonAction)
	io.Varint32(&x.Score)
	io.String(&x.ActivityID)
}

// ID returns the protocol ID for LessonProgress.
func (*LessonProgress) ID() uint32 { return IDLessonProgress }
