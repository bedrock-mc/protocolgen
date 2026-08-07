// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MapCreateLockedCopy struct {
	OriginalMapId int64
	NewMapId      int64
}

// Marshal reads or writes MapCreateLockedCopy using its canonical wire layout.
func (x *MapCreateLockedCopy) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.OriginalMapId)
	io.ActorUniqueID(&x.NewMapId)
}
