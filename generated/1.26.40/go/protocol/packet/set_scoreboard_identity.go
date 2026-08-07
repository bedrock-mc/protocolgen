// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetScoreboardIdentity struct {
	ScoreboardIdentityPacketType protocol.ScoreboardIdentityPacketType
	ScoreboardIdentityInfo       []protocol.ScoreboardIdentityPacketInfo
}

// Marshal reads or writes SetScoreboardIdentity using its canonical wire layout.
func (x *SetScoreboardIdentity) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.ScoreboardIdentityPacketType, io.Uint8)
	protocol.FuncSlice(io, &x.ScoreboardIdentityInfo, io.Varuint32, func(value *protocol.ScoreboardIdentityPacketInfo) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for SetScoreboardIdentity.
func (*SetScoreboardIdentity) ID() uint32 { return IDSetScoreboardIdentity }
