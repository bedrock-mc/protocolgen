// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ToastRequest struct {
	Title   string
	Content string
}

// Marshal reads or writes ToastRequest using its canonical wire layout.
func (x *ToastRequest) Marshal(io IO) {
	io.String(&x.Title)
	io.String(&x.Content)
}
