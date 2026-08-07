// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SpawnSettings struct {
	SpawnBiomeType       SpawnBiomeType
	UserDefinedBiomeName string
	Dimension            int32
}

// Marshal reads or writes SpawnSettings using its canonical wire layout.
func (x *SpawnSettings) Marshal(io IO) {
	IntegerFunc(&x.SpawnBiomeType, io.Int16)
	io.String(&x.UserDefinedBiomeName)
	io.Varint32(&x.Dimension)
}
