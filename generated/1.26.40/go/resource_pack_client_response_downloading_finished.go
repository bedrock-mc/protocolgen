// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ResourcePackClientResponseDownloadingFinished struct {
	ResponseType string
}

func (ResourcePackClientResponseDownloadingFinished) isResourcePackClientResponseResponse() {}

// Marshal reads or writes ResourcePackClientResponseDownloadingFinished using its canonical wire layout.
func (x *ResourcePackClientResponseDownloadingFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}
