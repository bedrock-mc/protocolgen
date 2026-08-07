// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateClientInputLocks struct {
	InputLockComponentData uint32
}

// Marshal reads or writes UpdateClientInputLocks using its canonical wire layout.
func (x *UpdateClientInputLocks) Marshal(io protocol.IO) {
	io.Varuint32(&x.InputLockComponentData)
}
