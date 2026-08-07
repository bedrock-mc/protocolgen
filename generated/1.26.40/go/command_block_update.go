// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CommandBlockUpdate struct {
	Target             CommandBlockUpdateTarget
	Command            string
	LastOutput         string
	Name               string
	FilteredName       string
	TrackOutput        bool
	TickDelay          int32
	ExecuteOnFirstTick bool
}
