// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Text struct {
	Localize        bool
	Body            TextBody
	SenderSXUID     string
	PlatformId      string
	FilteredMessage Optional[string]
}

// Marshal reads or writes Text using its canonical wire layout.
func (x *Text) Marshal(io IO) {
	io.Bool(&x.Localize)
	marshalTextBody(io, &x.Body)
	io.String(&x.SenderSXUID)
	io.String(&x.PlatformId)
	OptionalFunc(io, &x.FilteredMessage, io.String)
}
