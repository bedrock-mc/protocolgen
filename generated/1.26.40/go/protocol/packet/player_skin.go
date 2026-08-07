// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

type PlayerSkin struct {
	UUID                 uuid.UUID
	SerializedSkin       protocol.SerializedSkinRef
	LocalizedNewSkinName string
	LocalizedOldSkinName string
}

// Marshal reads or writes PlayerSkin using its canonical wire layout.
func (x *PlayerSkin) Marshal(io protocol.IO) {
	io.UUID(&x.UUID)
	x.SerializedSkin.Marshal(io)
	io.String(&x.LocalizedNewSkinName)
	io.String(&x.LocalizedOldSkinName)
}

// ID returns the protocol ID for PlayerSkin.
func (*PlayerSkin) ID() uint32 { return IDPlayerSkin }
