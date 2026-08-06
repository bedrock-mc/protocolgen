// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

// Shape is the frozen wire vocabulary rendered from the canonical manifest.
// A profile may implement Encoder/Decoder, but it cannot alter Shape.
type Shape struct {
	Kind string
	Semantic string
	TypeID string
	PrimitiveCode string
	Encoding string
	Representation string
	Prefix *Shape
	Element *Shape
	Length uint64
	Value *Shape
	Elements []Shape
	Key *Shape
	Fields []ShapeField
	Variants []ShapeVariant
	Control *Shape
	CompareTo string
	Cases []ShapeCase
	Default *Shape
	Target string
}

type ShapeField struct { Ordinal int; Name string; Shape Shape }
type ShapeVariant struct { Value int64; Name string; Shape Shape }
type ShapeCase struct { Value string; Shapes []Shape }

type Encoder interface { Write(path string, shape Shape, value any) error }
type Decoder interface { Read(path string, shape Shape) (any, error) }
