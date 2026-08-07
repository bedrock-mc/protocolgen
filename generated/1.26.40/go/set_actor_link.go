// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetActorLink struct {
	Link ActorLink
}

func (p *SetActorLink) Encode(w Encoder) error {
	if err := w.Write("SetActorLinkPacket.Link", Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Link); err != nil {
		return err
	}
	return nil
}

func DecodeSetActorLink(r Decoder) (SetActorLink, error) {
	var p SetActorLink
	{
		raw, err := r.Read("SetActorLinkPacket.Link", Shape{Kind: "struct", Semantic: "ActorLink", TypeID: "ActorLink", Fields: []ShapeField{{Ordinal: 0, Name: "Target A", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 1, Name: "Target B", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 2, Name: "Type", Shape: Shape{Kind: "enum", Semantic: "ActorLinkType", TypeID: "enums/ActorLinkType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Riding", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Passenger", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Immediate", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Passenger Initiated", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Vehicle Angular Velocity", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorLink)
		if !ok {
			return p, fmt.Errorf("field SetActorLinkPacket.Link has unexpected decoded type %T", raw)
		}
		p.Link = value
	}
	return p, nil
}
