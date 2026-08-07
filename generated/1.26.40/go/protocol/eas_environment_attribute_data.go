// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EASEnvironmentAttributeData struct {
	AttributeName          string
	FromAttribute          Optional[EAS]
	Attribute              EAS
	ToAttribute            Optional[EAS]
	CurrentTransitionTicks uint32
	TotalTransitionTicks   uint32
	Easing                 string
	LocalTransitionTicks   uint32
	NoiseTransition        bool
}

// Marshal reads or writes EASEnvironmentAttributeData using its canonical wire layout.
func (x *EASEnvironmentAttributeData) Marshal(io IO) {
	io.String(&x.AttributeName)
	OptionalFunc(io, &x.FromAttribute, func(value *EAS) {
		MarshalEAS(io, value)
	})
	MarshalEAS(io, &x.Attribute)
	OptionalFunc(io, &x.ToAttribute, func(value *EAS) {
		MarshalEAS(io, value)
	})
	io.Uint32(&x.CurrentTransitionTicks)
	io.Uint32(&x.TotalTransitionTicks)
	io.String(&x.Easing)
	io.Uint32(&x.LocalTransitionTicks)
	io.Bool(&x.NoiseTransition)
}
