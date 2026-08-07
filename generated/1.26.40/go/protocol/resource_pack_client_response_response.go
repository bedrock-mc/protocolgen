// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseResponse interface {
	isResourcePackClientResponseResponse()
}

// MarshalResourcePackClientResponseResponse reads or writes the ResourcePackClientResponseResponse union using its canonical wire layout.
func MarshalResourcePackClientResponseResponse(io IO, x *ResourcePackClientResponseResponse) {
	UnionFunc(io,
		func() {
			var tag int8
			io.Int8(&tag)
			switch int64(tag) {
			case 1:
				var value ResourcePackClientResponseCancel
				value.Marshal(io)
				*x = value
			case 2:
				var value ResourcePackClientResponseDownloading
				value.Marshal(io)
				*x = value
			case 3:
				var value ResourcePackClientResponseDownloadingFinished
				value.Marshal(io)
				*x = value
			case 4:
				var value ResourcePackClientResponseResourcePackStackFinished
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case ResourcePackClientResponseCancel:
				tag := int8(1)
				io.Int8(&tag)
				value.Marshal(io)
			case ResourcePackClientResponseDownloading:
				tag := int8(2)
				io.Int8(&tag)
				value.Marshal(io)
			case ResourcePackClientResponseDownloadingFinished:
				tag := int8(3)
				io.Int8(&tag)
				value.Marshal(io)
			case ResourcePackClientResponseResourcePackStackFinished:
				tag := int8(4)
				io.Int8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}
