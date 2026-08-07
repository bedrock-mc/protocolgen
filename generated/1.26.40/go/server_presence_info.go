// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerPresenceInfo struct {
	PresenceConfiguration Optional[ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerPresenceInfo using its canonical wire layout.
func (x *ServerPresenceInfo) Marshal(io IO) {
	io.Bool(&x.PresenceConfiguration.set)
	if x.PresenceConfiguration.set {
		x.PresenceConfiguration.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationPresenceConfiguration
		x.PresenceConfiguration.val = zero
	}
}
