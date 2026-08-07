// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ContainerRegistryCleanup struct {
	RemovedContainers []FullContainerName
}

// Marshal reads or writes ContainerRegistryCleanup using its canonical wire layout.
func (x *ContainerRegistryCleanup) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.RemovedContainers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.RemovedContainers), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.RemovedContainers))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.RemovedContainers = make([]FullContainerName, int(count1))
	}
	for index2 := range x.RemovedContainers {
		x.RemovedContainers[index2].Marshal(io)
	}
}
