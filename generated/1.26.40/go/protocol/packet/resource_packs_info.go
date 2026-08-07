// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateIDAndVersion  protocol.PackIDVersion
	ResourcePacks              []protocol.PackInfoData
}

// Marshal reads or writes ResourcePacksInfo using its canonical wire layout.
func (x *ResourcePacksInfo) Marshal(io protocol.IO) {
	io.Bool(&x.ResourcePackRequired)
	io.Bool(&x.HasAddonPacks)
	io.Bool(&x.HasScripts)
	io.Bool(&x.ForceDisableVibrantVisuals)
	x.WorldTemplateIDAndVersion.Marshal(io)
	protocol.Slice(io, &x.ResourcePacks)
}

// ID returns the protocol ID for ResourcePacksInfo.
func (*ResourcePacksInfo) ID() uint32 { return IDResourcePacksInfo }
