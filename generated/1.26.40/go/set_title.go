// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetTitle struct {
	TitleType            SetTitleTitleType
	TitleText            string
	FadeInTime           int32
	StayTime             int32
	FadeOutTime          int32
	Xuid                 string
	PlatformOnlineId     string
	FilteredTitleMessage string
}

// Marshal reads or writes SetTitle using its canonical wire layout.
func (x *SetTitle) Marshal(io IO) {
	IntegerFunc(&x.TitleType, io.Varint32)
	io.String(&x.TitleText)
	io.Varint32(&x.FadeInTime)
	io.Varint32(&x.StayTime)
	io.Varint32(&x.FadeOutTime)
	io.String(&x.Xuid)
	io.String(&x.PlatformOnlineId)
	io.String(&x.FilteredTitleMessage)
}
