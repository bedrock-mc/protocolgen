// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetActorData struct {
	TargetRuntimeID   uint64
	ActorData         protocol.SynchedActorDataCopyableDataList
	SynchedProperties protocol.PropertySyncData
	Tick              uint64
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
