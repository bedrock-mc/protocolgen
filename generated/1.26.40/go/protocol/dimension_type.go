// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DimensionType struct {
	Value int32
}

// Marshal reads or writes DimensionType using its canonical wire layout.
func (x *DimensionType) Marshal(io IO) {
	io.Varint32(&x.Value)
}
