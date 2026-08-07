// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ActorDataBoundingBoxComponent struct {
	ActorDataBoundingBox [3]float32
}

// Marshal reads or writes ActorDataBoundingBoxComponent using its canonical wire layout.
func (x *ActorDataBoundingBoxComponent) Marshal(io IO) {
	for index1 := range x.ActorDataBoundingBox {
		io.Float32(&x.ActorDataBoundingBox[index1])
	}
}
