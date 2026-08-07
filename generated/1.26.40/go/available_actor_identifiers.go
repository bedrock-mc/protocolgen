// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableActorIdentifiers struct {
	IdentifierList []byte
}

// Marshal reads or writes AvailableActorIdentifiers using its canonical wire layout.
func (x *AvailableActorIdentifiers) Marshal(io IO) {
	io.NBT(&x.IdentifierList)
}
