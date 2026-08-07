// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type HiddenLocation struct {
	PacketType PlayerLocationType
}

func (*HiddenLocation) isPlayerLocationData() {}

// Marshal reads or writes HiddenLocation using its canonical wire layout.
func (x *HiddenLocation) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Varint32)
}
