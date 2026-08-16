// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// SyncWorldClocks is sent by the server to initialise and synchronise world clocks with the client.
type SyncWorldClocks struct {
	Data protocol.SyncWorldClocksData
}

// Marshal reads or writes SyncWorldClocks using its canonical wire layout.
func (x *SyncWorldClocks) Marshal(io protocol.IO) {
	protocol.MarshalSyncWorldClocksData(io, &x.Data)
}

// ID returns the protocol ID for SyncWorldClocks.
func (*SyncWorldClocks) ID() uint32 { return IDSyncWorldClocks }
