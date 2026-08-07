// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AdventureSettings struct {
	NoPvM          bool
	NoMvP          bool
	ImmutableWorld bool
	ShowNameTags   bool
	AutoJump       bool
}

// Marshal reads or writes AdventureSettings using its canonical wire layout.
func (x *AdventureSettings) Marshal(io IO) {
	io.Bool(&x.NoPvM)
	io.Bool(&x.NoMvP)
	io.Bool(&x.ImmutableWorld)
	io.Bool(&x.ShowNameTags)
	io.Bool(&x.AutoJump)
}
