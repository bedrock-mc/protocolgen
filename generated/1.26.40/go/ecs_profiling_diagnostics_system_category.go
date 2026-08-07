// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ECSProfilingDiagnosticsSystemCategory struct {
	CategoryName string
	SystemIndex  uint64
}

// Marshal reads or writes ECSProfilingDiagnosticsSystemCategory using its canonical wire layout.
func (x *ECSProfilingDiagnosticsSystemCategory) Marshal(io IO) {
	io.String(&x.CategoryName)
	io.Uint64(&x.SystemIndex)
}
