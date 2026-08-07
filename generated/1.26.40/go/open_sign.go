// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type OpenSign struct {
	Pos         BlockPos
	IsFrontSide bool
}

// Marshal reads or writes OpenSign using its canonical wire layout.
func (x *OpenSign) Marshal(io IO) {
	x.Pos.Marshal(io)
	io.Bool(&x.IsFrontSide)
}
