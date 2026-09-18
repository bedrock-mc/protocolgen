// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type NpcRequest struct {
	NPCRuntimeID uint64
	RequestType  protocol.RequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}

// ID returns the protocol ID for NpcRequest.
func (*NpcRequest) ID() uint32 { return IDNpcRequest }

// Marshal reads or writes NpcRequest using its canonical wire layout.
func (pk *NpcRequest) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.NPCRuntimeID)
	pk.RequestType.Marshal(io)
	io.String(&pk.Actions)
	io.Uint8(&pk.ActionIndex)
	io.String(&pk.SceneName)
}
