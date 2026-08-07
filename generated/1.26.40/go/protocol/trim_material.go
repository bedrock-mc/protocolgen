// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TrimMaterial struct {
	MaterialID string
	Color      string
	ItemName   string
}

// Marshal reads or writes TrimMaterial using its canonical wire layout.
func (x *TrimMaterial) Marshal(io IO) {
	io.String(&x.MaterialID)
	io.String(&x.Color)
	io.String(&x.ItemName)
}
