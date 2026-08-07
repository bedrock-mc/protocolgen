// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

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

// Marshal reads or writes ServerboundDiagnostics using its canonical wire layout.
func (x *ServerboundDiagnostics) Marshal(io protocol.IO) {
	io.Float32(&x.AvgFps)
	io.Float32(&x.AvgServerSimTickTimeMS)
	io.Float32(&x.AvgClientSimTickTimeMS)
	io.Float32(&x.AvgBeginFrameTimeMS)
	io.Float32(&x.AvgInputTimeMS)
	io.Float32(&x.AvgRenderTimeMS)
	io.Float32(&x.AvgEndFrameTimeMS)
	io.Float32(&x.AvgRemainderTimePercent)
	io.Float32(&x.AvgUnaccountedTimePercent)
	protocol.Slice(io, &x.MemoryCategoryValues)
	protocol.Slice(io, &x.EntityDiagnostics)
	protocol.Slice(io, &x.SystemDiagnostics)
	protocol.Slice(io, &x.SystemCategories)
	protocol.Slice(io, &x.WhiskerScopes)
}

// ID returns the protocol ID for ServerboundDiagnostics.
func (*ServerboundDiagnostics) ID() uint32 { return IDServerboundDiagnostics }
