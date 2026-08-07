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
