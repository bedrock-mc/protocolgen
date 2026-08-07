// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type EditorNetwork struct {
	RouteToManager bool
	RawVariantName string
	RawVariantData string
}

// Marshal reads or writes EditorNetwork using its canonical wire layout.
func (x *EditorNetwork) Marshal(io protocol.IO) {
	io.Bool(&x.RouteToManager)
	io.String(&x.RawVariantName)
	io.String(&x.RawVariantData)
}
