// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ECSProfilingDiagnosticsSystemDiagnosticTimingInfo struct {
	DisplayName    string
	SystemIndex    uint64
	TimeInNS       uint64
	PercentOfTotal uint8
}

// Marshal reads or writes ECSProfilingDiagnosticsSystemDiagnosticTimingInfo using its canonical wire layout.
func (x *ECSProfilingDiagnosticsSystemDiagnosticTimingInfo) Marshal(io IO) {
	io.String(&x.DisplayName)
	io.Uint64(&x.SystemIndex)
	io.Uint64(&x.TimeInNS)
	io.Uint8(&x.PercentOfTotal)
}
