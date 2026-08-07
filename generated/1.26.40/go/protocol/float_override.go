// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type FloatOverride struct {
	Type  string
	Value float32
}

func (*FloatOverride) isPlayerUpdateEntityOverridesData() {}

// Marshal reads or writes FloatOverride using its canonical wire layout.
func (x *FloatOverride) Marshal(io IO) {
	io.String(&x.Type)
	io.Float32(&x.Value)
}
