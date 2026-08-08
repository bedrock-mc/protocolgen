
// ---------------------------------------------------------------------------
// Runtime tests
// ---------------------------------------------------------------------------

#[cfg(test)]
mod runtime_tests {
    use super::*;

    #[test]
    fn varint_rejects_overlong_encodings() {
        // 0x80 0x00 encodes 0 in two bytes; the canonical form is one byte.
        assert_eq!(
            Reader::new(&[0x80, 0x00]).read_var_u32(),
            Err(DecodeError::VarIntOverlong)
        );
        assert_eq!(Reader::new(&[0x00]).read_var_u32(), Ok(0));
    }

    /// A ten-byte varint must not silently drop the bits that do not fit.
    #[test]
    fn varint_rejects_values_wider_than_their_target() {
        let ten_bytes = [0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f];
        assert_eq!(
            Reader::new(&ten_bytes).read_var_u64(),
            Err(DecodeError::VarIntTooLarge)
        );
        let six_bytes = [0xff, 0xff, 0xff, 0xff, 0xff, 0x7f];
        assert_eq!(
            Reader::new(&six_bytes).read_var_u32(),
            Err(DecodeError::VarIntTooLarge)
        );
    }

    #[test]
    fn varint_round_trips_boundary_values() {
        for value in [0u64, 1, 127, 128, u32::MAX as u64, u64::MAX] {
            let mut writer = Writer::new();
            writer.write_var_u64(value);
            assert_eq!(
                Reader::new(writer.as_slice()).read_var_u64(),
                Ok(value),
                "varint {value}"
            );
        }
    }

    #[test]
    fn zigzag_round_trips_boundary_values() {
        for value in [0i32, -1, 1, i32::MIN, i32::MAX] {
            let mut writer = Writer::new();
            writer.write_zigzag_i32(value);
            assert_eq!(
                Reader::new(writer.as_slice()).read_zigzag_i32(),
                Ok(value),
                "zigzag {value}"
            );
        }
    }

    /// A declared count is rejected before it is used to reserve memory.
    #[test]
    fn declared_counts_are_bounded_before_allocation() {
        let reader = Reader::new(&[0u8; 4]);
        assert_eq!(
            reader.checked_count(u64::from(u32::MAX), 1, MAX_COLLECTION_ELEMENTS),
            Err(DecodeError::LengthLimitExceeded {
                limit: MAX_COLLECTION_ELEMENTS,
                actual: u32::MAX as usize,
            })
        );
        // Within the limit, but more elements than the input could hold.
        assert_eq!(
            reader.checked_count(100, 8, MAX_COLLECTION_ELEMENTS),
            Err(DecodeError::LengthNotRepresentable { needed: 800, remaining: 4 })
        );
        assert_eq!(reader.checked_count(4, 1, MAX_COLLECTION_ELEMENTS), Ok(4));
    }

    #[test]
    fn strings_reject_invalid_utf8() {
        assert_eq!(
            Reader::new(&[0x01, 0xff]).read_str(),
            Err(DecodeError::InvalidUtf8)
        );
    }

    #[test]
    fn trailing_bytes_are_rejected() {
        assert_eq!(
            Reader::new(&[0x00, 0x01]).expect_consumed(),
            Err(DecodeError::TrailingBytes(2))
        );
    }
}

#[cfg(test)]
mod bitset_tests {
    use super::*;

    #[test]
    fn bitsets_round_trip() {
        for words in [[0u64, 0, 0], [1, 0, 0], [0, 0, 1 << 2], [u64::MAX, u64::MAX, 0b111]] {
            let mut writer = Writer::new();
            encode_bitset(&mut writer, &words, 131);
            let decoded: [u64; 3] =
                decode_bitset(&mut Reader::new(writer.as_slice()), 131).expect("decode");
            assert_eq!(decoded, words, "bitset {words:?}");
        }
    }

    #[test]
    fn empty_bitset_is_one_byte() {
        let mut writer = Writer::new();
        encode_bitset(&mut writer, &[0u64; 3], 131);
        assert_eq!(writer.as_slice(), &[0x00]);
    }
}
