// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData Optional[ClientboundDebugRendererDebugMarkerData]
}

// Marshal reads or writes ClientboundDebugRenderer using its canonical wire layout.
func (x *ClientboundDebugRenderer) Marshal(io IO) {
	io.String(&x.Type)
	OptionalFunc(io, &x.DebugMarkerData, func(value *ClientboundDebugRendererDebugMarkerData) {
		value.Marshal(io)
	})
}
