// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EducationLocalLevelSettings struct {
	CodeBuilderOverrideUri Optional[string]
}

// Marshal reads or writes EducationLocalLevelSettings using its canonical wire layout.
func (x *EducationLocalLevelSettings) Marshal(io IO) {
	OptionalFunc(io, &x.CodeBuilderOverrideUri, io.String)
}
