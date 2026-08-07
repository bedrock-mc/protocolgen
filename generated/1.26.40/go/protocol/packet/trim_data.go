// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type TrimData struct {
	TrimPatternList  []protocol.TrimPattern
	TrimMaterialList []protocol.TrimMaterial
}

// Marshal reads or writes TrimData using its canonical wire layout.
func (x *TrimData) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.TrimPatternList)
	protocol.Slice(io, &x.TrimMaterialList)
}

// ID returns the protocol ID for TrimData.
func (*TrimData) ID() uint32 { return IDTrimData }
