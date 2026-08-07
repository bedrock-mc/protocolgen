// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientCacheStatus struct {
	IsCacheSupported bool
}

// Marshal reads or writes ClientCacheStatus using its canonical wire layout.
func (x *ClientCacheStatus) Marshal(io IO) {
	io.Bool(&x.IsCacheSupported)
}
