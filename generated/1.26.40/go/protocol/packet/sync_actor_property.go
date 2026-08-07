// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SyncActorProperty struct {
	PropertyData []byte
}

// Marshal reads or writes SyncActorProperty using its canonical wire layout.
func (x *SyncActorProperty) Marshal(io protocol.IO) {
	io.NBT(&x.PropertyData)
}

// ID returns the protocol ID for SyncActorProperty.
func (*SyncActorProperty) ID() uint32 { return IDSyncActorProperty }
