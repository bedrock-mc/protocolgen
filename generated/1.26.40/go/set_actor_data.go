// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetActorData struct {
	TargetRuntimeID   uint64
	ActorData         SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	Tick              uint64
}

// Marshal reads or writes SetActorData using its canonical wire layout.
func (x *SetActorData) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	io.PlayerInputTick(&x.Tick)
}
