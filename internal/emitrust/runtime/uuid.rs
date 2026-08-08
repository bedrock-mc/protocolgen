
// ---------------------------------------------------------------------------
// UUID
// ---------------------------------------------------------------------------

impl Encode for uuid::Uuid {
    fn encode(&self, writer: &mut Writer) {
        // Bedrock transmits a UUID as two little-endian 64-bit halves, not as
        // the raw byte order returned by `Uuid::as_bytes`.
        let (high, low) = self.as_u64_pair();
        writer.write_all(&high.to_le_bytes());
        writer.write_all(&low.to_le_bytes());
    }
}

impl Decode for uuid::Uuid {
    fn decode(reader: &mut Reader<'_>) -> DecodeResult<Self> {
        let high = u64::from_le_bytes(reader.read_bytes::<8>()?);
        let low = u64::from_le_bytes(reader.read_bytes::<8>()?);
        Ok(uuid::Uuid::from_u64_pair(high, low))
    }
}
