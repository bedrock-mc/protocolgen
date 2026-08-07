// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ContainerRegistryCleanup struct {
	RemovedContainers []FullContainerName
}

// Marshal reads or writes ContainerRegistryCleanup using its canonical wire layout.
func (x *ContainerRegistryCleanup) Marshal(io IO) {
	FuncSlice(io, &x.RemovedContainers, io.Varuint32, func(value *FullContainerName) {
		value.Marshal(io)
	})
}
