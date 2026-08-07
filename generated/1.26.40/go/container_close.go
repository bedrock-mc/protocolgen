// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ContainerClose struct {
	ContainerId          uint8
	ContainerType        uint8
	ServerInitiatedClose bool
}

// Marshal reads or writes ContainerClose using its canonical wire layout.
func (x *ContainerClose) Marshal(io IO) {
	io.Uint8(&x.ContainerId)
	io.Uint8(&x.ContainerType)
	io.Bool(&x.ServerInitiatedClose)
}
