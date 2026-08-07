// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerStats struct {
	ServerTime  float32
	NetworkTime float32
}

// Marshal reads or writes ServerStats using its canonical wire layout.
func (x *ServerStats) Marshal(io IO) {
	io.Float32(&x.ServerTime)
	io.Float32(&x.NetworkTime)
}
