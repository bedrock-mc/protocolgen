// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EduUriResource struct {
	EduSharedURIResource EduSharedUriResource
}

// Marshal reads or writes EduUriResource using its canonical wire layout.
func (x *EduUriResource) Marshal(io IO) {
	x.EduSharedURIResource.Marshal(io)
}
