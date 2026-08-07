// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetTitle struct {
	TitleType            protocol.TitleType
	TitleText            string
	FadeInTime           int32
	StayTime             int32
	FadeOutTime          int32
	Xuid                 string
	PlatformOnlineID     string
	FilteredTitleMessage string
}

// Marshal reads or writes SetTitle using its canonical wire layout.
func (x *SetTitle) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.TitleType, io.Varint32)
	io.String(&x.TitleText)
	io.Varint32(&x.FadeInTime)
	io.Varint32(&x.StayTime)
	io.Varint32(&x.FadeOutTime)
	io.String(&x.Xuid)
	io.String(&x.PlatformOnlineID)
	io.String(&x.FilteredTitleMessage)
}

// ID returns the protocol ID for SetTitle.
func (*SetTitle) ID() uint32 { return IDSetTitle }
