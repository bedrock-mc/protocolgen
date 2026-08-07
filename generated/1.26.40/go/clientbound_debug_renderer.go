// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData Optional[ClientboundDebugRendererDebugMarkerData]
}

// Marshal reads or writes ClientboundDebugRenderer using its canonical wire layout.
func (x *ClientboundDebugRenderer) Marshal(io IO) {
	io.String(&x.Type)
	io.Bool(&x.DebugMarkerData.set)
	if x.DebugMarkerData.set {
		x.DebugMarkerData.val.Marshal(io)
	} else if io.Reading() {
		var zero ClientboundDebugRendererDebugMarkerData
		x.DebugMarkerData.val = zero
	}
}
