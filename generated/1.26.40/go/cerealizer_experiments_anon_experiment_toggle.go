// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CerealizerExperimentsAnonExperimentToggle struct {
	Name    string
	Enabled bool
}

// Marshal reads or writes CerealizerExperimentsAnonExperimentToggle using its canonical wire layout.
func (x *CerealizerExperimentsAnonExperimentToggle) Marshal(io IO) {
	io.String(&x.Name)
	io.Bool(&x.Enabled)
}
