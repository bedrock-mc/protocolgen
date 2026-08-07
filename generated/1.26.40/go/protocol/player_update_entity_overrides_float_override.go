// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerUpdateEntityOverridesFloatOverride struct {
	Type  string
	Value float32
}

func (PlayerUpdateEntityOverridesFloatOverride) isPlayerUpdateEntityOverridesUpdate() {}

// Marshal reads or writes PlayerUpdateEntityOverridesFloatOverride using its canonical wire layout.
func (x *PlayerUpdateEntityOverridesFloatOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Float32(&x.Value)
}
