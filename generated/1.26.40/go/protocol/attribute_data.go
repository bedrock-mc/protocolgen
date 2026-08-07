// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AttributeData struct {
	MinValue        float32
	MaxValue        float32
	CurrentValue    float32
	DefaultMinValue float32
	DefaultMaxValue float32
	DefaultValue    float32
	Name            string
	Modifiers       []AttributeModifier
}

// Marshal reads or writes AttributeData using its canonical wire layout.
func (x *AttributeData) Marshal(io IO) {
	io.Float32(&x.MinValue)
	io.Float32(&x.MaxValue)
	io.Float32(&x.CurrentValue)
	io.Float32(&x.DefaultMinValue)
	io.Float32(&x.DefaultMaxValue)
	io.Float32(&x.DefaultValue)
	io.String(&x.Name)
	Slice(io, &x.Modifiers)
}
