// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type Experiments struct {
	Toggles                []CerealizerExperimentsAnonExperimentToggle
	ExperimentsEverToggled bool
}

// Marshal reads or writes Experiments using its canonical wire layout.
func (x *Experiments) Marshal(io IO) {
	FuncSlice(io, &x.Toggles, io.Uint32, func(value *CerealizerExperimentsAnonExperimentToggle) {
		value.Marshal(io)
	})
	io.Bool(&x.ExperimentsEverToggled)
}
