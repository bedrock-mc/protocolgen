// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SphereData struct {
	NumSegments uint8
}

func (SphereData) isPrimitiveShapeDataExtraShapeData() {}

// Marshal reads or writes SphereData using its canonical wire layout.
func (x *SphereData) Marshal(io IO) {
	io.Uint8(&x.NumSegments)
}
