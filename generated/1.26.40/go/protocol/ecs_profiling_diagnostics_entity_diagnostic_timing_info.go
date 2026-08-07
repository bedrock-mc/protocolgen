// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ECSProfilingDiagnosticsEntityDiagnosticTimingInfo struct {
	DisplayName    string
	Entity         string
	TimeInNS       uint64
	PercentOfTotal uint8
}

// Marshal reads or writes ECSProfilingDiagnosticsEntityDiagnosticTimingInfo using its canonical wire layout.
func (x *ECSProfilingDiagnosticsEntityDiagnosticTimingInfo) Marshal(io IO) {
	io.String(&x.DisplayName)
	io.String(&x.Entity)
	io.Uint64(&x.TimeInNS)
	io.Uint8(&x.PercentOfTotal)
}
