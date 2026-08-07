// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateIdAndVersion  protocol.PackIdVersion
	ResourcePacks              []protocol.PackInfoData
}

// Marshal reads or writes ResourcePacksInfo using its canonical wire layout.
func (x *ResourcePacksInfo) Marshal(io protocol.IO) {
	io.Bool(&x.ResourcePackRequired)
	io.Bool(&x.HasAddonPacks)
	io.Bool(&x.HasScripts)
	io.Bool(&x.ForceDisableVibrantVisuals)
	x.WorldTemplateIdAndVersion.Marshal(io)
	protocol.FuncSlice(io, &x.ResourcePacks, io.Varuint32, func(value *protocol.PackInfoData) {
		value.Marshal(io)
	})
}
