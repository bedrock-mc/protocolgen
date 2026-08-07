// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate BlockPos
}

// Marshal reads or writes LecternUpdate using its canonical wire layout.
func (x *LecternUpdate) Marshal(io IO) {
	io.Uint8(&x.NewPageToShow)
	io.Uint8(&x.TotalPages)
	x.PositionOfLecternToUpdate.Marshal(io)
}
