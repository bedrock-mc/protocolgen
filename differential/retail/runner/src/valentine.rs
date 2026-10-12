//! Cinnabar's current codec: the vendored Valentine crate at the corpus
//! commit, framed and decoded as `crates/protocol/src/codec.rs`'s
//! `decode_batch` does, without Cinnabar's UI and audio policy checks.

use bytes::{Buf, Bytes, BytesMut};
use valentine::bedrock::error::DecodeError;
use valentine::bedrock::version::v1_26_51::{McpePacketArgs, McpePacketData};
use valentine::protocol::wire;

use crate::verdict::{DIFFERS, Input, OK, SAME, Verdict, digest, guarded, header_text};

const BATCH_HEADER: u8 = 0xfe;

pub fn run(input: &Input, lines: &mut Vec<Verdict>) {
    let case = input.case.as_str();
    if input.batch.first() != Some(&BATCH_HEADER) {
        return lines.push(Verdict::frame_error(
            input,
            case,
            0,
            "missing the 0xfe batch header",
        ));
    }
    let mut rest = input.batch.slice(1..);
    let mut entry = 0;
    while rest.has_remaining() {
        let start = rest.clone();
        let before = rest.remaining();
        let declared = match wire::read_var_u32(&mut rest) {
            Ok(declared) => declared as usize,
            Err(error) => return lines.push(Verdict::frame_error(input, case, entry, error)),
        };
        let prefix = before - rest.remaining();
        if rest.remaining() < declared {
            let detail = format!(
                "entry declares {declared} bytes, {} remain",
                rest.remaining()
            );
            return lines.push(Verdict::frame_error(input, case, entry, detail));
        }
        let frame = start.slice(..prefix + declared);
        rest.advance(declared);
        let mut verdict = Verdict::new(input, case, entry);
        let mut payload = frame.slice(prefix..);
        verdict.payload = Some(digest(&payload));
        if let Ok(raw) = wire::read_var_u32(&mut payload) {
            verdict.header = Some(header_text(raw));
        }
        guarded(&mut verdict, |verdict| decode(frame, verdict));
        lines.push(verdict);
        entry += 1;
    }
}

fn decode(frame: Bytes, verdict: &mut Verdict) {
    let mut cursor = frame;
    let (header, data) = match McpePacketData::decode_inner(&mut cursor, McpePacketArgs) {
        Ok(decoded) => decoded,
        Err(error) => return verdict.fail(class(&error), &error),
    };
    if cursor.has_remaining() {
        let detail = format!("{} unread bytes in the declared entry", cursor.remaining());
        return verdict.fail("trailing", detail);
    }
    verdict.outcome = OK;
    let fields = format!("{data:?}");
    verdict.fields = Some(digest(fields.as_bytes()));
    let mut out = BytesMut::new();
    if let Err(error) =
        data.encode_inner_bytes_mut(&mut out, header.from_subclient, header.to_subclient)
    {
        verdict.detail = Some(format!("encode: {error}"));
        return;
    }
    let entry = out.freeze();
    let mut payload = entry.clone();
    if wire::read_var_u32(&mut payload).is_err() {
        verdict.detail = Some("the re-encoded entry has no length".to_string());
        return;
    }
    verdict.encoded = Some(digest(&payload));
    let mut again = entry;
    verdict.roundtrip = Some(
        match McpePacketData::decode_inner(&mut again, McpePacketArgs) {
            Ok((_, data)) if !again.has_remaining() && format!("{data:?}") == fields => SAME,
            _ => DIFFERS,
        },
    );
}

fn class(error: &DecodeError) -> &'static str {
    match error {
        DecodeError::UnexpectedEof { .. }
        | DecodeError::StringLengthExceeded { .. }
        | DecodeError::ArrayLengthExceeded { .. }
        | DecodeError::PacketLengthExceeded { .. } => "eof",
        DecodeError::NegativeLength { .. } => "length",
        DecodeError::InvalidEnumValue { .. } => "variant",
        DecodeError::InvalidPacketId { .. } => "unknown-id",
        DecodeError::InvalidMagicByte { .. } => "frame",
        DecodeError::VarIntTooLarge | DecodeError::VarLongTooLarge => "varint",
        DecodeError::InvalidBitset { .. } => "bitset",
        DecodeError::UnknownNbtTag { .. } | DecodeError::NbtDepthExceeded { .. } => "nbt",
        DecodeError::Utf8Error(_) => "utf8",
        DecodeError::Io(_) => "io",
    }
}
