// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AnimateEntity struct {
	MAnimation             string
	MNextState             string
	MStopExpression        string
	MStopExpressionVersion int32
	MController            string
	MBlendOutTime          float32
	MRuntimeIds            []uint64
}

// Marshal reads or writes AnimateEntity using its canonical wire layout.
func (x *AnimateEntity) Marshal(io protocol.IO) {
	io.String(&x.MAnimation)
	io.String(&x.MNextState)
	io.String(&x.MStopExpression)
	io.Int32(&x.MStopExpressionVersion)
	io.String(&x.MController)
	io.Float32(&x.MBlendOutTime)
	protocol.FuncSlice(io, &x.MRuntimeIds, io.Varuint32, io.ActorRuntimeID)
}
