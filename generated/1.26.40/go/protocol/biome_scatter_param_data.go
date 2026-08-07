// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BiomeScatterParamData struct {
	Coordinates       []BiomeCoordinateData
	EvalOrder         CoordinateEvaluationOrder
	ChancePercentType int32
	ChancePercent     uint16
	ChanceNumerator   int32
	ChanceDenominator int32
	IterationsType    int32
	Iterations        uint16
}

// Marshal reads or writes BiomeScatterParamData using its canonical wire layout.
func (x *BiomeScatterParamData) Marshal(io IO) {
	FuncSlice(io, &x.Coordinates, io.Varuint32, func(value *BiomeCoordinateData) {
		value.Marshal(io)
	})
	IntegerFunc(&x.EvalOrder, io.Varint32)
	io.Varint32(&x.ChancePercentType)
	io.Uint16(&x.ChancePercent)
	io.Int32(&x.ChanceNumerator)
	io.Int32(&x.ChanceDenominator)
	io.Varint32(&x.IterationsType)
	io.Uint16(&x.Iterations)
}
