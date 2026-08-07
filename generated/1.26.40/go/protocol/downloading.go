// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type Downloading struct {
	ResponseType     string
	DownloadingPacks []string
}

func (*Downloading) isResourcePackClientResponseData() {}

// Marshal reads or writes Downloading using its canonical wire layout.
func (x *Downloading) Marshal(io IO) {
	io.String(&x.ResponseType)
	FuncSlice(io, &x.DownloadingPacks, io.Varuint32, io.String)
}
