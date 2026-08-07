// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LessonProgress struct {
	LessonAction int32
	Score        int32
	ActivityId   string
}

func (p *LessonProgress) Encode(w Encoder) error {
	if err := w.Write("LessonProgressPacket.Lesson Action", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.LessonAction); err != nil {
		return err
	}
	if err := w.Write("LessonProgressPacket.Score", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Score); err != nil {
		return err
	}
	if err := w.Write("LessonProgressPacket.Activity Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ActivityId); err != nil {
		return err
	}
	return nil
}

func DecodeLessonProgress(r Decoder) (LessonProgress, error) {
	var p LessonProgress
	{
		raw, err := r.Read("LessonProgressPacket.Lesson Action", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LessonProgressPacket.Lesson Action has unexpected decoded type %T", raw)
		}
		p.LessonAction = value
	}
	{
		raw, err := r.Read("LessonProgressPacket.Score", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LessonProgressPacket.Score has unexpected decoded type %T", raw)
		}
		p.Score = value
	}
	{
		raw, err := r.Read("LessonProgressPacket.Activity Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field LessonProgressPacket.Activity Id has unexpected decoded type %T", raw)
		}
		p.ActivityId = value
	}
	return p, nil
}
