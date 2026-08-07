// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ResourcePackStack struct {
	TexturePackRequired bool
	TexturePackList     []protocol.PackInstanceID
	BaseGameVersion     string
	Experiments         protocol.Experiments
	IncludeEditorPacks  bool
}

// Marshal reads or writes ResourcePackStack using its canonical wire layout.
func (x *ResourcePackStack) Marshal(io protocol.IO) {
	io.Bool(&x.TexturePackRequired)
	protocol.Slice(io, &x.TexturePackList)
	io.String(&x.BaseGameVersion)
	x.Experiments.Marshal(io)
	io.Bool(&x.IncludeEditorPacks)
}

// ID returns the protocol ID for ResourcePackStack.
func (*ResourcePackStack) ID() uint32 { return IDResourcePackStack }
