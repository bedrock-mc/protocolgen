// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
type Text struct {
	Localize bool
	// TextType is the type of the text sent. When a client sends this to the server, it should always be
	// TextTypeChat. If the server sends it, it may be one of the other text types above.
	MessageCategory uint8
	Body            protocol.TextData
	// SourceName is the name of the source of the messages. This source is displayed in text types such as the
	// TextTypeChat and TextTypeWhisper, where typically the username is shown.
	SenderSXUID string
	// Message is the message of the packet. This field is set for each TextType and is the main component of the
	// packet.
	PlatformID string
	// FilteredMessage is a filtered version of Message with all the profanity removed. The client will use this
	// over Message if this field is not empty and they have the "Filter Profanity" setting enabled.
	FilteredMessage protocol.Optional[string]
}

// ID ...
func (*Text) ID() uint32 {
	return IDText
}

func (pk *Text) Marshal(io protocol.IO) {
	io.Bool(&pk.Localize)
	io.Uint8(&pk.MessageCategory)
	protocol.MarshalTextData(io, &pk.Body)
	io.StringLimits(&pk.SenderSXUID, 0, 64)
	io.StringLimits(&pk.PlatformID, 0, 256)
	protocol.OptionalFunc(io, &pk.FilteredMessage, func(value *string) {
		io.StringLimits(value, 0, 65536)
	})
}
