// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EASColorAttributeData struct {
	Value     [4]int32
	Operation string
}

func (EASColorAttributeData) isEAS() {}

// Marshal reads or writes EASColorAttributeData using its canonical wire layout.
func (x *EASColorAttributeData) Marshal(io IO) {
	for index1 := range x.Value {
		io.Int32(&x.Value[index1])
	}
	io.String(&x.Operation)
}
