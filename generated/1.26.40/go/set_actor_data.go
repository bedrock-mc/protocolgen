// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetActorData struct {
	TargetRuntimeID   ActorRuntimeID
	ActorData         SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	Tick              PlayerInputTick
}
