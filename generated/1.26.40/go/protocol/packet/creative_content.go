// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CreativeContent struct {
	Groups  []protocol.CreativeGroupInfo
	Entries []protocol.CreativeItemEntry
}

// Marshal reads or writes CreativeContent using its canonical wire layout.
func (x *CreativeContent) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.Groups, io.Varuint32, func(value *protocol.CreativeGroupInfo) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.Entries, io.Varuint32, func(value *protocol.CreativeItemEntry) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for CreativeContent.
func (*CreativeContent) ID() uint32 { return IDCreativeContent }
