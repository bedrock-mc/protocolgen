// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BossEventUpdateType uint8

const (
	BossEventUpdateTypeAdd              BossEventUpdateType = 0
	BossEventUpdateTypePlayerAdded      BossEventUpdateType = 1
	BossEventUpdateTypeRemove           BossEventUpdateType = 2
	BossEventUpdateTypePlayerRemoved    BossEventUpdateType = 3
	BossEventUpdateTypeUpdatePercent    BossEventUpdateType = 4
	BossEventUpdateTypeUpdateName       BossEventUpdateType = 5
	BossEventUpdateTypeUpdateProperties BossEventUpdateType = 6
	BossEventUpdateTypeUpdateStyle      BossEventUpdateType = 7
	BossEventUpdateTypeQuery            BossEventUpdateType = 8
)
