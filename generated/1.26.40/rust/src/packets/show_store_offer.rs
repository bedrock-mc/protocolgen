// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ShowStoreOffer {
    pub offer_id: uuid::Uuid,
    pub redirect_type: ShowStoreOfferRedirectType,
}

impl ShowStoreOffer {
    pub const ID: u32 = 91;
}
