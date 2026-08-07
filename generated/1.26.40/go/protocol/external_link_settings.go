// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ExternalLinkSettings struct {
	URL         string
	DisplayName string
}

// Marshal reads or writes ExternalLinkSettings using its canonical wire layout.
func (x *ExternalLinkSettings) Marshal(io IO) {
	io.String(&x.URL)
	io.String(&x.DisplayName)
}
