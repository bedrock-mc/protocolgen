// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NpcDialogue struct {
	NpcIdRawId            uint64
	NpcDialogueActionType NpcDialogueNpcDialogueActionType
	Dialogue              string
	SceneName             string
	NpcName               string
	ActionJSON            string
}

// Marshal reads or writes NpcDialogue using its canonical wire layout.
func (x *NpcDialogue) Marshal(io IO) {
	io.Uint64(&x.NpcIdRawId)
	IntegerFunc(&x.NpcDialogueActionType, io.Varint32)
	io.String(&x.Dialogue)
	io.String(&x.SceneName)
	io.String(&x.NpcName)
	io.String(&x.ActionJSON)
}
