// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ResourcePacksInfo is sent by the server to inform the client on what resource packs the server has. It
// sends a list of the resource packs it has and basic information on them like the version and description.
type ResourcePacksInfo struct {
	// TexturePackRequired specifies if the client must accept the texture packs the server has in order to join
	// the server. If set to true, the client gets the option to either download the resource packs and join, or
	// quit entirely. Behaviour packs never have to be downloaded.
	ResourcePackRequired bool
	// HasAddons specifies if any of the resource packs contain addons in them. If set to true, only clients that
	// support addons will be able to download them.
	HasAddonPacks bool
	// HasScripts specifies if any of the resource packs contain scripts in them. If set to true, only clients
	// that support scripts will be able to download them.
	HasScripts bool
	// ForceDisableVibrantVisuals specifies if the vibrant visuals feature should be forcibly disabled on the
	// server. If set to true, the server will ensure that vibrant visuals are not enabled, regardless of the
	// client's settings.
	ForceDisableVibrantVisuals bool
	// WorldTemplateUUID is the UUID of the template that has been used to generate the world. Templates can be
	// downloaded from the marketplace or installed via '.mctemplate' files. If the world was not generated from a
	// template, this field is empty.
	WorldTemplateIDAndVersion protocol.PackIDVersion
	ResourcePacks             []protocol.PackInfoData
}

// ID ...
func (*ResourcePacksInfo) ID() uint32 {
	return IDResourcePacksInfo
}

func (pk *ResourcePacksInfo) Marshal(io protocol.IO) {
	io.Bool(&pk.ResourcePackRequired)
	io.Bool(&pk.HasAddonPacks)
	io.Bool(&pk.HasScripts)
	io.Bool(&pk.ForceDisableVibrantVisuals)
	pk.WorldTemplateIDAndVersion.Marshal(io)
	protocol.SliceLimits(io, &pk.ResourcePacks, 0, 65535)
}
