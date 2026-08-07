// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type JigsawStructureData struct {
	JigsawStructureDataTag []byte
}

func (p *JigsawStructureData) Encode(w Encoder) error {
	if err := w.Write("JigsawStructureDataPacket.Jigsaw Structure Data Tag", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.JigsawStructureDataTag); err != nil {
		return err
	}
	return nil
}

func DecodeJigsawStructureData(r Decoder) (JigsawStructureData, error) {
	var p JigsawStructureData
	{
		raw, err := r.Read("JigsawStructureDataPacket.Jigsaw Structure Data Tag", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field JigsawStructureDataPacket.Jigsaw Structure Data Tag has unexpected decoded type %T", raw)
		}
		p.JigsawStructureDataTag = value
	}
	return p, nil
}
