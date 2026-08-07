// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistActorPriority struct {
	CameraAimAssistActorPriorityList []CameraAimAssistActorPriorityPriorityData
}

// Marshal reads or writes CameraAimAssistActorPriority using its canonical wire layout.
func (x *CameraAimAssistActorPriority) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.CameraAimAssistActorPriorityList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CameraAimAssistActorPriorityList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.CameraAimAssistActorPriorityList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.CameraAimAssistActorPriorityList = make([]CameraAimAssistActorPriorityPriorityData, int(count1))
	}
	for index2 := range x.CameraAimAssistActorPriorityList {
		x.CameraAimAssistActorPriorityList[index2].Marshal(io)
	}
}
