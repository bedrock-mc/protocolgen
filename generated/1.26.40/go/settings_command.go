// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SettingsCommand struct {
	Command        string
	SuppressOutput bool
}

// Marshal reads or writes SettingsCommand using its canonical wire layout.
func (x *SettingsCommand) Marshal(io IO) {
	io.String(&x.Command)
	io.Bool(&x.SuppressOutput)
}
