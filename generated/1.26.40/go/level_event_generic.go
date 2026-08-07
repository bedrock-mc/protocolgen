// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LevelEventGeneric struct {
	EventId int32
	CTD     []byte
}

// Marshal reads or writes LevelEventGeneric using its canonical wire layout.
func (x *LevelEventGeneric) Marshal(io IO) {
	io.Varint32(&x.EventId)
	io.NBT(&x.CTD)
}
