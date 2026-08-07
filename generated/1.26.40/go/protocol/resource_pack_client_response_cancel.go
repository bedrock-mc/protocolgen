// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseCancel struct {
	ResponseType string
}

func (ResourcePackClientResponseCancel) isResourcePackClientResponseResponse() {}

// Marshal reads or writes ResourcePackClientResponseCancel using its canonical wire layout.
func (x *ResourcePackClientResponseCancel) Marshal(io IO) {
	io.String(&x.ResponseType)
}
