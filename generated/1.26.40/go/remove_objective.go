// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RemoveObjective struct {
	ObjectiveName string
}

// Marshal reads or writes RemoveObjective using its canonical wire layout.
func (x *RemoveObjective) Marshal(io IO) {
	io.String(&x.ObjectiveName)
}
