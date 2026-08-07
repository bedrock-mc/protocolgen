// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerAction struct {
	PlayerRuntimeID ActorRuntimeID
	Action          PlayerActionType
	BlockPosition   BlockPos
	ResultPos       BlockPos
	Face            int32
}
