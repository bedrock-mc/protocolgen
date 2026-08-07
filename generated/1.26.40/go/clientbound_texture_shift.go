// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundTextureShift struct {
	ActionID             ClientboundTextureShiftAction
	CollectionName       string
	FromStep             string
	ToStep               string
	AllSteps             []string
	CurrentLengthInTicks uint64
	TotalLengthInTicks   uint64
	Enabled              bool
}
