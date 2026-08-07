// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePackChunkRequest struct {
	ResourceName string
	Chunk        int32
}

// Marshal reads or writes ResourcePackChunkRequest using its canonical wire layout.
func (x *ResourcePackChunkRequest) Marshal(io IO) {
	io.String(&x.ResourceName)
	io.Int32(&x.Chunk)
}
