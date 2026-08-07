// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CreatePhoto struct {
	RawID         uint64
	PhotoName     string
	PhotoItemName string
}

func (p *CreatePhoto) Encode(w Encoder) error {
	if err := w.Write("CreatePhotoPacket.Raw ID", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.RawID); err != nil {
		return err
	}
	if err := w.Write("CreatePhotoPacket.Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PhotoName); err != nil {
		return err
	}
	if err := w.Write("CreatePhotoPacket.Photo Item Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.PhotoItemName); err != nil {
		return err
	}
	return nil
}

func DecodeCreatePhoto(r Decoder) (CreatePhoto, error) {
	var p CreatePhoto
	{
		raw, err := r.Read("CreatePhotoPacket.Raw ID", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field CreatePhotoPacket.Raw ID has unexpected decoded type %T", raw)
		}
		p.RawID = value
	}
	{
		raw, err := r.Read("CreatePhotoPacket.Photo Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CreatePhotoPacket.Photo Name has unexpected decoded type %T", raw)
		}
		p.PhotoName = value
	}
	{
		raw, err := r.Read("CreatePhotoPacket.Photo Item Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field CreatePhotoPacket.Photo Item Name has unexpected decoded type %T", raw)
		}
		p.PhotoItemName = value
	}
	return p, nil
}
