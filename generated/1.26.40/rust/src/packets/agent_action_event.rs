// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AgentActionEvent {
    pub request_id: String,
    pub action: AgentActionType,
    pub response: String,
}

impl AgentActionEvent {
    pub const ID: u32 = 181;
}
