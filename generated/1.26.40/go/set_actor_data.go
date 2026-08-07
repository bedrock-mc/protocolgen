// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetActorData struct {
	TargetRuntimeID   ActorRuntimeID
	ActorData         SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	Tick              PlayerInputTick
}

// Marshal reads or writes SetActorData using its canonical wire layout.
func (x *SetActorData) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	x.Tick.Marshal(io)
}
