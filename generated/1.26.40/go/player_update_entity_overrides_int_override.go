// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerUpdateEntityOverridesIntOverride struct {
	Type  string
	Value int32
}

func (PlayerUpdateEntityOverridesIntOverride) isPlayerUpdateEntityOverridesUpdate() {}

// Marshal reads or writes PlayerUpdateEntityOverridesIntOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesIntOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Int32(&x.Value)
}
