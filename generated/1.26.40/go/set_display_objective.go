// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetDisplayObjective struct {
	DisplaySlotName      string
	ObjectiveName        string
	ObjectiveDisplayName string
	CriteriaName         string
	SortOrder            int32
}

// Marshal reads or writes SetDisplayObjective using its canonical wire layout.
func (x *SetDisplayObjective) Marshal(io IO) {
	io.String(&x.DisplaySlotName)
	io.String(&x.ObjectiveName)
	io.String(&x.ObjectiveDisplayName)
	io.String(&x.CriteriaName)
	io.Varint32(&x.SortOrder)
}
