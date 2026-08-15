// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// AttributeData represents a polymorphic attribute value.
type AttributeData struct {
	MinValue        float32
	MaxValue        float32
	CurrentValue    float32
	DefaultMinValue float32
	DefaultMaxValue float32
	DefaultValue    float32
	Name            string
	Modifiers       []AttributeModifier
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
	AttributeLayers []EASAttributeLayerData
}

func (*AttributeLayerData) isAttributeLayerSyncData() {}

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

func (*AttributeLayerSettings) isAttributeLayerSyncData() {}

// Marshal reads or writes AttributeLayerSettings using its canonical wire layout.
func (x *AttributeLayerSettings) Marshal(io IO) {
	io.StringLimits(&x.AttributeLayerName, 0, 128)
	x.AttributeLayerDimension.Marshal(io)
	x.AttributesLayerSettings.Marshal(io)
}

type AttributeLayerSyncData interface {
	isAttributeLayerSyncData()
}

// MarshalAttributeLayerSyncData reads or writes the AttributeLayerSyncData union using its canonical wire layout.
func MarshalAttributeLayerSyncData(io IO, x *AttributeLayerSyncData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(AttributeLayerData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(AttributeLayerSettings)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(EnvironmentAttributeData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(RemoveEnvironmentAttributes)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *AttributeLayerData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AttributeLayerSettings:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *EnvironmentAttributeData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *RemoveEnvironmentAttributes:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

// EnvironmentAttributeData represents an environment attribute with optional transition data.
type EnvironmentAttributeData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EASEnvironmentAttributeData
}

func (*EnvironmentAttributeData) isAttributeLayerSyncData() {}

// Marshal reads or writes EnvironmentAttributeData using its canonical wire layout.
func (x *EnvironmentAttributeData) Marshal(io IO) {
	io.StringLimits(&x.AttributeLayerName, 0, 128)
	x.AttributeLayerDimension.Marshal(io)
	SliceLimits(io, &x.Attributes, 0, 1024)
}
