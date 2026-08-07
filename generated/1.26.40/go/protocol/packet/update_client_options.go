// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// UpdateClientOptions is sent by the client when some of the client's options are updated, such as
// the graphics mode.
type UpdateClientOptions struct {
	// GraphicsMode is the graphics mode that the client is using. It is one of the constants above.
	GraphicsModeChange protocol.Optional[protocol.GraphicsMode]
	// FilterProfanity is if the client only uses filtered messages or not.
	FilterProfanityChange protocol.Optional[bool]
}

// Marshal reads or writes UpdateClientOptions using its canonical wire layout.
func (x *UpdateClientOptions) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &x.GraphicsModeChange, func(value *protocol.GraphicsMode) {
		protocol.IntegerFunc(value, io.Uint8)
	})
	protocol.OptionalFunc(io, &x.FilterProfanityChange, io.Bool)
}

// ID returns the protocol ID for UpdateClientOptions.
func (*UpdateClientOptions) ID() uint32 { return IDUpdateClientOptions }
