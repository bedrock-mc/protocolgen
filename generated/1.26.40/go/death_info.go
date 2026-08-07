// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DeathInfo struct {
	DeathCauseAttackName  string
	DeathCauseMessageList []string
}

// Marshal reads or writes DeathInfo using its canonical wire layout.
func (x *DeathInfo) Marshal(io IO) {
	io.String(&x.DeathCauseAttackName)
	FuncSlice(io, &x.DeathCauseMessageList, io.Varuint32, io.String)
}
