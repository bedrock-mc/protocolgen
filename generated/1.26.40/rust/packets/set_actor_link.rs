// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetActorLink {
    pub link: ActorLink,
}

pub const SETACTORLINK_LINK_SHAPE: &str = r##"{"kind":"struct","semantic":"ActorLink","type_id":"ActorLink","fields":[{"ordinal":0,"name":"Target A","semantic":"Target A","type_id":"ActorUniqueID.json#","encode":{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":1,"name":"Target B","semantic":"Target B","type_id":"ActorUniqueID.json#","encode":{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":2,"name":"Type","semantic":"Type","encode":{"kind":"enum","semantic":"ActorLinkType","type_id":"enums/ActorLinkType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"Riding","encode":{"kind":"void"}},{"value":2,"name":"Passenger","encode":{"kind":"void"}}]},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":3,"name":"Immediate","semantic":"Immediate","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":4,"name":"Passenger Initiated","semantic":"Passenger Initiated","encode":{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}},{"ordinal":5,"name":"Vehicle Angular Velocity","semantic":"Vehicle Angular Velocity","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone","mojang"]}}]}"##;

impl SetActorLink {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetActorLinkPacket.Link", SETACTORLINK_LINK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetActorLinkPacket.Link", SETACTORLINK_LINK_SHAPE);
    }
}
