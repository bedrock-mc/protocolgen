// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type Cancel struct {
	ResponseType string
}

func (*Cancel) isResourcePackClientResponseData() {}

// Marshal reads or writes Cancel using its canonical wire layout.
func (x *Cancel) Marshal(io IO) {
	io.String(&x.ResponseType)
}
