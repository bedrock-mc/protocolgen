// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LessonProgress struct {
	LessonAction int32
	Score        int32
	ActivityId   string
}

// Marshal reads or writes LessonProgress using its canonical wire layout.
func (x *LessonProgress) Marshal(io IO) {
	io.Varint32(&x.LessonAction)
	io.Varint32(&x.Score)
	io.String(&x.ActivityId)
}
