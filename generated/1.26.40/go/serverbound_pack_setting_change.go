// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type ServerboundPackSettingChange struct {
	PackId           uuid.UUID
	PackSettingName  string
	PackSettingValue ServerboundPackSettingChangePackSettingValue
}
