// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncActorProperty struct {
	PropertyData []byte
}

// Marshal reads or writes SyncActorProperty using its canonical wire layout.
func (x *SyncActorProperty) Marshal(io IO) {
	io.NBT(&x.PropertyData)
}
