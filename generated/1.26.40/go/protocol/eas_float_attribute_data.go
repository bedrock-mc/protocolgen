// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EASFloatAttributeData struct {
	Value         float32
	Operation     string
	ConstraintMin Optional[float32]
	ConstraintMax Optional[float32]
}

func (*EASFloatAttributeData) isEAS() {}

// Marshal reads or writes EASFloatAttributeData using its canonical wire layout.
func (x *EASFloatAttributeData) Marshal(io IO) {
	io.Float32(&x.Value)
	io.String(&x.Operation)
	OptionalFunc(io, &x.ConstraintMin, io.Float32)
	OptionalFunc(io, &x.ConstraintMax, io.Float32)
}
