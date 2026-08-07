// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EditorNetwork struct {
	RouteToManager bool
	RawVariantName string
	RawVariantData string
}

// Marshal reads or writes EditorNetwork using its canonical wire layout.
func (x *EditorNetwork) Marshal(io IO) {
	io.Bool(&x.RouteToManager)
	io.String(&x.RawVariantName)
	io.String(&x.RawVariantData)
}
