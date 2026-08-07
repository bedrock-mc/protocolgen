// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EASBoolAttributeData struct {
	Value     bool
	Operation string
}

func (EASBoolAttributeData) isEAS() {}

// Marshal reads or writes EASBoolAttributeData using its canonical wire layout.
func (x *EASBoolAttributeData) Marshal(io IO) {
	io.Bool(&x.Value)
	io.String(&x.Operation)
}
