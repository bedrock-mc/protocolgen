// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type DeathInfo struct {
	DeathCauseAttackName  string
	DeathCauseMessageList []string
}

// Marshal reads or writes DeathInfo using its canonical wire layout.
func (x *DeathInfo) Marshal(io protocol.IO) {
	io.String(&x.DeathCauseAttackName)
	protocol.FuncSlice(io, &x.DeathCauseMessageList, io.Varuint32, io.String)
}
