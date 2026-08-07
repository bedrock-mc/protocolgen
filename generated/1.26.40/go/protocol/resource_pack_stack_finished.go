// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackStackFinished struct {
	ResponseType string
}

func (*ResourcePackStackFinished) isResourcePackClientResponseData() {}

// Marshal reads or writes ResourcePackStackFinished using its canonical wire layout.
func (x *ResourcePackStackFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}
