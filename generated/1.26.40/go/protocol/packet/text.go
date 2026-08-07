// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Text struct {
	Localize        bool
	Body            protocol.TextData
	SenderSXUID     string
	PlatformID      string
	FilteredMessage protocol.Optional[string]
}

// Marshal reads or writes Text using its canonical wire layout.
func (x *Text) Marshal(io protocol.IO) {
	io.Bool(&x.Localize)
	protocol.MarshalTextData(io, &x.Body)
	io.String(&x.SenderSXUID)
	io.String(&x.PlatformID)
	protocol.OptionalFunc(io, &x.FilteredMessage, io.String)
}

// ID returns the protocol ID for Text.
func (*Text) ID() uint32 { return IDText }
