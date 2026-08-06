fn main() {}

#[cfg(test)]
mod tests {
    use serde::Deserialize;

    #[derive(Deserialize)]
    struct Golden {
        expected_hex: String,
    }

    fn encode_vocabulary() -> Vec<u8> {
        let mut bytes = Vec::new();

        // nested double optional: true, true, u8(7)
        bytes.extend([1, 1, 7]);
        // two fixed [u16le; 2] values behind an explicit VarUInt32 count
        bytes.extend([2, 1, 0, 3, 2, 4, 0, 6, 5]);
        // discriminated union tag 7, then length-prefixed arbitrary bytes
        bytes.extend([7, 3, 0xa0, 0x00, 0xff]);
        // explicit enum ordinal, reserved u16le, UTF-8 text, u32le semantic ID
        bytes.extend([9, 0x34, 0x12, 2, 0xc3, 0xa9, 4, 3, 2, 1]);
        bytes
    }

    fn read_var_u32(bytes: &[u8], offset: &mut usize) -> u32 {
        let mut value = 0;
        let mut shift = 0;
        loop {
            let byte = bytes[*offset];
            *offset += 1;
            value |= u32::from(byte & 0x7f) << shift;
            if byte & 0x80 == 0 {
                return value;
            }
            shift += 7;
        }
    }

    fn read_u16le(bytes: &[u8], offset: &mut usize) -> u16 {
        let value = u16::from_le_bytes([bytes[*offset], bytes[*offset + 1]]);
        *offset += 2;
        value
    }

    fn read_u32le(bytes: &[u8], offset: &mut usize) -> u32 {
        let value = u32::from_le_bytes([
            bytes[*offset],
            bytes[*offset + 1],
            bytes[*offset + 2],
            bytes[*offset + 3],
        ]);
        *offset += 4;
        value
    }

    fn hex_decode(value: &str) -> Vec<u8> {
        assert_eq!(value.len() % 2, 0);
        value
            .as_bytes()
            .chunks_exact(2)
            .map(|pair| {
                let high = (pair[0] as char).to_digit(16).unwrap();
                let low = (pair[1] as char).to_digit(16).unwrap();
                ((high << 4) | low) as u8
            })
            .collect()
    }

    #[test]
    fn shared_v2_golden_decodes_and_reencodes_identically() {
        let golden: Golden = serde_json::from_str(include_str!(
            "../../../../../testdata/goldens/v2-vocabulary.json"
        ))
        .expect("parse shared v2 golden");
        let expected = hex_decode(&golden.expected_hex);
        let encoded = encode_vocabulary();
        assert_eq!(encoded, expected);

        let mut offset = 0;
        assert_eq!(expected[offset], 1);
        offset += 1;
        assert_eq!(expected[offset], 1);
        offset += 1;
        assert_eq!(expected[offset], 7);
        offset += 1;
        assert_eq!(read_var_u32(&expected, &mut offset), 2);
        assert_eq!(read_u16le(&expected, &mut offset), 1);
        assert_eq!(read_u16le(&expected, &mut offset), 0x0203);
        assert_eq!(read_u16le(&expected, &mut offset), 4);
        assert_eq!(read_u16le(&expected, &mut offset), 0x0506);
        assert_eq!(read_var_u32(&expected, &mut offset), 7);
        assert_eq!(read_var_u32(&expected, &mut offset), 3);
        assert_eq!(&expected[offset..offset + 3], [0xa0, 0x00, 0xff]);
        offset += 3;
        assert_eq!(expected[offset], 9);
        offset += 1;
        assert_eq!(read_u16le(&expected, &mut offset), 0x1234);
        assert_eq!(read_var_u32(&expected, &mut offset), 2);
        assert_eq!(&expected[offset..offset + 2], "é".as_bytes());
        offset += 2;
        assert_eq!(read_u32le(&expected, &mut offset), 0x01020304);
        assert_eq!(offset, expected.len());
    }
}
