// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseData interface {
	isResourcePackClientResponseData()
}

// MarshalResourcePackClientResponseData reads or writes the ResourcePackClientResponseData union using its canonical wire layout.
func MarshalResourcePackClientResponseData(io IO, x *ResourcePackClientResponseData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(Cancel)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(Downloading)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(DownloadingFinished)
				value.Marshal(io)
				*x = value
			case 3:
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
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *Downloading:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DownloadingFinished:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ResourcePackStackFinished:
				tag := uint32(3)
				io.Varuint32(&tag)
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
