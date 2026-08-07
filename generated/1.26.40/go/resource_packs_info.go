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
	if !io.Reading() && uint64(len(x.ResourcePacks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ResourcePacks), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ResourcePacks))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ResourcePacks = make([]PackInfoData, int(count1))
	}
	for index2 := range x.ResourcePacks {
		x.ResourcePacks[index2].Marshal(io)
	}
}
