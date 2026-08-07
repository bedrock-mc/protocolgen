// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateAttributes struct {
	TargetRuntimeID uint64
	AttributeList   []AttributeData
	Tick            uint64
}

// Marshal reads or writes UpdateAttributes using its canonical wire layout.
func (x *UpdateAttributes) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	FuncSlice(io, &x.AttributeList, io.Varuint32, func(value *AttributeData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.PlayerInputTick(&x.Tick)
}
