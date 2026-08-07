// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ModalFormResponse struct {
	FormID           uint32
	JSONResponse     Optional[string]
	FormCancelReason Optional[ModalFormCancelReason]
}

// Marshal reads or writes ModalFormResponse using its canonical wire layout.
func (x *ModalFormResponse) Marshal(io IO) {
	io.Varuint32(&x.FormID)
	OptionalFunc(io, &x.JSONResponse, io.String)
	OptionalFunc(io, &x.FormCancelReason, func(value *ModalFormCancelReason) {
		item := *value
		IntegerFunc(&item, io.Uint8)
		*value = item
	})
}
