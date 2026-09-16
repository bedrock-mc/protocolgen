// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// UpdateTrade is sent by the server to update the trades offered by a villager to a player. It is
// sent at the moment that a player interacts with a villager.
type UpdateTrade struct {
	ContainerID uint8
	Type        uint8
	// Size is the amount of trading options that the villager has.
	Size       int32
	TraderTier int32
	// EntityUniqueID is the unique ID of the entity (usually a player) for which the trades are
	// updated. The updated trades may apply only to this entity.
	EntityUniqueID    int64
	LastTradingPlayer int64
	// DisplayName is the name displayed at the top of the trading UI. It is usually used to represent
	// the profession of the villager in the UI.
	DisplayName       string
	UseNewTradeScreen bool
	UsingEconomyTrade bool
	Data              []byte
}

// Marshal reads or writes UpdateTrade using its canonical wire layout.
func (x *UpdateTrade) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Uint8(&x.Type)
	io.Varint32(&x.Size)
	io.Varint32(&x.TraderTier)
	io.ActorUniqueID(&x.EntityUniqueID)
	io.ActorUniqueID(&x.LastTradingPlayer)
	io.String(&x.DisplayName)
	io.Bool(&x.UseNewTradeScreen)
	io.Bool(&x.UsingEconomyTrade)
	io.NBT(&x.Data, protocol.NBTNetwork)
}

// ID returns the protocol ID for UpdateTrade.
func (*UpdateTrade) ID() uint32 { return IDUpdateTrade }
