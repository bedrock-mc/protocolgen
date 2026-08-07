// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateClientInputLocks struct {
	InputLockComponentData uint32
}

// Marshal reads or writes UpdateClientInputLocks using its canonical wire layout.
func (x *UpdateClientInputLocks) Marshal(io IO) {
	io.Varuint32(&x.InputLockComponentData)
}
