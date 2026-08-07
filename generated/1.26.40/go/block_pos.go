// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BlockPos struct {
	X int32
	Y int32
	Z int32
}

// Marshal reads or writes BlockPos using its canonical wire layout.
func (x *BlockPos) Marshal(io IO) {
	io.Varint32(&x.X)
	io.Varint32(&x.Y)
	io.Varint32(&x.Z)
}
