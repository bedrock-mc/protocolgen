// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// AttributeData represents a polymorphic attribute value.
type AttributeData struct {
	MinValue        float32
	MaxValue        float32
	CurrentValue    float32
	DefaultMinValue float32
	DefaultMaxValue float32
	// FloatValue is the float value if Type is AttributeDataTypeFloat.
	DefaultValue float32
	Name         string
	Modifiers    []AttributeModifier
}

// Marshal reads or writes AttributeData using its canonical wire layout.
func (x *AttributeData) Marshal(io IO) {
	io.Float32(&x.MinValue)
	io.Float32(&x.MaxValue)
	io.Float32(&x.CurrentValue)
	io.Float32(&x.DefaultMinValue)
	io.Float32(&x.DefaultMaxValue)
	io.Float32(&x.DefaultValue)
	io.String(&x.Name)
	Slice(io, &x.Modifiers)
}

// AttributeLayerData represents a complete attribute layer.
type AttributeLayerData struct {
	// EnvironmentAttributes is the list of environment attributes in this layer.
	AttributeLayers []EASAttributeLayerData
}

func (*AttributeLayerData) tagAttributeLayerSyncData() uint32 { return 0 }

// Marshal reads or writes AttributeLayerData using its canonical wire layout.
func (x *AttributeLayerData) Marshal(io IO) {
	SliceLimits(io, &x.AttributeLayers, 0, 512)
}

// AttributeLayerSettings represents settings for an attribute layer.
type AttributeLayerSettings struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	AttributesLayerSettings EASAttributeLayerSettings
}

func (*AttributeLayerSettings) tagAttributeLayerSyncData() uint32 { return 1 }

// Marshal reads or writes AttributeLayerSettings using its canonical wire layout.
func (x *AttributeLayerSettings) Marshal(io IO) {
	io.StringLimits(&x.AttributeLayerName, 0, 128)
	x.AttributeLayerDimension.Marshal(io)
	x.AttributesLayerSettings.Marshal(io)
}

type AttributeLayerSyncData interface {
	Marshaler
	tagAttributeLayerSyncData() uint32
}

// MarshalAttributeLayerSyncData reads or writes the AttributeLayerSyncData union using its canonical wire layout.
func MarshalAttributeLayerSyncData(io IO, x *AttributeLayerSyncData) {
	Union(io, x, io.Varuint32, AttributeLayerSyncData.tagAttributeLayerSyncData, func(tag uint32) AttributeLayerSyncData {
		switch tag {
		case 0:
			return new(AttributeLayerData)
		case 1:
			return new(AttributeLayerSettings)
		case 2:
			return new(EnvironmentAttributeData)
		case 3:
			return new(RemoveEnvironmentAttributes)
		}
		return nil
	})
}

// EnvironmentAttributeData represents an environment attribute with optional transition data.
type EnvironmentAttributeData struct {
	// AttributeName is the name of the attribute.
	AttributeLayerName string
	// Attribute is the current attribute value.
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (*EnvironmentAttributeData) tagAttributeLayerSyncData() uint32 { return 2 }

// Marshal reads or writes EnvironmentAttributeData using its canonical wire layout.
func (x *EnvironmentAttributeData) Marshal(io IO) {
	io.StringLimits(&x.AttributeLayerName, 0, 128)
	x.AttributeLayerDimension.Marshal(io)
	SliceLimits(io, &x.Attributes, 0, 1024)
}
