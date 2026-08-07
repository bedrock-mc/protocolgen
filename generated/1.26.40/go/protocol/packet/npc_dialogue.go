// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type NpcDialogue struct {
	NpcIDRawID            uint64
	NpcDialogueActionType protocol.NpcDialogueActionType
	Dialogue              string
	SceneName             string
	NpcName               string
	ActionJSON            string
}

// Marshal reads or writes NpcDialogue using its canonical wire layout.
func (x *NpcDialogue) Marshal(io protocol.IO) {
	io.Uint64(&x.NpcIDRawID)
	protocol.IntegerFunc(&x.NpcDialogueActionType, io.Varint32)
	io.String(&x.Dialogue)
	io.String(&x.SceneName)
	io.String(&x.NpcName)
	io.String(&x.ActionJSON)
}

// ID returns the protocol ID for NpcDialogue.
func (*NpcDialogue) ID() uint32 { return IDNpcDialogue }
