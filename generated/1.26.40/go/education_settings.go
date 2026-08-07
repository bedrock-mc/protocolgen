// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EducationSettings struct {
	EducationLevelSettings EducationLevelSettings
}

// Marshal reads or writes EducationSettings using its canonical wire layout.
func (x *EducationSettings) Marshal(io IO) {
	x.EducationLevelSettings.Marshal(io)
}
