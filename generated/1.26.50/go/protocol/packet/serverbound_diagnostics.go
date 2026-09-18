// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ServerboundDiagnostics struct {
	AvgFps                    float32
	AvgServerSimTickTimeMS    float32
	AvgClientSimTickTimeMS    float32
	AvgBeginFrameTimeMS       float32
	AvgInputTimeMS            float32
	AvgRenderTimeMS           float32
	AvgEndFrameTimeMS         float32
	AvgRemainderTimePercent   float32
	AvgUnaccountedTimePercent float32
	MemoryCategoryValues      []protocol.MemoryCategoryCounter
	EntityDiagnostics         []protocol.ECSProfilingDiagnosticsEntityDiagnosticTimingInfo
	SystemDiagnostics         []protocol.ECSProfilingDiagnosticsSystemDiagnosticTimingInfo
	SystemCategories          []protocol.ECSProfilingDiagnosticsSystemCategory
	WhiskerScopes             []protocol.BedrockProfileWhiskerDiagnosticsScopeDataSummary
}

// ID returns the protocol ID for ServerboundDiagnostics.
func (*ServerboundDiagnostics) ID() uint32 { return IDServerboundDiagnostics }

// Marshal reads or writes ServerboundDiagnostics using its canonical wire layout.
func (pk *ServerboundDiagnostics) Marshal(io protocol.IO) {
	io.Float32(&pk.AvgFps)
	io.Float32(&pk.AvgServerSimTickTimeMS)
	io.Float32(&pk.AvgClientSimTickTimeMS)
	io.Float32(&pk.AvgBeginFrameTimeMS)
	io.Float32(&pk.AvgInputTimeMS)
	io.Float32(&pk.AvgRenderTimeMS)
	io.Float32(&pk.AvgEndFrameTimeMS)
	io.Float32(&pk.AvgRemainderTimePercent)
	io.Float32(&pk.AvgUnaccountedTimePercent)
	protocol.Slice(io, &pk.MemoryCategoryValues)
	protocol.Slice(io, &pk.EntityDiagnostics)
	protocol.Slice(io, &pk.SystemDiagnostics)
	protocol.Slice(io, &pk.SystemCategories)
	protocol.Slice(io, &pk.WhiskerScopes)
}
