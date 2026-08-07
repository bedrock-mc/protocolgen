// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration Optional[ServerConfigurationGatheringsConfigurationJoinInfo]
}
