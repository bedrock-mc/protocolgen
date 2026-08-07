// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerUpdateEntityOverridesClearOverride struct {
	Type string
}

func (PlayerUpdateEntityOverridesClearOverride) isPlayerUpdateEntityOverridesUpdate() {}

// Marshal reads or writes PlayerUpdateEntityOverridesClearOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesClearOverride) Marshal(io IO) {
	io.String(&x.Type)
}
