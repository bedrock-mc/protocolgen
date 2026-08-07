// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CurrentStructureFeature struct {
	CurrentStructureFeature string
}

// Marshal reads or writes CurrentStructureFeature using its canonical wire layout.
func (x *CurrentStructureFeature) Marshal(io IO) {
	io.String(&x.CurrentStructureFeature)
}
