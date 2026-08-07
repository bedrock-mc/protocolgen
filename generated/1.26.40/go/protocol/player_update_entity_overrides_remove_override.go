// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerUpdateEntityOverridesRemoveOverride struct {
	Type string
}

func (*PlayerUpdateEntityOverridesRemoveOverride) isPlayerUpdateEntityOverridesUpdate() {}

// Marshal reads or writes PlayerUpdateEntityOverridesRemoveOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesRemoveOverride) Marshal(io IO) {
	io.String(&x.Type)
}
