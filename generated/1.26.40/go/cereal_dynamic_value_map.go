// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CerealDynamicValueMap struct {
	Value []OrderedEntry[string, CerealDynamicValue]
}

func (CerealDynamicValueMap) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueMap using its canonical wire layout.
func (x *CerealDynamicValueMap) Marshal(io IO) {
	OrderedMap(io, &x.Value, io.Varuint32, io.String, func(value *CerealDynamicValue) {
		item := *value
		marshalCerealDynamicValue(io, &item)
		*value = item
	})
}
