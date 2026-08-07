// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistActorPriority struct {
	CameraAimAssistActorPriorityList []CameraAimAssistActorPriorityPriorityData
}

// Marshal reads or writes CameraAimAssistActorPriority using its canonical wire layout.
func (x *CameraAimAssistActorPriority) Marshal(io IO) {
	FuncSlice(io, &x.CameraAimAssistActorPriorityList, io.Varuint32, func(value *CameraAimAssistActorPriorityPriorityData) {
		value.Marshal(io)
	})
}
