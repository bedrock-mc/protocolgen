// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundDataStore struct {
	Updates []protocol.BedrockDDUI
}

// Marshal reads or writes ClientboundDataStore using its canonical wire layout.
func (x *ClientboundDataStore) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.Updates, io.Varuint32, func(value *protocol.BedrockDDUI) {
		protocol.MarshalBedrockDDUI(io, value)
	})
}
