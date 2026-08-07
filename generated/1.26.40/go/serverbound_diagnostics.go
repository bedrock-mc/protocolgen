// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

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

func (p *ServerboundDiagnostics) Encode(w Encoder) error {
	if err := w.Write("ServerboundDiagnosticsPacket.AvgFps", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgFps); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgServerSimTickTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgServerSimTickTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgClientSimTickTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgClientSimTickTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgBeginFrameTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgBeginFrameTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgInputTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgInputTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgRenderTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgRenderTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgEndFrameTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgEndFrameTimeMS); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgRemainderTimePercent", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgRemainderTimePercent); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.AvgUnaccountedTimePercent", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.AvgUnaccountedTimePercent); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.Memory Category Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "Memory::MemoryCategoryCounter", TypeID: "Memory::MemoryCategoryCounter", Fields: []ShapeField{{Ordinal: 0, Name: "Category", Shape: Shape{Kind: "enum", Semantic: "Memory::MemoryCategory", TypeID: "enums/Memory::MemoryCategory", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Invalid_SizeUnknown", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Actor", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "ActorAnimation", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "ActorRendering", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "BlockTickingQueues", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Biome_Storage", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Blobs", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Cereal", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "CircuitSystem", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "Client", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "Commands", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "DBStorage", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Debug", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Documentation", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "ECSSystems", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "FMOD", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "Fonts", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "ImGui", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "Input", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "JsonUI", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "JsonUI_ControlFactory_Json", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "JsonUI_ControlTree", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "JsonUI_ControlTree_ControlElement", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "JsonUI_ControlTree_PopulateDataBinding", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "JsonUI_ControlTree_PopulateFocus", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "JsonUI_ControlTree_PopulateLayout", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "JsonUI_ControlTree_PopulateOther", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "JsonUI_ControlTree_PopulateSprite", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "JsonUI_ControlTree_PopulateText", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "JsonUI_ControlTree_PopulateTTS", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "JsonUI_ControlTree_Visibility", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "JsonUI_CreateUI", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "JsonUI_Defs", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "JsonUI_LayoutManager", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "JsonUI_LayoutManager_RemoveDependencies", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "JsonUI_LayoutManager_InitVariable", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "Languages", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "Level", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "LevelStructures", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "LevelChunk", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "LevelChunkGen", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "LevelChunkGenThreadLocal", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "LightVolumeManager", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "Network", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "Marketplace", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "Material_DragonCompiledDefinition", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "Material_DragonMaterial", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "Material_DragonResource", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "Material_DragonUniformMap", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "Material_RenderMaterial", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "Material_RenderMaterialGroup", Shape: Shape{Kind: "void"}}, {Value: 52, Name: "Material_VariationManager", Shape: Shape{Kind: "void"}}, {Value: 53, Name: "Molang", Shape: Shape{Kind: "void"}}, {Value: 54, Name: "OreUI", Shape: Shape{Kind: "void"}}, {Value: 55, Name: "OreUI_Client", Shape: Shape{Kind: "void"}}, {Value: 56, Name: "Persona_Pieces", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "Persona_Animations", Shape: Shape{Kind: "void"}}, {Value: 58, Name: "Persona_Textures", Shape: Shape{Kind: "void"}}, {Value: 59, Name: "Persona_Characters", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "Persona_SkinPacks", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "Persona_Repo", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "Player", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "RenderChunk", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "RenderChunk_IndexBuffer", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "RenderChunk_VertexBuffer", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "Rendering", Shape: Shape{Kind: "void"}}, {Value: 67, Name: "Rendering_BgfxInit", Shape: Shape{Kind: "void"}}, {Value: 68, Name: "Rendering_BgfxStartFrame", Shape: Shape{Kind: "void"}}, {Value: 69, Name: "Rendering_BlockTessellator", Shape: Shape{Kind: "void"}}, {Value: 70, Name: "Rendering_EndFrame", Shape: Shape{Kind: "void"}}, {Value: 71, Name: "Rendering_GraphicsTasksInit", Shape: Shape{Kind: "void"}}, {Value: 72, Name: "Rendering_Library", Shape: Shape{Kind: "void"}}, {Value: 73, Name: "Rendering_PolygonOperatorPool", Shape: Shape{Kind: "void"}}, {Value: 74, Name: "Rendering_PBRTextureData", Shape: Shape{Kind: "void"}}, {Value: 75, Name: "Rendering_RenderRegistry", Shape: Shape{Kind: "void"}}, {Value: 76, Name: "Rendering_Setup", Shape: Shape{Kind: "void"}}, {Value: 77, Name: "Rendering_Vertices", Shape: Shape{Kind: "void"}}, {Value: 78, Name: "RequestLog", Shape: Shape{Kind: "void"}}, {Value: 79, Name: "ResourcePacks", Shape: Shape{Kind: "void"}}, {Value: 80, Name: "Sound", Shape: Shape{Kind: "void"}}, {Value: 81, Name: "SubChunk_BiomeData", Shape: Shape{Kind: "void"}}, {Value: 82, Name: "SubChunk_BlockData", Shape: Shape{Kind: "void"}}, {Value: 83, Name: "SubChunk_LightData", Shape: Shape{Kind: "void"}}, {Value: 84, Name: "Textures", Shape: Shape{Kind: "void"}}, {Value: 85, Name: "WeatherRenderer", Shape: Shape{Kind: "void"}}, {Value: 86, Name: "World_Generator", Shape: Shape{Kind: "void"}}, {Value: 87, Name: "Tasks", Shape: Shape{Kind: "void"}}, {Value: 88, Name: "Test", Shape: Shape{Kind: "void"}}, {Value: 89, Name: "Test_LoadTestTags", Shape: Shape{Kind: "void"}}, {Value: 90, Name: "Scripting", Shape: Shape{Kind: "void"}}, {Value: 91, Name: "Scripting_Runtime", Shape: Shape{Kind: "void"}}, {Value: 92, Name: "Scripting_Context", Shape: Shape{Kind: "void"}}, {Value: 93, Name: "Scripting_Context_Bindings_MC", Shape: Shape{Kind: "void"}}, {Value: 94, Name: "Scripting_Context_Bindings_GT", Shape: Shape{Kind: "void"}}, {Value: 95, Name: "Scripting_Context_Run", Shape: Shape{Kind: "void"}}, {Value: 96, Name: "DataDrivenUI", Shape: Shape{Kind: "void"}}, {Value: 97, Name: "DataDrivenUI_Defs", Shape: Shape{Kind: "void"}}, {Value: 98, Name: "Gameface", Shape: Shape{Kind: "void"}}, {Value: 99, Name: "Gameface_System", Shape: Shape{Kind: "void"}}, {Value: 100, Name: "Gameface_DOM", Shape: Shape{Kind: "void"}}, {Value: 101, Name: "Gameface_CSS", Shape: Shape{Kind: "void"}}, {Value: 102, Name: "Gameface_Display", Shape: Shape{Kind: "void"}}, {Value: 103, Name: "Gameface_TempAllocator", Shape: Shape{Kind: "void"}}, {Value: 104, Name: "Gameface_PoolAllocator", Shape: Shape{Kind: "void"}}, {Value: 105, Name: "Gameface_Dump", Shape: Shape{Kind: "void"}}, {Value: 106, Name: "Gameface_Media", Shape: Shape{Kind: "void"}}, {Value: 107, Name: "Gameface_JSON", Shape: Shape{Kind: "void"}}, {Value: 108, Name: "Gameface_ScriptEngine", Shape: Shape{Kind: "void"}}, {Value: 109, Name: "Gameface_Script", Shape: Shape{Kind: "void"}}, {Value: 110, Name: "Gameface_Layout", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Current Bytes", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.MemoryCategoryValues); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.Entity Diagnostics", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::EntityDiagnosticTimingInfo", TypeID: "ECS::Profiling::Diagnostics::EntityDiagnosticTimingInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Entity", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Time in NS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "Percent of Total", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}, p.EntityDiagnostics); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.System Diagnostics", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::SystemDiagnosticTimingInfo", TypeID: "ECS::Profiling::Diagnostics::SystemDiagnosticTimingInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "System Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 2, Name: "Time in NS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "Percent of Total", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}, p.SystemDiagnostics); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.System Categories", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::SystemCategory", TypeID: "ECS::Profiling::Diagnostics::SystemCategory", Fields: []ShapeField{{Ordinal: 0, Name: "Category Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "System Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.SystemCategories); err != nil {
		return err
	}
	if err := w.Write("ServerboundDiagnosticsPacket.Whisker Scopes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "Bedrock::Profile::Whisker::Diagnostics::ScopeDataSummary", TypeID: "Bedrock::Profile::Whisker::Diagnostics::ScopeDataSummary", Fields: []ShapeField{{Ordinal: 0, Name: "Label", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Indentation", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "TotalHighCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "TotalMidCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 4, Name: "TotalLowCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.WhiskerScopes); err != nil {
		return err
	}
	return nil
}

func DecodeServerboundDiagnostics(r Decoder) (ServerboundDiagnostics, error) {
	var p ServerboundDiagnostics
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgFps", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgFps has unexpected decoded type %T", raw)
		}
		p.AvgFps = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgServerSimTickTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgServerSimTickTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgServerSimTickTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgClientSimTickTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgClientSimTickTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgClientSimTickTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgBeginFrameTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgBeginFrameTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgBeginFrameTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgInputTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgInputTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgInputTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgRenderTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgRenderTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgRenderTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgEndFrameTimeMS", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgEndFrameTimeMS has unexpected decoded type %T", raw)
		}
		p.AvgEndFrameTimeMS = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgRemainderTimePercent", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgRemainderTimePercent has unexpected decoded type %T", raw)
		}
		p.AvgRemainderTimePercent = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.AvgUnaccountedTimePercent", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.AvgUnaccountedTimePercent has unexpected decoded type %T", raw)
		}
		p.AvgUnaccountedTimePercent = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.Memory Category Values", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "Memory::MemoryCategoryCounter", TypeID: "Memory::MemoryCategoryCounter", Fields: []ShapeField{{Ordinal: 0, Name: "Category", Shape: Shape{Kind: "enum", Semantic: "Memory::MemoryCategory", TypeID: "enums/Memory::MemoryCategory", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Invalid_SizeUnknown", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Actor", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "ActorAnimation", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "ActorRendering", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "BlockTickingQueues", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Biome_Storage", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Blobs", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "Cereal", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "CircuitSystem", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "Client", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "Commands", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "DBStorage", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "Debug", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "Documentation", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "ECSSystems", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "FMOD", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "Fonts", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "ImGui", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "Input", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "JsonUI", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "JsonUI_ControlFactory_Json", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "JsonUI_ControlTree", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "JsonUI_ControlTree_ControlElement", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "JsonUI_ControlTree_PopulateDataBinding", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "JsonUI_ControlTree_PopulateFocus", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "JsonUI_ControlTree_PopulateLayout", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "JsonUI_ControlTree_PopulateOther", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "JsonUI_ControlTree_PopulateSprite", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "JsonUI_ControlTree_PopulateText", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "JsonUI_ControlTree_PopulateTTS", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "JsonUI_ControlTree_Visibility", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "JsonUI_CreateUI", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "JsonUI_Defs", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "JsonUI_LayoutManager", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "JsonUI_LayoutManager_RemoveDependencies", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "JsonUI_LayoutManager_InitVariable", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "Languages", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "Level", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "LevelStructures", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "LevelChunk", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "LevelChunkGen", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "LevelChunkGenThreadLocal", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "LightVolumeManager", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "Network", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "Marketplace", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "Material_DragonCompiledDefinition", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "Material_DragonMaterial", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "Material_DragonResource", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "Material_DragonUniformMap", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "Material_RenderMaterial", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "Material_RenderMaterialGroup", Shape: Shape{Kind: "void"}}, {Value: 52, Name: "Material_VariationManager", Shape: Shape{Kind: "void"}}, {Value: 53, Name: "Molang", Shape: Shape{Kind: "void"}}, {Value: 54, Name: "OreUI", Shape: Shape{Kind: "void"}}, {Value: 55, Name: "OreUI_Client", Shape: Shape{Kind: "void"}}, {Value: 56, Name: "Persona_Pieces", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "Persona_Animations", Shape: Shape{Kind: "void"}}, {Value: 58, Name: "Persona_Textures", Shape: Shape{Kind: "void"}}, {Value: 59, Name: "Persona_Characters", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "Persona_SkinPacks", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "Persona_Repo", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "Player", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "RenderChunk", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "RenderChunk_IndexBuffer", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "RenderChunk_VertexBuffer", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "Rendering", Shape: Shape{Kind: "void"}}, {Value: 67, Name: "Rendering_BgfxInit", Shape: Shape{Kind: "void"}}, {Value: 68, Name: "Rendering_BgfxStartFrame", Shape: Shape{Kind: "void"}}, {Value: 69, Name: "Rendering_BlockTessellator", Shape: Shape{Kind: "void"}}, {Value: 70, Name: "Rendering_EndFrame", Shape: Shape{Kind: "void"}}, {Value: 71, Name: "Rendering_GraphicsTasksInit", Shape: Shape{Kind: "void"}}, {Value: 72, Name: "Rendering_Library", Shape: Shape{Kind: "void"}}, {Value: 73, Name: "Rendering_PolygonOperatorPool", Shape: Shape{Kind: "void"}}, {Value: 74, Name: "Rendering_PBRTextureData", Shape: Shape{Kind: "void"}}, {Value: 75, Name: "Rendering_RenderRegistry", Shape: Shape{Kind: "void"}}, {Value: 76, Name: "Rendering_Setup", Shape: Shape{Kind: "void"}}, {Value: 77, Name: "Rendering_Vertices", Shape: Shape{Kind: "void"}}, {Value: 78, Name: "RequestLog", Shape: Shape{Kind: "void"}}, {Value: 79, Name: "ResourcePacks", Shape: Shape{Kind: "void"}}, {Value: 80, Name: "Sound", Shape: Shape{Kind: "void"}}, {Value: 81, Name: "SubChunk_BiomeData", Shape: Shape{Kind: "void"}}, {Value: 82, Name: "SubChunk_BlockData", Shape: Shape{Kind: "void"}}, {Value: 83, Name: "SubChunk_LightData", Shape: Shape{Kind: "void"}}, {Value: 84, Name: "Textures", Shape: Shape{Kind: "void"}}, {Value: 85, Name: "WeatherRenderer", Shape: Shape{Kind: "void"}}, {Value: 86, Name: "World_Generator", Shape: Shape{Kind: "void"}}, {Value: 87, Name: "Tasks", Shape: Shape{Kind: "void"}}, {Value: 88, Name: "Test", Shape: Shape{Kind: "void"}}, {Value: 89, Name: "Test_LoadTestTags", Shape: Shape{Kind: "void"}}, {Value: 90, Name: "Scripting", Shape: Shape{Kind: "void"}}, {Value: 91, Name: "Scripting_Runtime", Shape: Shape{Kind: "void"}}, {Value: 92, Name: "Scripting_Context", Shape: Shape{Kind: "void"}}, {Value: 93, Name: "Scripting_Context_Bindings_MC", Shape: Shape{Kind: "void"}}, {Value: 94, Name: "Scripting_Context_Bindings_GT", Shape: Shape{Kind: "void"}}, {Value: 95, Name: "Scripting_Context_Run", Shape: Shape{Kind: "void"}}, {Value: 96, Name: "DataDrivenUI", Shape: Shape{Kind: "void"}}, {Value: 97, Name: "DataDrivenUI_Defs", Shape: Shape{Kind: "void"}}, {Value: 98, Name: "Gameface", Shape: Shape{Kind: "void"}}, {Value: 99, Name: "Gameface_System", Shape: Shape{Kind: "void"}}, {Value: 100, Name: "Gameface_DOM", Shape: Shape{Kind: "void"}}, {Value: 101, Name: "Gameface_CSS", Shape: Shape{Kind: "void"}}, {Value: 102, Name: "Gameface_Display", Shape: Shape{Kind: "void"}}, {Value: 103, Name: "Gameface_TempAllocator", Shape: Shape{Kind: "void"}}, {Value: 104, Name: "Gameface_PoolAllocator", Shape: Shape{Kind: "void"}}, {Value: 105, Name: "Gameface_Dump", Shape: Shape{Kind: "void"}}, {Value: 106, Name: "Gameface_Media", Shape: Shape{Kind: "void"}}, {Value: 107, Name: "Gameface_JSON", Shape: Shape{Kind: "void"}}, {Value: 108, Name: "Gameface_ScriptEngine", Shape: Shape{Kind: "void"}}, {Value: 109, Name: "Gameface_Script", Shape: Shape{Kind: "void"}}, {Value: 110, Name: "Gameface_Layout", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Current Bytes", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]MemoryMemoryCategoryCounter)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.Memory Category Values has unexpected decoded type %T", raw)
		}
		p.MemoryCategoryValues = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.Entity Diagnostics", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::EntityDiagnosticTimingInfo", TypeID: "ECS::Profiling::Diagnostics::EntityDiagnosticTimingInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Entity", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Time in NS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "Percent of Total", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ECSProfilingDiagnosticsEntityDiagnosticTimingInfo)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.Entity Diagnostics has unexpected decoded type %T", raw)
		}
		p.EntityDiagnostics = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.System Diagnostics", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::SystemDiagnosticTimingInfo", TypeID: "ECS::Profiling::Diagnostics::SystemDiagnosticTimingInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Display Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "System Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 2, Name: "Time in NS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "Percent of Total", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ECSProfilingDiagnosticsSystemDiagnosticTimingInfo)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.System Diagnostics has unexpected decoded type %T", raw)
		}
		p.SystemDiagnostics = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.System Categories", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ECS::Profiling::Diagnostics::SystemCategory", TypeID: "ECS::Profiling::Diagnostics::SystemCategory", Fields: []ShapeField{{Ordinal: 0, Name: "Category Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "System Index", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ECSProfilingDiagnosticsSystemCategory)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.System Categories has unexpected decoded type %T", raw)
		}
		p.SystemCategories = value
	}
	{
		raw, err := r.Read("ServerboundDiagnosticsPacket.Whisker Scopes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "Bedrock::Profile::Whisker::Diagnostics::ScopeDataSummary", TypeID: "Bedrock::Profile::Whisker::Diagnostics::ScopeDataSummary", Fields: []ShapeField{{Ordinal: 0, Name: "Label", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Indentation", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "TotalHighCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 3, Name: "TotalMidCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 4, Name: "TotalLowCostNS", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]BedrockProfileWhiskerDiagnosticsScopeDataSummary)
		if !ok {
			return p, fmt.Errorf("field ServerboundDiagnosticsPacket.Whisker Scopes has unexpected decoded type %T", raw)
		}
		p.WhiskerScopes = value
	}
	return p, nil
}
