// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetScore struct {
	ScoreInfo []protocol.SetScoreInfoItem
}

// Marshal reads or writes SetScore using its canonical wire layout.
func (x *SetScore) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.ScoreInfo, io.Varuint32, func(value *protocol.SetScoreInfoItem) {
		protocol.MarshalSetScoreInfoItem(io, value)
	})
}

// ID returns the protocol ID for SetScore.
func (*SetScore) ID() uint32 { return IDSetScore }
