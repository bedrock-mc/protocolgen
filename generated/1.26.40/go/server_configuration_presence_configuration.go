// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerConfigurationPresenceConfiguration struct {
	RichPresenceId Optional[string]
}

// Marshal reads or writes ServerConfigurationPresenceConfiguration using its canonical wire layout.
func (x *ServerConfigurationPresenceConfiguration) Marshal(io IO) {
	OptionalFunc(io, &x.RichPresenceId, io.String)
}
