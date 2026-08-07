// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

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
	MemoryCategoryValues      []MemoryMemoryCategoryCounter
	EntityDiagnostics         []ECSProfilingDiagnosticsEntityDiagnosticTimingInfo
	SystemDiagnostics         []ECSProfilingDiagnosticsSystemDiagnosticTimingInfo
	SystemCategories          []ECSProfilingDiagnosticsSystemCategory
	WhiskerScopes             []BedrockProfileWhiskerDiagnosticsScopeDataSummary
}

// Marshal reads or writes ServerboundDiagnostics using its canonical wire layout.
func (x *ServerboundDiagnostics) Marshal(io IO) {
	io.Float32(&x.AvgFps)
	io.Float32(&x.AvgServerSimTickTimeMS)
	io.Float32(&x.AvgClientSimTickTimeMS)
	io.Float32(&x.AvgBeginFrameTimeMS)
	io.Float32(&x.AvgInputTimeMS)
	io.Float32(&x.AvgRenderTimeMS)
	io.Float32(&x.AvgEndFrameTimeMS)
	io.Float32(&x.AvgRemainderTimePercent)
	io.Float32(&x.AvgUnaccountedTimePercent)
	FuncSlice(io, &x.MemoryCategoryValues, io.Varuint32, func(value *MemoryMemoryCategoryCounter) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.EntityDiagnostics, io.Varuint32, func(value *ECSProfilingDiagnosticsEntityDiagnosticTimingInfo) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.SystemDiagnostics, io.Varuint32, func(value *ECSProfilingDiagnosticsSystemDiagnosticTimingInfo) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.SystemCategories, io.Varuint32, func(value *ECSProfilingDiagnosticsSystemCategory) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.WhiskerScopes, io.Varuint32, func(value *BedrockProfileWhiskerDiagnosticsScopeDataSummary) {
		value.Marshal(io)
	})
}
