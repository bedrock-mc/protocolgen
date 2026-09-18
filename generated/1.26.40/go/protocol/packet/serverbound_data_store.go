// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ServerboundDataStore struct {
	Update protocol.BedrockDDUIDataStoreUpdate
}

// ID returns the protocol ID for ServerboundDataStore.
func (*ServerboundDataStore) ID() uint32 { return IDServerboundDataStore }

// Marshal reads or writes ServerboundDataStore using its canonical wire layout.
func (pk *ServerboundDataStore) Marshal(io protocol.IO) {
	pk.Update.Marshal(io)
}
