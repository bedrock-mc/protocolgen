// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerboundDiagnostics {
    pub avg_fps: f32,
    pub avg_server_sim_tick_time_ms: f32,
    pub avg_client_sim_tick_time_ms: f32,
    pub avg_begin_frame_time_ms: f32,
    pub avg_input_time_ms: f32,
    pub avg_render_time_ms: f32,
    pub avg_end_frame_time_ms: f32,
    pub avg_remainder_time_percent: f32,
    pub avg_unaccounted_time_percent: f32,
    pub memory_category_values: Vec<MemoryMemoryCategoryCounter>,
    pub entity_diagnostics: Vec<ECSProfilingDiagnosticsEntityDiagnosticTimingInfo>,
    pub system_diagnostics: Vec<ECSProfilingDiagnosticsSystemDiagnosticTimingInfo>,
    pub system_categories: Vec<ECSProfilingDiagnosticsSystemCategory>,
    pub whisker_scopes: Vec<BedrockProfileWhiskerDiagnosticsScopeDataSummary>,
}

impl ServerboundDiagnostics {
    pub const ID: u32 = 315;
}
