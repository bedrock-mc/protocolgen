// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CameraAimAssistActorPriority struct {
	CameraAimAssistActorPriorityList []protocol.CameraAimAssistActorPriorityPriorityData
}

// Marshal reads or writes CameraAimAssistActorPriority using its canonical wire layout.
func (x *CameraAimAssistActorPriority) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.CameraAimAssistActorPriorityList, io.Varuint32, func(value *protocol.CameraAimAssistActorPriorityPriorityData) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for CameraAimAssistActorPriority.
func (*CameraAimAssistActorPriority) ID() uint32 { return IDCameraAimAssistActorPriority }
