// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TypedServerNetIdStructRecipeNetIdTag struct {
	RawId uint32
}

// Marshal reads or writes TypedServerNetIdStructRecipeNetIdTag using its canonical wire layout.
func (x *TypedServerNetIdStructRecipeNetIdTag) Marshal(io IO) {
	io.Varuint32(&x.RawId)
}
