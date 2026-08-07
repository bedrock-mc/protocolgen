// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type ServerConfigurationGatheringsConfigurationJoinInfo struct {
	ExperienceId   uuid.UUID
	ExperienceName string
	WorldId        Optional[uuid.UUID]
	WorldName      Optional[string]
	CreatorId      string
	TargetId       Optional[uuid.UUID]
	ScenarioId     Optional[string]
	ServerId       Optional[string]
}

// Marshal reads or writes ServerConfigurationGatheringsConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationGatheringsConfigurationJoinInfo) Marshal(io IO) {
	io.UUID(&x.ExperienceId)
	io.String(&x.ExperienceName)
	OptionalFunc(io, &x.WorldId, io.UUID)
	OptionalFunc(io, &x.WorldName, io.String)
	io.String(&x.CreatorId)
	OptionalFunc(io, &x.TargetId, io.UUID)
	OptionalFunc(io, &x.ScenarioId, io.String)
	OptionalFunc(io, &x.ServerId, io.String)
}
