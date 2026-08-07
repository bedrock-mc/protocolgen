// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TrimMaterial struct {
	MaterialId string
	Color      string
	ItemName   string
}

// Marshal reads or writes TrimMaterial using its canonical wire layout.
func (x *TrimMaterial) Marshal(io IO) {
	io.String(&x.MaterialId)
	io.String(&x.Color)
	io.String(&x.ItemName)
}
