package emitrust

import (
	"fmt"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func contractFixture() manifest.Manifest {
	field := func(ordinal int, name string, node manifest.Node) manifest.Field {
		return manifest.Field{Ordinal: ordinal, Name: name, Encode: node, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}
	}
	prefix := manifest.Primitive("var_u32")
	bounded := manifest.String(prefix)
	bounded.Constraints = &manifest.Constraints{MaxLength: pointerTo(uint64(64))}
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets: []manifest.Packet{{ID: 9, Name: "ChatPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			field(0, "Message", manifest.String(prefix)),
			field(1, "Parameters", manifest.Array(prefix, manifest.String(prefix))),
			field(2, "Filtered", manifest.Optional(manifest.String(prefix))),
			field(3, "Payload", manifest.Bytes(prefix)),
			field(4, "Counts", manifest.Array(prefix, manifest.Primitive("u16le"))),
			field(5, "Tags", manifest.Map(prefix, manifest.String(prefix), manifest.Primitive("u8"))),
			field(6, "Title", bounded),
		}}},
	}
}

func TestContractOutputIsOptIn(t *testing.T) {
	files, err := GenerateFiles(contractFixture())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["src/fields.rs"]; ok {
		t.Fatal("profile-neutral output emitted contract field sites")
	}
	for name, source := range files {
		for _, marker := range []string{"EncodeResult", "crate::fields::", "wire::Str", "decode_frame"} {
			if strings.Contains(source, marker) {
				t.Fatalf("profile-neutral %s contains contract marker %q", name, marker)
			}
		}
	}
}

func TestContractRejectsUnknownVersions(t *testing.T) {
	for _, version := range []int{-1, 1, CodecContractVersion + 1} {
		if _, err := GenerateFilesWithOptions(contractFixture(), Options{CodecContract: version}); err == nil {
			t.Fatalf("unimplemented codec contract version %d was accepted", version)
		}
	}
}

// Every length-prefixed site gets a FieldId, including each string element,
// map key and optional inner value, so a consumer can bound one field alone.
func TestContractNamesEveryLengthPrefixedSite(t *testing.T) {
	files, err := GenerateFilesWithOptions(contractFixture(), Options{CodecContract: CodecContractVersion})
	if err != nil {
		t.Fatal(err)
	}
	fields := files["src/fields.rs"]
	for _, path := range []string{"Chat.message", "Chat.parameters", "Chat.parameters[]", "Chat.filtered", "Chat.payload", "Chat.counts", "Chat.tags", "Chat.tags{key}", "Chat.title"} {
		if !strings.Contains(fields, `"`+path+`"`) {
			t.Fatalf("fields.rs omits site %q:\n%s", path, fields)
		}
	}
	packets := files["src/packets.rs"]
	for _, want := range []string{
		"reader.read_str(crate::fields::CHAT_MESSAGE)?",
		"reader.checked_count(crate::fields::CHAT_PARAMETERS, declared, 1)?",
		"wire::alloc_vec(count)?",
		"wire::push(&mut out, reader.read_str(crate::fields::CHAT_PARAMETERS_ITEM)?)?",
		"reader.read_shared_bytes(crate::fields::CHAT_PAYLOAD)?",
		"wire::decode_collection::<wire::U16LE>(reader, crate::fields::CHAT_COUNTS, 2)?",
		"wire::decode_string_limits(reader, crate::fields::CHAT_TITLE, 0, 64)?",
		"pub message: wire::Str,",
	} {
		if !strings.Contains(packets, want) {
			t.Fatalf("contract packets omit %q:\n%s", want, packets)
		}
	}
}

func TestContractEncodingIsFallible(t *testing.T) {
	files, err := GenerateFilesWithOptions(contractFixture(), Options{CodecContract: CodecContractVersion})
	if err != nil {
		t.Fatal(err)
	}
	packets := files["src/packets.rs"]
	for _, want := range []string{
		"fn encode(&self, writer: &mut wire::Writer<'_>) -> wire::EncodeResult<()>",
		"self.message.encode(writer)?;",
		"writer.write_len_var_u32(self.parameters.len())?;",
		"None => writer.write_u8(0)?,",
		"wire::encode_string_limits(writer, &self.title, 0, 64)?;",
		"        Ok(())\n",
	} {
		if !strings.Contains(packets, want) {
			t.Fatalf("contract encode omits %q:\n%s", want, packets)
		}
	}
	if strings.Contains(packets, "assert_") {
		t.Fatalf("contract encode still panics on schema violations:\n%s", packets)
	}
}

func TestContractFramesPreserveUnknownPackets(t *testing.T) {
	files, err := GenerateFilesWithOptions(contractFixture(), Options{CodecContract: CodecContractVersion})
	if err != nil {
		t.Fatal(err)
	}
	packets := files["src/packets.rs"]
	for _, want := range []string{"pub enum Decoded", "Unknown(UnknownPacket)", "pub fn decode_frame(", "pub fn encode_frame(", "use crate::wire::{Direction, Peer};"} {
		if !strings.Contains(packets, want) {
			t.Fatalf("contract packets omit %q", want)
		}
	}
	if strings.Contains(packets, "pub enum Peer") {
		t.Fatal("contract packets redefine Peer instead of using the wire runtime")
	}
	lib := files["src/lib.rs"]
	for _, want := range []string{fmt.Sprintf("pub const CODEC_CONTRACT: u32 = %d;", CodecContractVersion), "pub const MANIFEST_SHA256: &str = \"", "pub mod fields;"} {
		if !strings.Contains(lib, want) {
			t.Fatalf("contract lib.rs omits %q:\n%s", want, lib)
		}
	}
	if wire := files["src/wire.rs"]; !strings.Contains(wire, "pub struct Frames") || !strings.Contains(wire, "pub struct Limits") {
		t.Fatal("contract runtime omits frames or limits")
	}
	roundtrip := files["tests/roundtrip.rs"]
	for _, want := range []string{"unknown_packets_keep_exact_boundaries", "unknown_packets_forward_byte_for_byte"} {
		if !strings.Contains(roundtrip, want) {
			t.Fatalf("contract tests omit %s", want)
		}
	}
	// Re-encoding from the parsed header would drop bits above the target
	// subclient, so forwarding must copy the received entry.
	if !strings.Contains(packets, "wire::Writer::new(out).write_all(&self.frame)") {
		t.Fatal("unknown packets are not forwarded from their received bytes")
	}
}
