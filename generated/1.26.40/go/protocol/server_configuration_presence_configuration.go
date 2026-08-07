// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerConfigurationPresenceConfiguration struct {
	RichPresenceID Optional[string]
}

// Marshal reads or writes ServerConfigurationPresenceConfiguration using its canonical wire layout.
func (x *ServerConfigurationPresenceConfiguration) Marshal(io IO) {
	OptionalFunc(io, &x.RichPresenceID, io.String)
}
