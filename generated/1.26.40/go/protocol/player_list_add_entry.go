// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import (
	"image/color"

	"github.com/google/uuid"
)

type PlayerListAddEntry struct {
	UUID             uuid.UUID
	ActorUniqueID    int64
	PlayerName       string
	XBLXUID          string
	PlatformOnlineID string
	BuildPlatform    BuildPlatform
	SerializedSkin   SerializedSkinRef
	IsTeacher        bool
	IsHost           bool
	IsSubClient      bool
	PlayerColor      color.RGBA
}

func (*PlayerListAddEntry) isPlayerListEntriesItem() {}

// Marshal reads or writes PlayerListAddEntry using its canonical wire layout.
func (x *PlayerListAddEntry) Marshal(io IO) {
	io.UUID(&x.UUID)
	io.ActorUniqueID(&x.ActorUniqueID)
	io.String(&x.PlayerName)
	io.String(&x.XBLXUID)
	io.String(&x.PlatformOnlineID)
	IntegerFunc(&x.BuildPlatform, io.Int32)
	x.SerializedSkin.Marshal(io)
	io.Bool(&x.IsTeacher)
	io.Bool(&x.IsHost)
	io.Bool(&x.IsSubClient)
	io.RGBA(&x.PlayerColor)
}
