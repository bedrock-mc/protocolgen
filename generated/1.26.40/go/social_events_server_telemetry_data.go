// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SocialEventsServerTelemetryData struct {
	ServerId   string
	ScenarioId string
	WorldId    string
	OwnerId    string
}

// Marshal reads or writes SocialEventsServerTelemetryData using its canonical wire layout.
func (x *SocialEventsServerTelemetryData) Marshal(io IO) {
	io.String(&x.ServerId)
	io.String(&x.ScenarioId)
	io.String(&x.WorldId)
	io.String(&x.OwnerId)
}
