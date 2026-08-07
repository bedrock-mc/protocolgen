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
	enumValue1 := int32(x.NpcDialogueActionType)
	io.Varint32(&enumValue1)
	x.NpcDialogueActionType = NpcDialogueNpcDialogueActionType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.Dialogue)
	io.String(&x.SceneName)
	io.String(&x.NpcName)
	io.String(&x.ActionJSON)
}
