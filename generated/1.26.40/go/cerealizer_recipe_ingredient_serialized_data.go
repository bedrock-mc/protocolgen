// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CerealizerRecipeIngredientSerializedData struct {
	Descriptor []OrderedEntry[string, string]
	AuxValue   int32
	StackSize  int32
}

// Marshal reads or writes CerealizerRecipeIngredientSerializedData using its canonical wire layout.
func (x *CerealizerRecipeIngredientSerializedData) Marshal(io IO) {
	OrderedMap(io, &x.Descriptor, io.Varuint32, io.String, io.String)
	io.Varint32(&x.AuxValue)
	io.Varint32(&x.StackSize)
}
