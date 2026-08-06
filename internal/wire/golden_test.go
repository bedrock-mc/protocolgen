package wire

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"protocolgen/internal/manifest"
)

func TestVocabularyGoldenRoundTripsBothDirections(t *testing.T) {
	var golden struct {
		ExpectedHex string `json:"expected_hex"`
	}
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "goldens", "v2-vocabulary.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &golden); err != nil {
		t.Fatal(err)
	}
	node := vocabularyNode()
	value := map[string]any{
		"Outer":     OptionalValue{Present: true, Value: OptionalValue{Present: true, Value: uint8(7)}},
		"Values":    []any{[]any{uint16(1), uint16(0x0203)}, []any{uint16(4), uint16(0x0506)}},
		"Choice":    UnionValue{Tag: 7, Value: []byte{0xa0, 0x00, 0xff}},
		"Mode":      uint8(9),
		"Reserved":  uint16(0x1234),
		"Text":      "é",
		"NetworkID": uint32(0x01020304),
	}
	encoded, err := Encode(node, value)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	if got := hex.EncodeToString(encoded); got != golden.ExpectedHex {
		t.Fatalf("encoded = %s, want %s", got, golden.ExpectedHex)
	}
	decoded, err := DecodeAll(node, encoded)
	if err != nil {
		t.Fatalf("DecodeAll: %v", err)
	}
	decodedObject, ok := decoded.(map[string]any)
	if !ok {
		t.Fatalf("decoded value = %#v, want object", decoded)
	}
	if !reflect.DeepEqual(decodedObject["Text"], "é") ||
		!reflect.DeepEqual(decodedObject["NetworkID"], uint32(0x01020304)) ||
		!reflect.DeepEqual(decodedObject["Mode"], uint8(9)) ||
		!reflect.DeepEqual(decodedObject["Reserved"], uint16(0x1234)) {
		t.Fatalf("decoded semantic values = %#v", decodedObject)
	}
	if optional, ok := decodedObject["Outer"].(OptionalValue); !ok || !optional.Present {
		t.Fatalf("decoded nested optional = %#v", decodedObject["Outer"])
	}
	decodedBytes, err := Encode(node, decodedObject)
	if err != nil {
		t.Fatalf("re-encode decoded golden: %v", err)
	}
	if !bytes.Equal(encoded, decodedBytes) {
		t.Fatalf("decoded/re-encoded bytes = %x, want %x", decodedBytes, encoded)
	}
	semanticEquivalent := vocabularyNode()
	semanticEquivalent.Fields[len(semanticEquivalent.Fields)-1].Semantic = "DifferentAPIIdentity"
	encodedEquivalent, err := Encode(semanticEquivalent, value)
	if err != nil {
		t.Fatalf("Encode semantic-equivalent node: %v", err)
	}
	if !reflect.DeepEqual(encoded, encodedEquivalent) {
		t.Fatalf("semantic identity changed wire bytes: %x != %x", encoded, encodedEquivalent)
	}
}

func TestSignedVarintAndZigZagHaveDistinctWireBytes(t *testing.T) {
	varint, err := Encode(manifest.Primitive("var_i32"), int32(-1))
	if err != nil {
		t.Fatal(err)
	}
	zigzag, err := Encode(manifest.Primitive("zigzag_i32"), int32(-1))
	if err != nil {
		t.Fatal(err)
	}
	if string(varint) == string(zigzag) || len(zigzag) != 1 || zigzag[0] != 1 {
		t.Fatalf("varint=%x zigzag=%x", varint, zigzag)
	}
	decoded, err := DecodeAll(manifest.Primitive("var_i32"), varint)
	if err != nil || decoded != int32(-1) {
		t.Fatalf("decoded varint=%v err=%v", decoded, err)
	}
	decoded, err = DecodeAll(manifest.Primitive("zigzag_i32"), zigzag)
	if err != nil || decoded != int32(-1) {
		t.Fatalf("decoded zigzag=%v err=%v", decoded, err)
	}
}

func vocabularyNode() manifest.Node {
	return manifest.Struct(
		manifest.Field{Ordinal: 0, Name: "Outer", Encode: manifest.Optional(manifest.Optional(manifest.Primitive("u8"))), Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 1, Name: "Values", Encode: manifest.Array(manifest.Primitive("var_u32"), manifest.FixedArray(2, manifest.Primitive("u16le"))), Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 2, Name: "Choice", Encode: manifest.Union(manifest.Primitive("var_u32"), manifest.Variant{Value: 0, Name: "None", Encode: manifest.Void()}, manifest.Variant{Value: 7, Name: "Payload", Encode: manifest.Bytes(manifest.Primitive("var_u32"))}), Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 3, Name: "Mode", Encode: manifest.Enum("u8", manifest.EnumValue{Name: "Ready", Value: 4}, manifest.EnumValue{Name: "Later", Value: 9}), Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 4, Name: "Reserved", Encode: manifest.Reserved(manifest.Primitive("u16le")), Reserved: true, Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 5, Name: "Text", Encode: manifest.String(manifest.Primitive("var_u32")), Symmetry: manifest.Symmetric},
		manifest.Field{Ordinal: 6, Name: "NetworkID", Semantic: "NetworkRuntimeID", Encode: manifest.Primitive("u32le"), Symmetry: manifest.Symmetric},
	)
}
