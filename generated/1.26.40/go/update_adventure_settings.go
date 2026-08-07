// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateAdventureSettings struct {
	AdventureSettings AdventureSettings
}

// Marshal reads or writes UpdateAdventureSettings using its canonical wire layout.
func (x *UpdateAdventureSettings) Marshal(io IO) {
	x.AdventureSettings.Marshal(io)
}
