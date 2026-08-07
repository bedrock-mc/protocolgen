// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DeathInfo struct {
	DeathCauseAttackName  string
	DeathCauseMessageList []string
}

// Marshal reads or writes DeathInfo using its canonical wire layout.
func (x *DeathInfo) Marshal(io IO) {
	io.String(&x.DeathCauseAttackName)
	if !io.Reading() && uint64(len(x.DeathCauseMessageList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.DeathCauseMessageList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.DeathCauseMessageList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.DeathCauseMessageList = make([]string, int(count1))
	}
	for index2 := range x.DeathCauseMessageList {
		io.String(&x.DeathCauseMessageList[index2])
	}
}
