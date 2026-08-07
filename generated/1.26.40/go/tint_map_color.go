// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "image/color"

type TintMapColor struct {
	Colors [4]color.RGBA
}

// Marshal reads or writes TintMapColor using its canonical wire layout.
func (x *TintMapColor) Marshal(io IO) {
	for index1 := range x.Colors {
		io.RGBA(&x.Colors[index1])
	}
}
