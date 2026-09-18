// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ClientboundDataStore is sent by the server to update, change or remove data store entries on the client.
type ClientboundDataStore struct {
	// Updates is an array of data store changes. Each entry has its own change type discriminator.
	Updates []protocol.BedrockDDUI
}

// ID returns the protocol ID for ClientboundDataStore.
func (*ClientboundDataStore) ID() uint32 { return IDClientboundDataStore }

// Marshal reads or writes ClientboundDataStore using its canonical wire layout.
func (pk *ClientboundDataStore) Marshal(io protocol.IO) {
	protocol.FuncSliceLimits(io, &pk.Updates, io.Varuint32, 0, 500, func(value *protocol.BedrockDDUI) {
		protocol.MarshalBedrockDDUI(io, value)
	})
}
