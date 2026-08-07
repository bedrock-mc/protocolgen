// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseData interface {
	isResourcePackClientResponseData()
}

// MarshalResourcePackClientResponseData reads or writes the ResourcePackClientResponseData union using its canonical wire layout.
func MarshalResourcePackClientResponseData(io IO, x *ResourcePackClientResponseData) {
	UnionFunc(io,
		func() {
			var tag int8
			io.Int8(&tag)
			switch int64(tag) {
			case 1:
				value := new(Cancel)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(Downloading)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(DownloadingFinished)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(ResourcePackStackFinished)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *Cancel:
				tag := int8(1)
				io.Int8(&tag)
				value.Marshal(io)
			case *Downloading:
				tag := int8(2)
				io.Int8(&tag)
				value.Marshal(io)
			case *DownloadingFinished:
				tag := int8(3)
				io.Int8(&tag)
				value.Marshal(io)
			case *ResourcePackStackFinished:
				tag := int8(4)
				io.Int8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type ResourcePackStackFinished struct {
	ResponseType string
}

func (*ResourcePackStackFinished) isResourcePackClientResponseData() {}

// Marshal reads or writes ResourcePackStackFinished using its canonical wire layout.
func (x *ResourcePackStackFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}
