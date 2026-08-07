// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SocialEventsServerTelemetryData struct {
	ServerID   string
	ScenarioID string
	WorldID    string
	OwnerID    string
}

// Marshal reads or writes SocialEventsServerTelemetryData using its canonical wire layout.
func (x *SocialEventsServerTelemetryData) Marshal(io IO) {
	io.String(&x.ServerID)
	io.String(&x.ScenarioID)
	io.String(&x.WorldID)
	io.String(&x.OwnerID)
}
