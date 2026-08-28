// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// ResourcePackStack is sent by the server to send the order in which resource packs and behaviour
// packs should be applied (and downloaded) by the client.
type ResourcePackStack struct {
	// TexturePackRequired specifies if the client must accept the texture packs the server has in order
	// to join the server. If set to true, the client gets the option to either download the resource
	// packs and join, or quit entirely. Behaviour packs never have to be downloaded.
	TexturePackRequired bool
	TexturePackList     []protocol.PackInstanceID
	// BaseGameVersion is the vanilla version that the client should set its resource pack stack to.
	BaseGameVersion string
	// Experiments holds a list of experiments that are either enabled or disabled in the world that the
	// player spawns in. It is not clear why experiments are sent both here and in the StartGame packet.
	Experiments protocol.Experiments
	// IncludeEditorPacks specifies if vanilla editor packs should be included in the resource pack
	// stack when connecting to an editor world.
	IncludeEditorPacks bool
}

// Marshal reads or writes ResourcePackStack using its canonical wire layout.
func (x *ResourcePackStack) Marshal(io protocol.IO) {
	io.Bool(&x.TexturePackRequired)
	protocol.SliceLimits(io, &x.TexturePackList, 0, 65535)
	io.String(&x.BaseGameVersion)
	x.Experiments.Marshal(io)
	io.Bool(&x.IncludeEditorPacks)
}

// ID returns the protocol ID for ResourcePackStack.
func (*ResourcePackStack) ID() uint32 { return IDResourcePackStack }
