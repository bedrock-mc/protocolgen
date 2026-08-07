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
	FuncSlice(io, &x.TexturePackList, io.Varuint32, func(value *PackInstanceId) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.String(&x.BaseGameVersion)
	x.Experiments.Marshal(io)
	io.Bool(&x.IncludeEditorPacks)
}
