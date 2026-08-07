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
	if !io.Reading() && uint64(len(x.MemoryCategoryValues)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MemoryCategoryValues), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.MemoryCategoryValues))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.MemoryCategoryValues = make([]MemoryMemoryCategoryCounter, int(count1))
	}
	for index2 := range x.MemoryCategoryValues {
		x.MemoryCategoryValues[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.EntityDiagnostics)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EntityDiagnostics), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.EntityDiagnostics))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.EntityDiagnostics = make([]ECSProfilingDiagnosticsEntityDiagnosticTimingInfo, int(count3))
	}
	for index4 := range x.EntityDiagnostics {
		x.EntityDiagnostics[index4].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.SystemDiagnostics)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SystemDiagnostics), "collection length overflows uint32")
		return
	}
	count5 := uint32(len(x.SystemDiagnostics))
	io.Varuint32(&count5)
	if io.Reading() {
		if uint64(count5) > uint64(^uint(0)>>1) {
			io.InvalidValue(count5, "collection length overflows int")
			return
		}
		x.SystemDiagnostics = make([]ECSProfilingDiagnosticsSystemDiagnosticTimingInfo, int(count5))
	}
	for index6 := range x.SystemDiagnostics {
		x.SystemDiagnostics[index6].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.SystemCategories)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SystemCategories), "collection length overflows uint32")
		return
	}
	count7 := uint32(len(x.SystemCategories))
	io.Varuint32(&count7)
	if io.Reading() {
		if uint64(count7) > uint64(^uint(0)>>1) {
			io.InvalidValue(count7, "collection length overflows int")
			return
		}
		x.SystemCategories = make([]ECSProfilingDiagnosticsSystemCategory, int(count7))
	}
	for index8 := range x.SystemCategories {
		x.SystemCategories[index8].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.WhiskerScopes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.WhiskerScopes), "collection length overflows uint32")
		return
	}
	count9 := uint32(len(x.WhiskerScopes))
	io.Varuint32(&count9)
	if io.Reading() {
		if uint64(count9) > uint64(^uint(0)>>1) {
			io.InvalidValue(count9, "collection length overflows int")
			return
		}
		x.WhiskerScopes = make([]BedrockProfileWhiskerDiagnosticsScopeDataSummary, int(count9))
	}
	for index10 := range x.WhiskerScopes {
		x.WhiskerScopes[index10].Marshal(io)
	}
}
