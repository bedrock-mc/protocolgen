// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct InventoryTransaction {
    pub legacy_request_id: TypedClientNetIdStructItemStackLegacyRequestIdTagInt32T0,
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    pub transaction: Option<InventoryTransactionTransactionValue>,
}

impl InventoryTransaction {
    pub const ID: u32 = 30;
}
