// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistCategoryDefinition struct {
	Name       string
	Priorities CameraAimAssistCategoryPriorities
}

// Marshal reads or writes CameraAimAssistCategoryDefinition using its canonical wire layout.
func (x *CameraAimAssistCategoryDefinition) Marshal(io IO) {
	io.String(&x.Name)
	x.Priorities.Marshal(io)
}
