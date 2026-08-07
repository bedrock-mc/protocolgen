// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DimensionData struct {
	Definitions []OrderedEntry[string, DimensionDefinition]
}

// Marshal reads or writes DimensionData using its canonical wire layout.
func (x *DimensionData) Marshal(io IO) {
	OrderedMap(io, &x.Definitions, io.Varuint32, io.String, func(value *DimensionDefinition) {
		value.Marshal(io)
	})
}
