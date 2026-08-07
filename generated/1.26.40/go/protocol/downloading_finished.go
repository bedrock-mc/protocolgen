// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DownloadingFinished struct {
	ResponseType string
}

func (*DownloadingFinished) isResourcePackClientResponseData() {}

// Marshal reads or writes DownloadingFinished using its canonical wire layout.
func (x *DownloadingFinished) Marshal(io IO) {
	io.String(&x.ResponseType)
}
