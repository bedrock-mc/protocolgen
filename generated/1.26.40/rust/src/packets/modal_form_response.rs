// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ModalFormResponse {
    pub form_id: u32,
    pub json_response: Option<String>,
    pub form_cancel_reason: Option<ModalFormCancelReason>,
}

impl ModalFormResponse {
    pub const ID: u32 = 101;
}
