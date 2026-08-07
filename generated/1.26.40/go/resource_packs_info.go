// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateIdAndVersion  PackIdVersion
	ResourcePacks              []PackInfoData
}

// Marshal reads or writes ResourcePacksInfo using its canonical wire layout.
func (x *ResourcePacksInfo) Marshal(io IO) {
	io.Bool(&x.ResourcePackRequired)
	io.Bool(&x.HasAddonPacks)
	io.Bool(&x.HasScripts)
	io.Bool(&x.ForceDisableVibrantVisuals)
	x.WorldTemplateIdAndVersion.Marshal(io)
	FuncSlice(io, &x.ResourcePacks, io.Varuint32, func(value *PackInfoData) {
		value.Marshal(io)
	})
}
