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
	enumValue1 := int32(x.TitleType)
	io.Varint32(&enumValue1)
	x.TitleType = SetTitleTitleType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.String(&x.TitleText)
	io.Varint32(&x.FadeInTime)
	io.Varint32(&x.StayTime)
	io.Varint32(&x.FadeOutTime)
	io.String(&x.Xuid)
	io.String(&x.PlatformOnlineId)
	io.String(&x.FilteredTitleMessage)
}
