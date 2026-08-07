// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type ServerConfigurationGatheringsConfigurationJoinInfo struct {
	ExperienceID   uuid.UUID
	ExperienceName string
	WorldID        Optional[uuid.UUID]
	WorldName      Optional[string]
	CreatorID      string
	TargetID       Optional[uuid.UUID]
	ScenarioID     Optional[string]
	ServerID       Optional[string]
}

// Marshal reads or writes ServerConfigurationGatheringsConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationGatheringsConfigurationJoinInfo) Marshal(io IO) {
	io.UUID(&x.ExperienceID)
	io.String(&x.ExperienceName)
	OptionalFunc(io, &x.WorldID, io.UUID)
	OptionalFunc(io, &x.WorldName, io.String)
	io.String(&x.CreatorID)
	OptionalFunc(io, &x.TargetID, io.UUID)
	OptionalFunc(io, &x.ScenarioID, io.String)
	OptionalFunc(io, &x.ServerID, io.String)
}
