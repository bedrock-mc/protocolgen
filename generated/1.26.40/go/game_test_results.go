// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GameTestResults struct {
	Succeeded bool
	Error     string
	TestName  string
}

// Marshal reads or writes GameTestResults using its canonical wire layout.
func (x *GameTestResults) Marshal(io IO) {
	io.Bool(&x.Succeeded)
	io.String(&x.Error)
	io.String(&x.TestName)
}
