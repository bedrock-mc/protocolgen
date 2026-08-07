// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type NpcRequest struct {
	NPCRuntimeID uint64
	RequestType  protocol.NpcRequestRequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}

// Marshal reads or writes NpcRequest using its canonical wire layout.
func (x *NpcRequest) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.NPCRuntimeID)
	protocol.IntegerFunc(&x.RequestType, io.Uint8)
	io.String(&x.Actions)
	io.Uint8(&x.ActionIndex)
	io.String(&x.SceneName)
}

// ID returns the protocol ID for NpcRequest.
func (*NpcRequest) ID() uint32 { return IDNpcRequest }
