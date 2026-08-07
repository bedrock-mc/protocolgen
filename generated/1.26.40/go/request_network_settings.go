// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type RequestNetworkSettings struct {
	ClientNetworkVersion int32
}

// Marshal reads or writes RequestNetworkSettings using its canonical wire layout.
func (x *RequestNetworkSettings) Marshal(io IO) {
	io.BEInt32(&x.ClientNetworkVersion)
}
