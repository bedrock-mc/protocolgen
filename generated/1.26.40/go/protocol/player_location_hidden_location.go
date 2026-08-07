// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerLocationHiddenLocation struct {
	PacketType PlayerLocationType
}

func (*PlayerLocationHiddenLocation) isPlayerLocationLocation() {}

// Marshal reads or writes PlayerLocationHiddenLocation using its canonical wire layout.
func (x *PlayerLocationHiddenLocation) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Varint32)
}
