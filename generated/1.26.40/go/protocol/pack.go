// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type PackIDVersion struct {
	PackUUID    uuid.UUID
	PackVersion SemVersion
}

// Marshal reads or writes PackIDVersion using its canonical wire layout.
func (x *PackIDVersion) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

type PackIDVersionData struct {
	PackUUID    uuid.UUID
	PackVersion SemVersionData
}

// Marshal reads or writes PackIDVersionData using its canonical wire layout.
func (x *PackIDVersionData) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

type PackInfoData struct {
	PackIDVersion       PackIDVersionData
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
	x.PackIDVersion.Marshal(io)
	io.Uint64(&x.PackSize)
	io.String(&x.ContentKey)
	io.String(&x.SubpackName)
	x.ContentIdentity.Marshal(io)
	io.Bool(&x.HasScripts)
	io.Bool(&x.IsAddonPack)
	io.Bool(&x.IsRayTracingCapable)
	io.String(&x.CDNURL)
}

type PackInstanceID struct {
	PackID      string
	Version     string
	SubPackName string
}

// Marshal reads or writes PackInstanceID using its canonical wire layout.
func (x *PackInstanceID) Marshal(io IO) {
	io.String(&x.PackID)
	io.String(&x.Version)
	io.String(&x.SubPackName)
}
