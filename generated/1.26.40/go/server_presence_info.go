// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerPresenceInfo struct {
	PresenceConfiguration Optional[ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerPresenceInfo using its canonical wire layout.
func (x *ServerPresenceInfo) Marshal(io IO) {
	OptionalFunc(io, &x.PresenceConfiguration, func(value *ServerConfigurationPresenceConfiguration) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
