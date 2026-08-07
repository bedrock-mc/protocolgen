// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ShowStoreOffer {
    pub offer_id: uuid::Uuid,
    pub redirect_type: ShowStoreOfferRedirectType,
}

pub const SHOWSTOREOFFER_OFFER_ID_SHAPE: &str = r##"{"kind":"primitive","semantic":"mce::UUID","type_id":"mce__UUID.json#","primitive":{"code":"uuid","width":128,"signed":false,"zigzag":false,"endianness":"none"}}"##;
pub const SHOWSTOREOFFER_REDIRECT_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"ShowStoreOfferRedirectType","type_id":"enums/ShowStoreOfferRedirectType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"MarketplaceOffer","encode":{"kind":"void"}},{"value":1,"name":"DressingRoomOffer","encode":{"kind":"void"}},{"value":2,"name":"ThirdPartyServerPage","encode":{"kind":"void"}}]}"#;

impl ShowStoreOffer {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ShowStoreOfferPacket.Offer Id", SHOWSTOREOFFER_OFFER_ID_SHAPE);
        encoder.field("ShowStoreOfferPacket.Redirect Type", SHOWSTOREOFFER_REDIRECT_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ShowStoreOfferPacket.Offer Id", SHOWSTOREOFFER_OFFER_ID_SHAPE);
        decoder.field("ShowStoreOfferPacket.Redirect Type", SHOWSTOREOFFER_REDIRECT_TYPE_SHAPE);
    }
}
