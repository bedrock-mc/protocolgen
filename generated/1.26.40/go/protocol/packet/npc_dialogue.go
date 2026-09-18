// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type NpcDialogue struct {
	NpcIDRawID            uint64
	NpcDialogueActionType protocol.NpcDialogueActionType
	Dialogue              string
	SceneName             string
	NpcName               string
	ActionJSON            string
}

// ID returns the protocol ID for NpcDialogue.
func (*NpcDialogue) ID() uint32 { return IDNpcDialogue }

// Marshal reads or writes NpcDialogue using its canonical wire layout.
func (pk *NpcDialogue) Marshal(io protocol.IO) {
	io.Uint64(&pk.NpcIDRawID)
	pk.NpcDialogueActionType.Marshal(io)
	io.String(&pk.Dialogue)
	io.String(&pk.SceneName)
	io.String(&pk.NpcName)
	io.String(&pk.ActionJSON)
}
