// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ClientboundDataStore struct {
	Updates []protocol.BedrockDDUI
}

// ID ...
func (*ClientboundDataStore) ID() uint32 {
	return IDClientboundDataStore
}

func (pk *ClientboundDataStore) Marshal(io protocol.IO) {
	protocol.FuncSliceLimits(io, &pk.Updates, io.Varuint32, 0, 500, func(value *protocol.BedrockDDUI) {
		protocol.MarshalBedrockDDUI(io, value)
	})
}
