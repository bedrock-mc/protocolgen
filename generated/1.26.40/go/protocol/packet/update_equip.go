// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// UpdateEquip is sent by the server to the client upon opening a horse inventory. It is used to set
// the content of the inventory and specify additional properties, such as the items that are
// allowed to be put in slots of the inventory.
type UpdateEquip struct {
	ContainerID uint8
	Type        uint8
	// Size is the size of the horse inventory that should be opened. A bigger size does, in fact,
	// change the amount of slots displayed.
	Size int32
	// EntityUniqueID is the unique ID of the entity whose equipment was 'updated' to the player. It is
	// typically the horse entity that had its inventory opened.
	EntityUniqueID int64
	Data           []byte
}

// Marshal reads or writes UpdateEquip using its canonical wire layout.
func (x *UpdateEquip) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Uint8(&x.Type)
	io.Varint32(&x.Size)
	io.ActorUniqueID(&x.EntityUniqueID)
	io.NBT(&x.Data, protocol.NBTNetwork)
}

// ID returns the protocol ID for UpdateEquip.
func (*UpdateEquip) ID() uint32 { return IDUpdateEquip }
