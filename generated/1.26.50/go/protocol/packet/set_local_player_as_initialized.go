// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type SetLocalPlayerAsInitialized struct {
	PlayerID uint64
}

// ID ...
func (*SetLocalPlayerAsInitialized) ID() uint32 {
	return IDSetLocalPlayerAsInitialized
}

func (pk *SetLocalPlayerAsInitialized) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.PlayerID)
}
