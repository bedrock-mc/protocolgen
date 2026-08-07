// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Login struct {
	ClientNetworkVersion int32
	ConnectionRequest    string
}

// Marshal reads or writes Login using its canonical wire layout.
func (x *Login) Marshal(io IO) {
	io.BEInt32(&x.ClientNetworkVersion)
	io.String(&x.ConnectionRequest)
}
