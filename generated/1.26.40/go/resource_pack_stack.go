// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePackStack struct {
	TexturePackRequired bool
	TexturePackList     []PackInstanceId
	BaseGameVersion     string
	Experiments         Experiments
	IncludeEditorPacks  bool
}

// Marshal reads or writes ResourcePackStack using its canonical wire layout.
func (x *ResourcePackStack) Marshal(io IO) {
	io.Bool(&x.TexturePackRequired)
	if !io.Reading() && uint64(len(x.TexturePackList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.TexturePackList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.TexturePackList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.TexturePackList = make([]PackInstanceId, int(count1))
	}
	for index2 := range x.TexturePackList {
		x.TexturePackList[index2].Marshal(io)
	}
	io.String(&x.BaseGameVersion)
	x.Experiments.Marshal(io)
	io.Bool(&x.IncludeEditorPacks)
}
