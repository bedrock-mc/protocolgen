// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// Text is sent by the client to the server to send chat messages, and by the server to the client
// to forward or send messages, which may be chat, popups, tips etc.
type Text struct {
	Localize        bool
	MessageCategory uint8
	Body            protocol.TextData
	SenderSXUID     string
	PlatformID      string
	// FilteredMessage is a filtered version of Message with all the profanity removed. The client will
	// use this over Message if this field is not empty and they have the "Filter Profanity" setting
	// enabled.
	FilteredMessage protocol.Optional[string]
}

// Marshal reads or writes Text using its canonical wire layout.
func (x *Text) Marshal(io protocol.IO) {
	io.Bool(&x.Localize)
	io.Uint8(&x.MessageCategory)
	protocol.MarshalTextData(io, &x.Body)
	io.String(&x.SenderSXUID)
	io.String(&x.PlatformID)
	protocol.OptionalFunc(io, &x.FilteredMessage, io.String)
}

// ID returns the protocol ID for Text.
func (*Text) ID() uint32 { return IDText }
