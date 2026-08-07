// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EducationLevelSettings struct {
	CodeBuilderDefaultURI        string
	CodeBuilderTitle             string
	CanResizeCodeBuilder         bool
	DisableLegacyTitleBar        bool
	PostProcessFilter            string
	ScreenshotBorderResourcePath string
	AgentCapabilities            Optional[AgentCapabilities]
	LocalSettings                EducationLocalLevelSettings
	DeprecatedAlwaysFalse        bool
	ExternalLinkSettings         Optional[ExternalLinkSettings]
}

// Marshal reads or writes EducationLevelSettings using its canonical wire layout.
func (x *EducationLevelSettings) Marshal(io IO) {
	io.String(&x.CodeBuilderDefaultURI)
	io.String(&x.CodeBuilderTitle)
	io.Bool(&x.CanResizeCodeBuilder)
	io.Bool(&x.DisableLegacyTitleBar)
	io.String(&x.PostProcessFilter)
	io.String(&x.ScreenshotBorderResourcePath)
	OptionalFunc(io, &x.AgentCapabilities, func(value *AgentCapabilities) {
		value.Marshal(io)
	})
	x.LocalSettings.Marshal(io)
	io.Bool(&x.DeprecatedAlwaysFalse)
	OptionalFunc(io, &x.ExternalLinkSettings, func(value *ExternalLinkSettings) {
		value.Marshal(io)
	})
}
