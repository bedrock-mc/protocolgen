// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraAimAssistActorPriorityPriorityData struct {
	PresetIndex   int32
	CategoryIndex int32
	ActorIndex    int32
	PriorityValue int32
}

// Marshal reads or writes CameraAimAssistActorPriorityPriorityData using its canonical wire layout.
func (x *CameraAimAssistActorPriorityPriorityData) Marshal(io IO) {
	io.Int32(&x.PresetIndex)
	io.Int32(&x.CategoryIndex)
	io.Int32(&x.ActorIndex)
	io.Int32(&x.PriorityValue)
}
