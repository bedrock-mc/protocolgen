// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SetActorData is sent by the server to update the entity metadata of an entity. It includes flags
// such as if the entity is on fire, but also properties such as the air it has left until it starts
// drowning.
type SetActorData struct {
	TargetRuntimeID   uint64
	ActorData         protocol.SynchedActorDataCopyableDataList
	SynchedProperties protocol.PropertySyncData
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// Marshal reads or writes SetActorData using its canonical wire layout.
func (x *SetActorData) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for SetActorData.
func (*SetActorData) ID() uint32 { return IDSetActorData }
