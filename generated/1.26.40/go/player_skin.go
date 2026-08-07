// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type PlayerSkin struct {
	UUID                 uuid.UUID
	SerializedSkin       SerializedSkinRef
	LocalizedNewSkinName string
	LocalizedOldSkinName string
}
