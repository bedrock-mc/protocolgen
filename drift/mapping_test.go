package main

import "testing"

func TestCompatibleEncoding(t *testing.T) {
	cases := []struct {
		name string
		doc  DocField
		op   WireOp
		want string
	}{
		{
			name: "varuint64 matches uint64+Compression",
			doc:  DocField{UnderlyingType: "uint64", Options: []string{"Compression"}, JSONType: "integer"},
			op:   WireOp{Kind: "io", Method: "Varuint64"},
			want: statusOK,
		},
		{
			name: "fixed uint64 mismatches uint64+Compression",
			doc:  DocField{UnderlyingType: "uint64", Options: []string{"Compression"}, JSONType: "integer"},
			op:   WireOp{Kind: "io", Method: "Uint64"},
			want: statusMismatch,
		},
		{
			name: "boolean underlying-type matches Bool",
			doc:  DocField{UnderlyingType: "boolean", JSONType: "boolean"},
			op:   WireOp{Kind: "io", Method: "Bool"},
			want: statusOK,
		},
		{
			name: "fixed int32 matches int32 without compression",
			doc:  DocField{UnderlyingType: "int32", JSONType: "integer"},
			op:   WireOp{Kind: "io", Method: "Int32"},
			want: statusOK,
		},
		{
			name: "varint32 mismatches int32 without compression",
			doc:  DocField{UnderlyingType: "int32", JSONType: "integer"},
			op:   WireOp{Kind: "io", Method: "Varint32"},
			want: statusMismatch,
		},
		{
			name: "float matches Float32",
			doc:  DocField{UnderlyingType: "float", JSONType: "number"},
			op:   WireOp{Kind: "io", Method: "Float32"},
			want: statusOK,
		},
		{
			name: "string matches String",
			doc:  DocField{UnderlyingType: "string", JSONType: "string"},
			op:   WireOp{Kind: "io", Method: "String"},
			want: statusOK,
		},
		{
			name: "nbt blob (doc string) matches NBT op",
			doc:  DocField{UnderlyingType: "string", JSONType: "string"},
			op:   WireOp{Kind: "io", Method: "NBT"},
			want: statusOK,
		},
		{
			name: "enum without Enum-as-Value matches String",
			doc:  DocField{UnderlyingType: "uint8", JSONType: "string", Enum: true},
			op:   WireOp{Kind: "io", Method: "String"},
			want: statusOK,
		},
		{
			name: "enum with Enum-as-Value matches underlying uint8",
			doc:  DocField{UnderlyingType: "uint8", JSONType: "string", Enum: true, Options: []string{"Enum-as-Value"}},
			op:   WireOp{Kind: "io", Method: "Uint8"},
			want: statusOK,
		},
		{
			name: "composite BlockPos matches BlockPos op",
			doc:  DocField{Composite: true, RefTitle: "BlockPos"},
			op:   WireOp{Kind: "io", Method: "BlockPos"},
			want: statusOK,
		},
		{
			name: "composite BlockPos mismatches Vec3 op",
			doc:  DocField{Composite: true, RefTitle: "BlockPos"},
			op:   WireOp{Kind: "io", Method: "Vec3"},
			want: statusMismatch,
		},
		{
			name: "protocol helper needs manual review",
			doc:  DocField{UnderlyingType: "uint8", JSONType: "string", Enum: true},
			op:   WireOp{Kind: "protocol", Method: "OptionalFunc", Inner: "String"},
			want: statusReview,
		},
		{
			name: "custom handcoded op needs manual review",
			doc:  DocField{UnderlyingType: "uint8"},
			op:   WireOp{Kind: "custom", Method: "swingSourceFromString"},
			want: statusReview,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := compatibleEncoding(c.doc, c.op)
			if got != c.want {
				t.Errorf("compatibleEncoding(%+v, %+v) = %q, want %q", c.doc, c.op, got, c.want)
			}
		})
	}
}
