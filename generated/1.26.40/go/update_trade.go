// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateTrade struct {
	ContainerId       uint8
	Type              uint8
	Size              int32
	TraderTier        int32
	EntityUniqueId    int64
	LastTradingPlayer int64
	DisplayName       string
	UseNewTradeScreen bool
	UsingEconomyTrade bool
	Data              []byte
}

// Marshal reads or writes UpdateTrade using its canonical wire layout.
func (x *UpdateTrade) Marshal(io IO) {
	io.Uint8(&x.ContainerId)
	io.Uint8(&x.Type)
	io.Varint32(&x.Size)
	io.Varint32(&x.TraderTier)
	io.ActorUniqueID(&x.EntityUniqueId)
	io.ActorUniqueID(&x.LastTradingPlayer)
	io.String(&x.DisplayName)
	io.Bool(&x.UseNewTradeScreen)
	io.Bool(&x.UsingEconomyTrade)
	io.NBT(&x.Data)
}
