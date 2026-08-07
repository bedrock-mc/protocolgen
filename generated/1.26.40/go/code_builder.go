// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CodeBuilder struct {
	URL                   string
	ShouldOpenCodeBuilder bool
}

// Marshal reads or writes CodeBuilder using its canonical wire layout.
func (x *CodeBuilder) Marshal(io IO) {
	io.String(&x.URL)
	io.Bool(&x.ShouldOpenCodeBuilder)
}
