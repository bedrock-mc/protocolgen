// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration Optional[ServerConfigurationGatheringsConfigurationJoinInfo]
}

// Marshal reads or writes Transfer using its canonical wire layout.
func (x *Transfer) Marshal(io IO) {
	io.String(&x.ServerAddress)
	io.Uint16(&x.ServerPort)
	io.Bool(&x.ReloadWorld)
	io.Bool(&x.GatheringsConfiguration.set)
	if x.GatheringsConfiguration.set {
		x.GatheringsConfiguration.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationGatheringsConfigurationJoinInfo
		x.GatheringsConfiguration.val = zero
	}
}
