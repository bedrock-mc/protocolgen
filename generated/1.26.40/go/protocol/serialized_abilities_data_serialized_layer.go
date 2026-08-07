// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SerializedAbilitiesDataSerializedLayer struct {
	SerializedLayer  uint16
	AbilitiesSet     uint32
	AbilityValues    uint32
	FlySpeed         float32
	VerticalFlySpeed float32
	WalkSpeed        float32
}

// Marshal reads or writes SerializedAbilitiesDataSerializedLayer using its canonical wire layout.
func (x *SerializedAbilitiesDataSerializedLayer) Marshal(io IO) {
	io.Uint16(&x.SerializedLayer)
	io.Uint32(&x.AbilitiesSet)
	io.Uint32(&x.AbilityValues)
	io.Float32(&x.FlySpeed)
	io.Float32(&x.VerticalFlySpeed)
	io.Float32(&x.WalkSpeed)
}
