// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateAttributes struct {
	TargetRuntimeID ActorRuntimeID
	AttributeList   []AttributeData
	Tick            PlayerInputTick
}

// Marshal reads or writes UpdateAttributes using its canonical wire layout.
func (x *UpdateAttributes) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	if !io.Reading() && uint64(len(x.AttributeList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.AttributeList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.AttributeList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.AttributeList = make([]AttributeData, int(count1))
	}
	for index2 := range x.AttributeList {
		x.AttributeList[index2].Marshal(io)
	}
	x.Tick.Marshal(io)
}
