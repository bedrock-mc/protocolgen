// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type VoxelShapesRegistryHandle struct {
	Value uint16
}

// Marshal reads or writes VoxelShapesRegistryHandle using its canonical wire layout.
func (x *VoxelShapesRegistryHandle) Marshal(io IO) {
	io.Uint16(&x.Value)
}
