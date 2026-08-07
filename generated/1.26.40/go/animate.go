// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Animate struct {
	Action               AnimateAction
	TargetActorRuntimeID ActorRuntimeID
	Data                 float32
	SwingSource          Optional[string]
}
