// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerList struct {
	Entries []protocol.PlayerListEntriesItem
}

// Marshal reads or writes PlayerList using its canonical wire layout.
func (x *PlayerList) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.Entries, io.Varuint32, func(value *protocol.PlayerListEntriesItem) {
		protocol.MarshalPlayerListEntriesItem(io, value)
	})
}

// ID returns the protocol ID for PlayerList.
func (*PlayerList) ID() uint32 { return IDPlayerList }
