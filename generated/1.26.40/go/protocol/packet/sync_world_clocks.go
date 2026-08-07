// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SyncWorldClocks struct {
	Data protocol.SyncWorldClocksData
}

// Marshal reads or writes SyncWorldClocks using its canonical wire layout.
func (x *SyncWorldClocks) Marshal(io protocol.IO) {
	protocol.MarshalSyncWorldClocksData(io, &x.Data)
}

// ID returns the protocol ID for SyncWorldClocks.
func (*SyncWorldClocks) ID() uint32 { return IDSyncWorldClocks }
