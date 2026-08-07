// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RecipeNetID struct {
	RawID uint32
}

// Marshal reads or writes RecipeNetID using its canonical wire layout.
func (x *RecipeNetID) Marshal(io IO) {
	io.Varuint32(&x.RawID)
}
