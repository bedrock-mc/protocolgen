// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PhotoTransfer struct {
	PhotoName    string
	PhotoData    string
	BookID       string
	Type         PhotoType
	SourceType   PhotoType
	OwnerID      int64
	NewPhotoName string
}

func (p *PhotoTransfer) Encode(w Encoder) error {
	if err := w.Write("PhotoTransferPacket.Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PhotoName); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.Photo Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PhotoData); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.Book ID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.BookID); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.Type", Shape{Kind: "enum", Semantic: "PhotoType", TypeID: "enums/PhotoType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Portfolio", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PhotoItem", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Book", Shape: Shape{Kind: "void"}}}}, p.Type); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.Source Type", Shape{Kind: "enum", Semantic: "PhotoType", TypeID: "enums/PhotoType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Portfolio", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PhotoItem", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Book", Shape: Shape{Kind: "void"}}}}, p.SourceType); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.Owner ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"}, p.OwnerID); err != nil {
		return err
	}
	if err := w.Write("PhotoTransferPacket.New Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.NewPhotoName); err != nil {
		return err
	}
	return nil
}

func DecodePhotoTransfer(r Decoder) (PhotoTransfer, error) {
	var p PhotoTransfer
	{
		raw, err := r.Read("PhotoTransferPacket.Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Photo Name has unexpected decoded type %T", raw)
		}
		p.PhotoName = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.Photo Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Photo Data has unexpected decoded type %T", raw)
		}
		p.PhotoData = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.Book ID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Book ID has unexpected decoded type %T", raw)
		}
		p.BookID = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.Type", Shape{Kind: "enum", Semantic: "PhotoType", TypeID: "enums/PhotoType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Portfolio", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PhotoItem", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Book", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PhotoType)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Type has unexpected decoded type %T", raw)
		}
		p.Type = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.Source Type", Shape{Kind: "enum", Semantic: "PhotoType", TypeID: "enums/PhotoType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Portfolio", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "PhotoItem", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Book", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PhotoType)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Source Type has unexpected decoded type %T", raw)
		}
		p.SourceType = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.Owner ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int64)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.Owner ID has unexpected decoded type %T", raw)
		}
		p.OwnerID = value
	}
	{
		raw, err := r.Read("PhotoTransferPacket.New Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PhotoTransferPacket.New Photo Name has unexpected decoded type %T", raw)
		}
		p.NewPhotoName = value
	}
	return p, nil
}
