// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ResourcePackClientResponseDownloading struct {
	ResponseType     string
	DownloadingPacks []string
}

func (ResourcePackClientResponseDownloading) isResourcePackClientResponseResponse() {}

// Marshal reads or writes ResourcePackClientResponseDownloading using its canonical wire layout.
func (x *ResourcePackClientResponseDownloading) Marshal(io IO) {
	io.String(&x.ResponseType)
	FuncSlice(io, &x.DownloadingPacks, io.Varuint32, io.String)
}
