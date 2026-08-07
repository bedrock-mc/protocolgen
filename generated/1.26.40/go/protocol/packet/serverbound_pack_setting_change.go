// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

type ServerboundPackSettingChange struct {
	PackId           uuid.UUID
	PackSettingName  string
	PackSettingValue protocol.ServerboundPackSettingChangePackSettingValue
}

// Marshal reads or writes ServerboundPackSettingChange using its canonical wire layout.
func (x *ServerboundPackSettingChange) Marshal(io protocol.IO) {
	io.UUID(&x.PackId)
	io.String(&x.PackSettingName)
	protocol.MarshalServerboundPackSettingChangePackSettingValue(io, &x.PackSettingValue)
}
