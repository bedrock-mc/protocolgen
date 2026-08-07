// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RemoveOverride struct {
	Type string
}

func (*RemoveOverride) isPlayerUpdateEntityOverridesData() {}

// Marshal reads or writes RemoveOverride using its canonical wire layout.
func (x *RemoveOverride) Marshal(io IO) {
	io.String(&x.Type)
}
