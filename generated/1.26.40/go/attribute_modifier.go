// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AttributeModifier struct {
	Id             string
	Name           string
	Amount         float32
	Operation      int32
	Operand        int32
	IsSerializable bool
}

// Marshal reads or writes AttributeModifier using its canonical wire layout.
func (x *AttributeModifier) Marshal(io IO) {
	io.String(&x.Id)
	io.String(&x.Name)
	io.Float32(&x.Amount)
	io.Int32(&x.Operation)
	io.Int32(&x.Operand)
	io.Bool(&x.IsSerializable)
}
