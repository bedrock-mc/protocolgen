// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ClearOverride struct {
	Type string
}

func (*ClearOverride) isPlayerUpdateEntityOverridesData() {}

// Marshal reads or writes ClearOverride using its canonical wire layout.
func (x *ClearOverride) Marshal(io IO) {
	io.String(&x.Type)
}
