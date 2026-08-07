// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type AddBehaviorTree struct {
	BehaviorTreeStructureJSON string
}

func (p *AddBehaviorTree) Encode(w Encoder) error {
	if err := w.Write("AddBehaviorTreePacket.Behavior Tree Structure (JSON)", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.BehaviorTreeStructureJSON); err != nil {
		return err
	}
	return nil
}

func DecodeAddBehaviorTree(r Decoder) (AddBehaviorTree, error) {
	var p AddBehaviorTree
	{
		raw, err := r.Read("AddBehaviorTreePacket.Behavior Tree Structure (JSON)", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field AddBehaviorTreePacket.Behavior Tree Structure (JSON) has unexpected decoded type %T", raw)
		}
		p.BehaviorTreeStructureJSON = value
	}
	return p, nil
}
