// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MobBorn struct {
	BornBabyEntityType    int32
	BornBabyEntityVariant int32
	BornBabyColor         uint8
}

func (*MobBorn) isEvent() {}

// Marshal reads or writes MobBorn using its canonical wire layout.
func (x *MobBorn) Marshal(io IO) {
	io.Varint32(&x.BornBabyEntityType)
	io.Varint32(&x.BornBabyEntityVariant)
	io.Uint8(&x.BornBabyColor)
}
