// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CommandPermissionLevel uint8

const (
	CommandPermissionLevelAny           CommandPermissionLevel = 0
	CommandPermissionLevelGameDirectors CommandPermissionLevel = 1
	CommandPermissionLevelAdmin         CommandPermissionLevel = 2
	CommandPermissionLevelHost          CommandPermissionLevel = 3
	CommandPermissionLevelOwner         CommandPermissionLevel = 4
	CommandPermissionLevelInternal      CommandPermissionLevel = 5
)
