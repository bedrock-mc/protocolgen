// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateTrade struct {
	ContainerId       uint8
	Type              uint8
	Size              int32
	TraderTier        int32
	EntityUniqueId    ActorUniqueID
	LastTradingPlayer ActorUniqueID
	DisplayName       string
	UseNewTradeScreen bool
	UsingEconomyTrade bool
	Data              []byte
}
