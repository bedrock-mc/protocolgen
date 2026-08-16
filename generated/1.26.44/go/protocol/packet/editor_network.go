// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// EditorNetwork is a packet sent from the server to the client and vise-versa to communicate
// editor-mode related information. It carries a single compound tag containing the relevant
// information.
type EditorNetwork struct {
	// RouteToManager ...
	RouteToManager bool
	// Payload is a network little endian compound tag holding data relevant to the editor.
	Payload []byte
}

// Marshal reads or writes EditorNetwork using its canonical wire layout.
func (x *EditorNetwork) Marshal(io protocol.IO) {
	io.Bool(&x.RouteToManager)
	io.NBT(&x.Payload, protocol.NBTNetwork)
}

// ID returns the protocol ID for EditorNetwork.
func (*EditorNetwork) ID() uint32 { return IDEditorNetwork }
