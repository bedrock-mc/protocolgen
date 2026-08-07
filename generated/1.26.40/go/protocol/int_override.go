// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type IntOverride struct {
	Type  string
	Value int32
}

func (*IntOverride) isPlayerUpdateEntityOverridesData() {}

// Marshal reads or writes IntOverride using its canonical wire layout.
func (x *IntOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Int32(&x.Value)
}
