// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PropertySyncDataPropertySyncFloatEntry struct {
	PropertyIndex uint32
	Data          float32
}

// Marshal reads or writes PropertySyncDataPropertySyncFloatEntry using its canonical wire layout.
func (x *PropertySyncDataPropertySyncFloatEntry) Marshal(io IO) {
	io.Varuint32(&x.PropertyIndex)
	io.Float32(&x.Data)
}
