// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type ServerboundPackSettingChange struct {
	PackId           uuid.UUID
	PackSettingName  string
	PackSettingValue ServerboundPackSettingChangePackSettingValue
}

// Marshal reads or writes ServerboundPackSettingChange using its canonical wire layout.
func (x *ServerboundPackSettingChange) Marshal(io IO) {
	io.UUID(&x.PackId)
	io.String(&x.PackSettingName)
	marshalServerboundPackSettingChangePackSettingValue(io, &x.PackSettingValue)
}
