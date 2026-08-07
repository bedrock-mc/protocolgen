// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ServerboundDataStore struct {
	Update protocol.BedrockDDUIDataStoreUpdate
}

// Marshal reads or writes ServerboundDataStore using its canonical wire layout.
func (x *ServerboundDataStore) Marshal(io protocol.IO) {
	x.Update.Marshal(io)
}
