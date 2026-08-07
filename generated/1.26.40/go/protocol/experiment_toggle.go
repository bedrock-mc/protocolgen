// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ExperimentToggle struct {
	Name    string
	Enabled bool
}

// Marshal reads or writes ExperimentToggle using its canonical wire layout.
func (x *ExperimentToggle) Marshal(io IO) {
	io.String(&x.Name)
	io.Bool(&x.Enabled)
}
