// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PurchaseReceipt {
    pub purchase_receipts: Vec<String>,
}

pub const PURCHASERECEIPT_PURCHASE_RECEIPTS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;

impl PurchaseReceipt {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PurchaseReceiptPacket.PurchaseReceipts", PURCHASERECEIPT_PURCHASE_RECEIPTS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PurchaseReceiptPacket.PurchaseReceipts", PURCHASERECEIPT_PURCHASE_RECEIPTS_SHAPE);
    }
}
