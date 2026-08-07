// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetDisplayObjective struct {
	DisplaySlotName      string
	ObjectiveName        string
	ObjectiveDisplayName string
	CriteriaName         string
	SortOrder            int32
}

func (p *SetDisplayObjective) Encode(w Encoder) error {
	if err := w.Write("SetDisplayObjectivePacket.DisplaySlotName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.DisplaySlotName); err != nil {
		return err
	}
	if err := w.Write("SetDisplayObjectivePacket.ObjectiveName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ObjectiveName); err != nil {
		return err
	}
	if err := w.Write("SetDisplayObjectivePacket.ObjectiveDisplayName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ObjectiveDisplayName); err != nil {
		return err
	}
	if err := w.Write("SetDisplayObjectivePacket.CriteriaName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.CriteriaName); err != nil {
		return err
	}
	if err := w.Write("SetDisplayObjectivePacket.SortOrder", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.SortOrder); err != nil {
		return err
	}
	return nil
}

func DecodeSetDisplayObjective(r Decoder) (SetDisplayObjective, error) {
	var p SetDisplayObjective
	{
		raw, err := r.Read("SetDisplayObjectivePacket.DisplaySlotName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetDisplayObjectivePacket.DisplaySlotName has unexpected decoded type %T", raw)
		}
		p.DisplaySlotName = value
	}
	{
		raw, err := r.Read("SetDisplayObjectivePacket.ObjectiveName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetDisplayObjectivePacket.ObjectiveName has unexpected decoded type %T", raw)
		}
		p.ObjectiveName = value
	}
	{
		raw, err := r.Read("SetDisplayObjectivePacket.ObjectiveDisplayName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetDisplayObjectivePacket.ObjectiveDisplayName has unexpected decoded type %T", raw)
		}
		p.ObjectiveDisplayName = value
	}
	{
		raw, err := r.Read("SetDisplayObjectivePacket.CriteriaName", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field SetDisplayObjectivePacket.CriteriaName has unexpected decoded type %T", raw)
		}
		p.CriteriaName = value
	}
	{
		raw, err := r.Read("SetDisplayObjectivePacket.SortOrder", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field SetDisplayObjectivePacket.SortOrder has unexpected decoded type %T", raw)
		}
		p.SortOrder = value
	}
	return p, nil
}
