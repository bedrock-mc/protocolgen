// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EducationEditionOffer uint32

const (
	EducationEditionOfferNone            EducationEditionOffer = 0
	EducationEditionOfferRestOfWorld     EducationEditionOffer = 1
	EducationEditionOfferChinaDeprecated EducationEditionOffer = 2
)

// Marshal reads or writes EducationEditionOffer through its uint32 wire encoding.
func (x *EducationEditionOffer) Marshal(io IO) { io.Varuint32((*uint32)(x)) }

type EducationLevelSettings struct {
	CodeBuilderDefaultURI        string
	CodeBuilderTitle             string
	CanResizeCodeBuilder         bool
	DisableLegacyTitleBar        bool
	PostProcessFilter            string
	ScreenshotBorderResourcePath string
	AgentCapabilities            Optional[bool]
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
	OptionalFunc(io, &x.AgentCapabilities, io.Bool)
	x.LocalSettings.Marshal(io)
	io.Bool(&x.DeprecatedAlwaysFalse)
	OptionalMarshaler(io, &x.ExternalLinkSettings)
}

type EducationLocalLevelSettings struct {
	CodeBuilderOverrideURI Optional[string]
}

// Marshal reads or writes EducationLocalLevelSettings using its canonical wire layout.
func (x *EducationLocalLevelSettings) Marshal(io IO) {
	OptionalFunc(io, &x.CodeBuilderOverrideURI, io.String)
}
