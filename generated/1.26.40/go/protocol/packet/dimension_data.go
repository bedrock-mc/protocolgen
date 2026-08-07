// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type DimensionData struct {
	Definitions []protocol.OrderedEntry[string, protocol.DimensionDefinition]
}

// Marshal reads or writes DimensionData using its canonical wire layout.
func (x *DimensionData) Marshal(io protocol.IO) {
	protocol.OrderedMap(io, &x.Definitions, io.Varuint32, io.String, func(value *protocol.DimensionDefinition) {
		value.Marshal(io)
	})
}
