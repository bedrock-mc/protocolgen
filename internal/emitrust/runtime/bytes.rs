
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
        let declared = u64::from(reader.read_var_u32()?);
        let count = reader.checked_count_with(declared, 1, reader.byte_buffer_limit())?;
        reader.take_shared(count)
    }
}
