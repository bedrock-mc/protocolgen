// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePackClientResponse struct {
	Response ResourcePackClientResponseResponse
}

// Marshal reads or writes ResourcePackClientResponse using its canonical wire layout.
func (x *ResourcePackClientResponse) Marshal(io IO) {
	marshalResourcePackClientResponseResponse(io, &x.Response)
}
