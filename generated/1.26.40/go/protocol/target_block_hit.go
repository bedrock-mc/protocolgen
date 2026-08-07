// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TargetBlockHit struct {
	RedstoneLevel int32
}

func (*TargetBlockHit) isEvent() {}

// Marshal reads or writes TargetBlockHit using its canonical wire layout.
func (x *TargetBlockHit) Marshal(io IO) {
	io.Varint32(&x.RedstoneLevel)
}
