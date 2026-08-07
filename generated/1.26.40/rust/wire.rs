// Code generated from canonical protocol manifest v2. DO NOT EDIT.

pub trait WireEncoder {
    fn field(&mut self, path: &'static str, shape: &'static str);
}

pub trait WireDecoder {
    fn field(&mut self, path: &'static str, shape: &'static str);
}
