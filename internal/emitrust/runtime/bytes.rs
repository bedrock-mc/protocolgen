
// ---------------------------------------------------------------------------
// Length-prefixed byte buffers
// ---------------------------------------------------------------------------

impl Encode for bytes::Bytes {
    fn encode(&self, writer: &mut Writer) {
        writer.write_byte_slice(self);
    }
}

impl Decode for bytes::Bytes {
    fn decode(reader: &mut Reader<'_>) -> DecodeResult<Self> {
        reader.read_byte_slice().map(bytes::Bytes::copy_from_slice)
    }
}
