// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BedrockProfileWhiskerDiagnosticsScopeDataSummary struct {
	Label           string
	Indentation     string
	TotalHighCostNS uint64
	TotalMidCostNS  uint64
	TotalLowCostNS  uint64
}

// Marshal reads or writes BedrockProfileWhiskerDiagnosticsScopeDataSummary using its canonical wire layout.
func (x *BedrockProfileWhiskerDiagnosticsScopeDataSummary) Marshal(io IO) {
	io.String(&x.Label)
	io.String(&x.Indentation)
	io.Uint64(&x.TotalHighCostNS)
	io.Uint64(&x.TotalMidCostNS)
	io.Uint64(&x.TotalLowCostNS)
}
