// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LessonProgress struct {
	LessonAction int32
	Score        int32
	ActivityId   string
}

// Marshal reads or writes LessonProgress using its canonical wire layout.
func (x *LessonProgress) Marshal(io protocol.IO) {
	io.Varint32(&x.LessonAction)
	io.Varint32(&x.Score)
	io.String(&x.ActivityId)
}

// ID returns the protocol ID for LessonProgress.
func (*LessonProgress) ID() uint32 { return IDLessonProgress }
