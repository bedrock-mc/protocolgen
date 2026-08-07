// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AnimateEntity struct {
	MAnimation             string
	MNextState             string
	MStopExpression        string
	MStopExpressionVersion int32
	MController            string
	MBlendOutTime          float32
	MRuntimeIds            []ActorRuntimeID
}

// Marshal reads or writes AnimateEntity using its canonical wire layout.
func (x *AnimateEntity) Marshal(io IO) {
	io.String(&x.MAnimation)
	io.String(&x.MNextState)
	io.String(&x.MStopExpression)
	io.Int32(&x.MStopExpressionVersion)
	io.String(&x.MController)
	io.Float32(&x.MBlendOutTime)
	if !io.Reading() && uint64(len(x.MRuntimeIds)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MRuntimeIds), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.MRuntimeIds))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.MRuntimeIds = make([]ActorRuntimeID, int(count1))
	}
	for index2 := range x.MRuntimeIds {
		x.MRuntimeIds[index2].Marshal(io)
	}
}
