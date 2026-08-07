// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use std::io::{self, Read, Write};

pub trait WireCodec: Sized {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()>;
    fn decode<R: Read>(reader: &mut R) -> io::Result<Self>;
}

macro_rules! fixed_codec {
    ($name:ident, $inner:ty, $size:expr, $write:ident, $read:ident) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                writer.write_all(&self.0.$write())
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                let mut bytes = [0u8; $size];
                reader.read_exact(&mut bytes)?;
                Ok(Self(<$inner>::$read(bytes)))
            }
        }
    };
}

macro_rules! fixed_float_codec {
    ($name:ident, $inner:ty, $size:expr, $write:ident, $read:ident) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                writer.write_all(&self.0.$write())
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                let mut bytes = [0u8; $size];
                reader.read_exact(&mut bytes)?;
                Ok(Self(<$inner>::$read(bytes)))
            }
        }
    };
}

fixed_codec!(I8, i8, 1, to_ne_bytes, from_ne_bytes);
fixed_codec!(U8, u8, 1, to_ne_bytes, from_ne_bytes);
fixed_codec!(I16LE, i16, 2, to_le_bytes, from_le_bytes);
fixed_codec!(I16BE, i16, 2, to_be_bytes, from_be_bytes);
fixed_codec!(U16LE, u16, 2, to_le_bytes, from_le_bytes);
fixed_codec!(U16BE, u16, 2, to_be_bytes, from_be_bytes);
fixed_codec!(I32LE, i32, 4, to_le_bytes, from_le_bytes);
fixed_codec!(I32BE, i32, 4, to_be_bytes, from_be_bytes);
fixed_codec!(U32LE, u32, 4, to_le_bytes, from_le_bytes);
fixed_codec!(U32BE, u32, 4, to_be_bytes, from_be_bytes);
fixed_codec!(I64LE, i64, 8, to_le_bytes, from_le_bytes);
fixed_codec!(I64BE, i64, 8, to_be_bytes, from_be_bytes);
fixed_codec!(U64LE, u64, 8, to_le_bytes, from_le_bytes);
fixed_codec!(U64BE, u64, 8, to_be_bytes, from_be_bytes);
fixed_float_codec!(F32LE, f32, 4, to_le_bytes, from_le_bytes);
fixed_float_codec!(F32BE, f32, 4, to_be_bytes, from_be_bytes);
fixed_float_codec!(F64LE, f64, 8, to_le_bytes, from_le_bytes);
fixed_float_codec!(F64BE, f64, 8, to_be_bytes, from_be_bytes);

fn write_var_u64<W: Write>(writer: &mut W, mut value: u64) -> io::Result<()> {
    while value >= 0x80 {
        writer.write_all(&[(value as u8) | 0x80])?;
        value >>= 7;
    }
    writer.write_all(&[value as u8])
}

fn read_var_u64<R: Read>(reader: &mut R) -> io::Result<u64> {
    let mut value = 0u64;
    for shift in (0..64).step_by(7) {
        let mut byte = [0u8; 1];
        reader.read_exact(&mut byte)?;
        value |= u64::from(byte[0] & 0x7f) << shift;
        if byte[0] & 0x80 == 0 {
            return Ok(value);
        }
    }
    Err(io::Error::new(
        io::ErrorKind::InvalidData,
        "varint too large",
    ))
}

macro_rules! var_codec {
    ($name:ident, $inner:ty) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                write_var_u64(writer, self.0 as u64)
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                Ok(Self(read_var_u64(reader)? as $inner))
            }
        }
    };
}

var_codec!(VarInt, i32);
var_codec!(VarUInt, u32);
var_codec!(VarLong, i64);
var_codec!(VarULong, u64);

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ZigZag32(pub i32);

impl WireCodec for ZigZag32 {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
        write_var_u64(writer, ((self.0 << 1) ^ (self.0 >> 31)) as u32 as u64)
    }

    fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
        let raw = read_var_u64(reader)? as u32;
        Ok(Self(((raw >> 1) as i32) ^ -((raw & 1) as i32)))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ZigZag64(pub i64);

impl WireCodec for ZigZag64 {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
        write_var_u64(writer, ((self.0 << 1) ^ (self.0 >> 63)) as u64)
    }

    fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
        let raw = read_var_u64(reader)?;
        Ok(Self(((raw >> 1) as i64) ^ -((raw & 1) as i64)))
    }
}
