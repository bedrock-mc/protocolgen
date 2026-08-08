
// ---------------------------------------------------------------------------
// Vectors
// ---------------------------------------------------------------------------

impl Encode for glam::Vec2 {
    fn encode(&self, writer: &mut Writer) {
        writer.write_all(&self.x.to_le_bytes());
        writer.write_all(&self.y.to_le_bytes());
    }
}

impl Decode for glam::Vec2 {
    fn decode(reader: &mut Reader<'_>) -> DecodeResult<Self> {
        let x = f32::from_le_bytes(reader.read_bytes::<4>()?);
        let y = f32::from_le_bytes(reader.read_bytes::<4>()?);
        Ok(glam::Vec2::new(x, y))
    }
}

impl Encode for glam::Vec3 {
    fn encode(&self, writer: &mut Writer) {
        writer.write_all(&self.x.to_le_bytes());
        writer.write_all(&self.y.to_le_bytes());
        writer.write_all(&self.z.to_le_bytes());
    }
}

impl Decode for glam::Vec3 {
    fn decode(reader: &mut Reader<'_>) -> DecodeResult<Self> {
        let x = f32::from_le_bytes(reader.read_bytes::<4>()?);
        let y = f32::from_le_bytes(reader.read_bytes::<4>()?);
        let z = f32::from_le_bytes(reader.read_bytes::<4>()?);
        Ok(glam::Vec3::new(x, y, z))
    }
}
