
// ---------------------------------------------------------------------------
// NBT
// ---------------------------------------------------------------------------

/// Integer and length encoding used by an NBT payload.
///
/// Bedrock uses both forms and they are not interchangeable: the network form
/// varint-encodes integers and lengths, the persistent form stores them as
/// fixed-width little endian. A payload read with the wrong variant desyncs the
/// surrounding packet rather than failing locally.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum NbtVariant {
    Network,
    Persistent,
}

const TAG_END: u8 = 0;
const TAG_BYTE: u8 = 1;
const TAG_SHORT: u8 = 2;
const TAG_INT: u8 = 3;
const TAG_LONG: u8 = 4;
const TAG_FLOAT: u8 = 5;
const TAG_DOUBLE: u8 = 6;
const TAG_BYTE_ARRAY: u8 = 7;
const TAG_STRING: u8 = 8;
const TAG_LIST: u8 = 9;
const TAG_COMPOUND: u8 = 10;
const TAG_INT_ARRAY: u8 = 11;
const TAG_LONG_ARRAY: u8 = 12;

impl<'a> Reader<'a> {
    fn nbt_len(&mut self, variant: NbtVariant) -> DecodeResult<usize> {
        let declared = match variant {
            NbtVariant::Network => i64::from(self.read_zigzag_i32()?),
            NbtVariant::Persistent => i64::from(i32::from_le_bytes(self.read_bytes::<4>()?)),
        };
        if declared < 0 {
            return Err(DecodeError::NegativeLength(declared));
        }
        self.checked_count(declared as u64, 1, MAX_BYTE_BUFFER_LEN)
    }

    fn nbt_name(&mut self, variant: NbtVariant) -> DecodeResult<()> {
        let count = match variant {
            NbtVariant::Network => {
                let declared = u64::from(self.read_var_u32()?);
                self.checked_count(declared, 1, MAX_BYTE_BUFFER_LEN)?
            }
            NbtVariant::Persistent => usize::from(u16::from_le_bytes(self.read_bytes::<2>()?)),
        };
        self.take(count)?;
        Ok(())
    }

    fn scan_nbt_payload(&mut self, tag: u8, variant: NbtVariant, depth: u32) -> DecodeResult<()> {
        if depth > MAX_NBT_DEPTH {
            return Err(DecodeError::NbtDepthExceeded);
        }
        match tag {
            TAG_BYTE => {
                self.take(1)?;
            }
            TAG_SHORT => {
                self.take(2)?;
            }
            TAG_INT => match variant {
                NbtVariant::Network => {
                    self.read_zigzag_i32()?;
                }
                NbtVariant::Persistent => {
                    self.take(4)?;
                }
            },
            TAG_LONG => match variant {
                NbtVariant::Network => {
                    self.read_zigzag_i64()?;
                }
                NbtVariant::Persistent => {
                    self.take(8)?;
                }
            },
            TAG_FLOAT => {
                self.take(4)?;
            }
            TAG_DOUBLE => {
                self.take(8)?;
            }
            TAG_BYTE_ARRAY => {
                let count = self.nbt_len(variant)?;
                self.take(count)?;
            }
            TAG_STRING => self.nbt_name(variant)?,
            TAG_LIST => {
                let element = self.read_u8()?;
                let count = self.nbt_len(variant)?;
                if element == TAG_END && count != 0 {
                    return Err(DecodeError::NbtMalformed("non-empty list of TAG_End"));
                }
                for _ in 0..count {
                    self.scan_nbt_payload(element, variant, depth + 1)?;
                }
            }
            TAG_COMPOUND => loop {
                let entry = self.read_u8()?;
                if entry == TAG_END {
                    break;
                }
                self.nbt_name(variant)?;
                self.scan_nbt_payload(entry, variant, depth + 1)?;
            },
            TAG_INT_ARRAY => {
                let count = self.nbt_len(variant)?;
                for _ in 0..count {
                    match variant {
                        NbtVariant::Network => {
                            self.read_zigzag_i32()?;
                        }
                        NbtVariant::Persistent => {
                            self.take(4)?;
                        }
                    }
                }
            }
            TAG_LONG_ARRAY => {
                let count = self.nbt_len(variant)?;
                for _ in 0..count {
                    match variant {
                        NbtVariant::Network => {
                            self.read_zigzag_i64()?;
                        }
                        NbtVariant::Persistent => {
                            self.take(8)?;
                        }
                    }
                }
            }
            _ => return Err(DecodeError::NbtMalformed("unknown tag")),
        }
        Ok(())
    }

    /// Advances past exactly one root NBT value and borrows its bytes.
    pub fn read_nbt(&mut self, variant: NbtVariant) -> DecodeResult<&'a [u8]> {
        let start = self.pos;
        let tag = self.read_u8()?;
        if tag != TAG_END {
            self.nbt_name(variant)?;
            self.scan_nbt_payload(tag, variant, 0)?;
        }
        Ok(&self.buf[start..self.pos])
    }
}

macro_rules! nbt_codec {
    ($name:ident, $variant:expr, $doc:expr) => {
        #[doc = $doc]
        #[derive(Clone, Debug, PartialEq, Eq)]
        pub struct $name(pub bytes::Bytes);

        impl Default for $name {
            /// A lone TAG_End, which is the empty payload the scanner accepts.
            fn default() -> Self {
                Self(bytes::Bytes::from_static(&[0u8]))
            }
        }

        impl $name {
            pub const VARIANT: NbtVariant = $variant;
        }

        impl Encode for $name {
            fn encode(&self, writer: &mut Writer) {
                writer.write_all(&self.0);
            }
        }

        impl Decode for $name {
            fn decode(reader: &mut Reader<'_>) -> DecodeResult<Self> {
                let bytes = reader.read_nbt($variant)?;
                Ok(Self(bytes::Bytes::copy_from_slice(bytes)))
            }
        }
    };
}

nbt_codec!(
    NetworkNbt,
    NbtVariant::Network,
    "NBT with varint integers and lengths."
);
nbt_codec!(
    PersistentNbt,
    NbtVariant::Persistent,
    "NBT with fixed-width little-endian integers and lengths."
);

#[cfg(test)]
mod nbt_tests {
    use super::*;

    /// The same bytes mean different things under the two variants, so the
    /// scanner must consume a different number of them.
    #[test]
    fn network_and_persistent_integers_differ() {
        // TAG_Int("", 1): network varint-encodes the value, persistent uses i32le.
        let network = [0x03u8, 0x00, 0x02];
        let mut reader = Reader::new(&network);
        assert_eq!(reader.read_nbt(NbtVariant::Network), Ok(&network[..]));

        let persistent = [0x03u8, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00];
        let mut reader = Reader::new(&persistent);
        assert_eq!(reader.read_nbt(NbtVariant::Persistent), Ok(&persistent[..]));

        // Reading the persistent payload as network leaves bytes unconsumed.
        let mut reader = Reader::new(&persistent);
        reader.read_nbt(NbtVariant::Network).expect("scan");
        assert_ne!(reader.remaining(), 0);
    }

    #[test]
    fn empty_root_is_a_single_end_tag() {
        let mut reader = Reader::new(&[0x00]);
        assert_eq!(reader.read_nbt(NbtVariant::Network), Ok(&[0x00][..]));
        assert_eq!(reader.remaining(), 0);
    }

    #[test]
    fn nesting_is_bounded() {
        // Root compound with an empty name, then a chain of compounds nested
        // deeper than the cap. Each entry is TAG_Compound plus a zero-length
        // name; the scanner errors before reaching the closing TAG_Ends.
        let depth = MAX_NBT_DEPTH + 8;
        let mut deep = vec![0x0a, 0x00];
        for _ in 0..depth {
            deep.extend_from_slice(&[0x0a, 0x00]);
        }
        deep.extend(std::iter::repeat_n(0x00, depth as usize + 1));
        let mut reader = Reader::new(&deep);
        assert_eq!(
            reader.read_nbt(NbtVariant::Network),
            Err(DecodeError::NbtDepthExceeded)
        );
    }

    #[test]
    fn truncated_payloads_fail_rather_than_over_read() {
        let mut reader = Reader::new(&[0x03, 0x00]);
        assert!(reader.read_nbt(NbtVariant::Persistent).is_err());
    }
}
