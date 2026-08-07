// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NpcRequest struct {
	NPCRuntimeID ActorRuntimeID
	RequestType  NpcRequestRequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}

// Marshal reads or writes NpcRequest using its canonical wire layout.
func (x *NpcRequest) Marshal(io IO) {
	x.NPCRuntimeID.Marshal(io)
	enumValue1 := uint8(x.RequestType)
	io.Uint8(&enumValue1)
	x.RequestType = NpcRequestRequestType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.Actions)
	io.Uint8(&x.ActionIndex)
	io.String(&x.SceneName)
}
