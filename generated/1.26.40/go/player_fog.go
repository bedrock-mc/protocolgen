// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerFog struct {
	FogStack []string
}

// Marshal reads or writes PlayerFog using its canonical wire layout.
func (x *PlayerFog) Marshal(io IO) {
	FuncSlice(io, &x.FogStack, io.Varuint32, io.String)
}
