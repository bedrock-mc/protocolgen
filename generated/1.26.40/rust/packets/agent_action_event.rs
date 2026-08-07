// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AgentActionEvent {
    pub request_id: String,
    pub action: AgentActionType,
    pub response: String,
}

pub const AGENTACTIONEVENT_REQUEST_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const AGENTACTIONEVENT_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"AgentActionType","type_id":"enums/AgentActionType","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"},"variants":[{"value":1,"name":"Attack","encode":{"kind":"void"}},{"value":2,"name":"Collect","encode":{"kind":"void"}},{"value":3,"name":"Destroy","encode":{"kind":"void"}},{"value":4,"name":"DetectRedstone","encode":{"kind":"void"}},{"value":5,"name":"DetectObstacle","encode":{"kind":"void"}},{"value":6,"name":"Drop","encode":{"kind":"void"}},{"value":7,"name":"DropAll","encode":{"kind":"void"}},{"value":8,"name":"Inspect","encode":{"kind":"void"}},{"value":9,"name":"InspectData","encode":{"kind":"void"}},{"value":10,"name":"InspectItemCount","encode":{"kind":"void"}},{"value":11,"name":"InspectItemDetail","encode":{"kind":"void"}},{"value":12,"name":"InspectItemSpace","encode":{"kind":"void"}},{"value":13,"name":"Interact","encode":{"kind":"void"}},{"value":14,"name":"Move","encode":{"kind":"void"}},{"value":15,"name":"PlaceBlock","encode":{"kind":"void"}},{"value":16,"name":"Till","encode":{"kind":"void"}},{"value":17,"name":"TransferItemTo","encode":{"kind":"void"}},{"value":18,"name":"Turn","encode":{"kind":"void"}}]}"#;
pub const AGENTACTIONEVENT_RESPONSE_SHAPE: &str = r##"{"kind":"string","semantic":"Json::Value","type_id":"Json__Value.json#","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"##;

impl AgentActionEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AgentActionEventPacket.Request Id", AGENTACTIONEVENT_REQUEST_ID_SHAPE);
        encoder.field("AgentActionEventPacket.Action", AGENTACTIONEVENT_ACTION_SHAPE);
        encoder.field("AgentActionEventPacket.Response", AGENTACTIONEVENT_RESPONSE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AgentActionEventPacket.Request Id", AGENTACTIONEVENT_REQUEST_ID_SHAPE);
        decoder.field("AgentActionEventPacket.Action", AGENTACTIONEVENT_ACTION_SHAPE);
        decoder.field("AgentActionEventPacket.Response", AGENTACTIONEVENT_RESPONSE_SHAPE);
    }
}
