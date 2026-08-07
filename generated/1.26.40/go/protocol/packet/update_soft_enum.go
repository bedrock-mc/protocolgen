// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateSoftEnum struct {
	EnumName   string
	Values     []string
	UpdateType protocol.SoftEnumUpdateType
}

// Marshal reads or writes UpdateSoftEnum using its canonical wire layout.
func (x *UpdateSoftEnum) Marshal(io protocol.IO) {
	io.String(&x.EnumName)
	protocol.FuncSlice(io, &x.Values, io.Varuint32, io.String)
	protocol.IntegerFunc(&x.UpdateType, io.Uint8)
}

// ID returns the protocol ID for UpdateSoftEnum.
func (*UpdateSoftEnum) ID() uint32 { return IDUpdateSoftEnum }
