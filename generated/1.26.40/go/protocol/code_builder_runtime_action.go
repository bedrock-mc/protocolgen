// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CodeBuilderRuntimeAction struct {
	CodeBuilderRuntimeAction string
}

func (*CodeBuilderRuntimeAction) isEventData() {}

// Marshal reads or writes CodeBuilderRuntimeAction using its canonical wire layout.
func (x *CodeBuilderRuntimeAction) Marshal(io IO) {
	io.String(&x.CodeBuilderRuntimeAction)
}
