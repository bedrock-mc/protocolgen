// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CodeBuilderScoreboard struct {
	ObjectiveName string
	Score         int32
}

func (*CodeBuilderScoreboard) isEvent() {}

// Marshal reads or writes CodeBuilderScoreboard using its canonical wire layout.
func (x *CodeBuilderScoreboard) Marshal(io IO) {
	io.String(&x.ObjectiveName)
	io.Varint32(&x.Score)
}
