// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PurchaseReceipt struct {
	PurchaseReceipts []string
}

func (p *PurchaseReceipt) Encode(w Encoder) error {
	if err := w.Write("PurchaseReceiptPacket.PurchaseReceipts", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.PurchaseReceipts); err != nil {
		return err
	}
	return nil
}

func DecodePurchaseReceipt(r Decoder) (PurchaseReceipt, error) {
	var p PurchaseReceipt
	{
		raw, err := r.Read("PurchaseReceiptPacket.PurchaseReceipts", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]string)
		if !ok {
			return p, fmt.Errorf("field PurchaseReceiptPacket.PurchaseReceipts has unexpected decoded type %T", raw)
		}
		p.PurchaseReceipts = value
	}
	return p, nil
}
