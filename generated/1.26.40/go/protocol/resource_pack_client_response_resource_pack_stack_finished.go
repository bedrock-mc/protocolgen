// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseResourcePackStackFinished struct {
	ResponseType string
}

func (ResourcePackClientResponseResourcePackStackFinished) isResourcePackClientResponseResponse() {}

// Marshal reads or writes ResourcePackClientResponseResourcePackStackFinished using its canonical wire layout.
func (x *ResourcePackClientResponseResourcePackStackFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}
