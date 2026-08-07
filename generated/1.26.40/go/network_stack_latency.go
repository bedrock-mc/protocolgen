// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NetworkStackLatency struct {
	CreationTime uint64
	IsFromServer bool
}

// Marshal reads or writes NetworkStackLatency using its canonical wire layout.
func (x *NetworkStackLatency) Marshal(io IO) {
	io.Uint64(&x.CreationTime)
	io.Bool(&x.IsFromServer)
}
