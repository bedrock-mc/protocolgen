// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PartyChanged {
    pub party_info: Option<PlayerPartyInfo>,
}

pub const PARTYCHANGED_PARTY_INFO_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"PlayerPartyInfo","type_id":"PlayerPartyInfo","fields":[{"ordinal":0,"name":"party_id","semantic":"party_id","encode":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"is_party_leader","semantic":"is_party_leader","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl PartyChanged {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PartyChangedPacket.party_info", PARTYCHANGED_PARTY_INFO_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PartyChangedPacket.party_info", PARTYCHANGED_PARTY_INFO_SHAPE);
    }
}
