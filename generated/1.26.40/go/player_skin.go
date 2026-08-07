// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type PlayerSkin struct {
	UUID                 uuid.UUID
	SerializedSkin       SerializedSkinRef
	LocalizedNewSkinName string
	LocalizedOldSkinName string
}

// Marshal reads or writes PlayerSkin using its canonical wire layout.
func (x *PlayerSkin) Marshal(io IO) {
	io.UUID(&x.UUID)
	x.SerializedSkin.Marshal(io)
	io.String(&x.LocalizedNewSkinName)
	io.String(&x.LocalizedOldSkinName)
}
