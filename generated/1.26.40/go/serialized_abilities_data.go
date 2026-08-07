// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SerializedAbilitiesData struct {
	TargetPlayerRawId  int64
	PlayerPermissions  PlayerPermissionLevel
	CommandPermissions CommandPermissionLevel
	Layers             []SerializedAbilitiesDataSerializedLayer
}

// Marshal reads or writes SerializedAbilitiesData using its canonical wire layout.
func (x *SerializedAbilitiesData) Marshal(io IO) {
	io.Int64(&x.TargetPlayerRawId)
	IntegerFunc(&x.PlayerPermissions, io.Int8)
	IntegerFunc(&x.CommandPermissions, io.Uint8)
	FuncSlice(io, &x.Layers, io.Varuint32, func(value *SerializedAbilitiesDataSerializedLayer) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
