// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CodeBuilderSource struct {
	Operation  CodeBuilderStorageQueryOptionsOperation
	Category   CodeBuilderStorageQueryOptionsCategory
	CodeStatus CodeBuilderExecutionStateCodeStatus
}

func (p *CodeBuilderSource) Encode(w Encoder) error {
	if err := w.Write("CodeBuilderSourcePacket.Operation", Shape{Kind: "enum", Semantic: "CodeBuilderStorageQueryOptions::Operation", TypeID: "enums/CodeBuilderStorageQueryOptions::Operation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Get", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Set", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Reset", Shape: Shape{Kind: "void"}}}}, p.Operation); err != nil {
		return err
	}
	if err := w.Write("CodeBuilderSourcePacket.Category", Shape{Kind: "enum", Semantic: "CodeBuilderStorageQueryOptions::Category", TypeID: "enums/CodeBuilderStorageQueryOptions::Category", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "CodeStatus", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Instantiation", Shape: Shape{Kind: "void"}}}}, p.Category); err != nil {
		return err
	}
	if err := w.Write("CodeBuilderSourcePacket.CodeStatus", Shape{Kind: "enum", Semantic: "CodeBuilderExecutionState::CodeStatus", TypeID: "enums/CodeBuilderExecutionState::CodeStatus", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "NotStarted", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "InProgress", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Paused", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Error", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Succeeded", Shape: Shape{Kind: "void"}}}}, p.CodeStatus); err != nil {
		return err
	}
	return nil
}

func DecodeCodeBuilderSource(r Decoder) (CodeBuilderSource, error) {
	var p CodeBuilderSource
	{
		raw, err := r.Read("CodeBuilderSourcePacket.Operation", Shape{Kind: "enum", Semantic: "CodeBuilderStorageQueryOptions::Operation", TypeID: "enums/CodeBuilderStorageQueryOptions::Operation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Get", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Set", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Reset", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CodeBuilderStorageQueryOptionsOperation)
		if !ok {
			return p, fmt.Errorf("field CodeBuilderSourcePacket.Operation has unexpected decoded type %T", raw)
		}
		p.Operation = value
	}
	{
		raw, err := r.Read("CodeBuilderSourcePacket.Category", Shape{Kind: "enum", Semantic: "CodeBuilderStorageQueryOptions::Category", TypeID: "enums/CodeBuilderStorageQueryOptions::Category", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "CodeStatus", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Instantiation", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CodeBuilderStorageQueryOptionsCategory)
		if !ok {
			return p, fmt.Errorf("field CodeBuilderSourcePacket.Category has unexpected decoded type %T", raw)
		}
		p.Category = value
	}
	{
		raw, err := r.Read("CodeBuilderSourcePacket.CodeStatus", Shape{Kind: "enum", Semantic: "CodeBuilderExecutionState::CodeStatus", TypeID: "enums/CodeBuilderExecutionState::CodeStatus", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "NotStarted", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "InProgress", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Paused", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Error", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Succeeded", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CodeBuilderExecutionStateCodeStatus)
		if !ok {
			return p, fmt.Errorf("field CodeBuilderSourcePacket.CodeStatus has unexpected decoded type %T", raw)
		}
		p.CodeStatus = value
	}
	return p, nil
}
