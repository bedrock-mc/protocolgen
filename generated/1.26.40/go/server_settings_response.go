// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerSettingsResponse struct {
	FormID     uint32
	FormUIJSON string
}

// Marshal reads or writes ServerSettingsResponse using its canonical wire layout.
func (x *ServerSettingsResponse) Marshal(io IO) {
	io.Varuint32(&x.FormID)
	io.String(&x.FormUIJSON)
}
