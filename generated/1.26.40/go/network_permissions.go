// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NetworkPermissions struct {
	ServerAuthSoundEnabled bool
}

// Marshal reads or writes NetworkPermissions using its canonical wire layout.
func (x *NetworkPermissions) Marshal(io IO) {
	io.Bool(&x.ServerAuthSoundEnabled)
}
