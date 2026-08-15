// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

type ServerboundPackSettingChange struct {
	PackID           uuid.UUID
	PackSettingName  string
	PackSettingValue protocol.ServerboundPackSettingChangePackSettingValue
}

// Marshal reads or writes ServerboundPackSettingChange using its canonical wire layout.
func (x *ServerboundPackSettingChange) Marshal(io protocol.IO) {
	io.UUID(&x.PackID)
	io.StringLimits(&x.PackSettingName, 0, 128)
	protocol.MarshalServerboundPackSettingChangePackSettingValue(io, &x.PackSettingValue)
}

// ID returns the protocol ID for ServerboundPackSettingChange.
func (*ServerboundPackSettingChange) ID() uint32 { return IDServerboundPackSettingChange }
