// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PackInfoData struct {
	PackIdVersion       PackIdVersionData
	PackSize            uint64
	ContentKey          string
	SubpackName         string
	ContentIdentity     ContentIdentity
	HasScripts          bool
	IsAddonPack         bool
	IsRayTracingCapable bool
	CDNURL              string
}

// Marshal reads or writes PackInfoData using its canonical wire layout.
func (x *PackInfoData) Marshal(io IO) {
	x.PackIdVersion.Marshal(io)
	io.Uint64(&x.PackSize)
	io.String(&x.ContentKey)
	io.String(&x.SubpackName)
	x.ContentIdentity.Marshal(io)
	io.Bool(&x.HasScripts)
	io.Bool(&x.IsAddonPack)
	io.Bool(&x.IsRayTracingCapable)
	io.String(&x.CDNURL)
}
