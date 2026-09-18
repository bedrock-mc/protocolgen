// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// AddPlayer is sent by the server to the client to make a player entity show up client-side. It is one of the
// few entities that cannot be sent using the AddActor packet.
type AddPlayer struct {
	// UUID is the UUID of the player. It is the same UUID that the client sent in the Login packet at the start
	// of the session. A player with this UUID must exist in the player list (built up using the PlayerList
	// packet), for it to show up in-game.
	UUID            uuid.UUID
	PlayerName      string
	TargetRuntimeID uint64
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to chat
	// with each other.
	PlatformChatID string
	// Position is the position to spawn the player on. If the player is on a distance that the viewer cannot see
	// it, the player will still show up if the viewer moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the player spawns with. This velocity will initiate client side movement
	// of the player.
	Velocity          mgl32.Vec3
	Rotation          mgl32.Vec2
	YHeadRotation     float32
	CarriedItem       protocol.NetworkItemStackDescriptorSerializedData
	PlayerGameType    protocol.GameType
	EntityData        protocol.SynchedActorDataCopyableDataList
	SynchedProperties protocol.PropertySyncData
	AbilitiesData     protocol.SerializedAbilitiesData
	ActorLinks        []protocol.EntityLink
	// DeviceID is the device ID set in one of the files found in the storage of the device of the player. It may
	// be changed freely, so it should not be relied on for anything.
	DeviceID string
	// BuildPlatform is the build platform/device OS of the player that is about to be added, as it sent in the
	// Login packet when joining.
	BuildPlatform protocol.BuildPlatform
}

// ID returns the protocol ID for AddPlayer.
func (*AddPlayer) ID() uint32 { return IDAddPlayer }

// Marshal reads or writes AddPlayer using its canonical wire layout.
func (pk *AddPlayer) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	io.String(&pk.PlayerName)
	io.ActorRuntimeID(&pk.TargetRuntimeID)
	io.String(&pk.PlatformChatID)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Vec2(&pk.Rotation)
	io.Float32(&pk.YHeadRotation)
	pk.CarriedItem.Marshal(io)
	pk.PlayerGameType.Marshal(io)
	pk.EntityData.Marshal(io)
	pk.SynchedProperties.Marshal(io)
	pk.AbilitiesData.Marshal(io)
	protocol.Slice(io, &pk.ActorLinks)
	io.String(&pk.DeviceID)
	pk.BuildPlatform.Marshal(io)
}
