package vanilladata

// Resource-pack metadata is split out of compatibility.go because it is a
// direct projection of the two login packets, rather than a compatibility
// representation from another data project. The generated codec adapters fill
// these version-neutral values from the exact target packet package.

import (
	"encoding/hex"
	"fmt"
)

type generatedResourcePack struct {
	UUID                string
	Version             string
	Size                uint64
	ContentKey          string
	SubpackName         string
	ContentIdentity     string
	HasScripts          bool
	IsAddonPack         bool
	IsRayTracingCapable bool
	CDNURL              string
}

type generatedResourcePackInstance struct {
	PackID      string
	Version     string
	SubpackName string
}

type generatedResourcePackExperiment struct {
	Name    string
	Enabled bool
}

type generatedResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateUUID          string
	WorldTemplateVersion       string
	ResourcePacks              []generatedResourcePack
}

type generatedResourcePackStack struct {
	TexturePackRequired    bool
	TexturePackList        []generatedResourcePackInstance
	BaseGameVersion        string
	Experiments            []generatedResourcePackExperiment
	ExperimentsEverToggled bool
	IncludeEditorPacks     bool
}

type generatedResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       []byte
	IsPremiumPack  bool
	PackType       uint8
}

type resourcePacksDocument struct {
	// Info and Stack mirror the session negotiation packets. They are kept
	// separate from Packs so consumers do not mistake connection flags or the
	// world-template identity for pack metadata.
	Info  resourcePacksInfoDocument `json:"info"`
	Packs []resourcePackDocument    `json:"packs"`
	Stack resourcePackStackDocument `json:"stack"`
	Data  *resourcePackDataDocument `json:"data_info,omitempty"`
}

type resourcePacksInfoDocument struct {
	ResourcePackRequired       bool                  `json:"resource_pack_required"`
	HasAddonPacks              bool                  `json:"has_addon_packs"`
	HasScripts                 bool                  `json:"has_scripts"`
	ForceDisableVibrantVisuals bool                  `json:"force_disable_vibrant_visuals"`
	WorldTemplate              resourcePackIDVersion `json:"world_template"`
}

type resourcePackIDVersion struct {
	UUID    string `json:"uuid"`
	Version string `json:"version"`
}

type resourcePackDocument struct {
	UUID                string `json:"uuid"`
	Version             string `json:"version"`
	Size                uint64 `json:"size"`
	ContentKey          string `json:"content_key"`
	SubpackName         string `json:"subpack_name"`
	ContentIdentity     string `json:"content_identity"`
	HasScripts          bool   `json:"has_scripts"`
	IsAddonPack         bool   `json:"is_addon_pack"`
	IsRayTracingCapable bool   `json:"is_ray_tracing_capable"`
	CDNURL              string `json:"cdn_url"`
}

type resourcePackStackDocument struct {
	TexturePackRequired    bool                             `json:"texture_pack_required"`
	Packs                  []resourcePackInstanceDocument   `json:"packs"`
	BaseGameVersion        string                           `json:"base_game_version"`
	Experiments            []resourcePackExperimentDocument `json:"experiments"`
	ExperimentsEverToggled bool                             `json:"experiments_ever_toggled"`
	IncludeEditorPacks     bool                             `json:"include_editor_packs"`
}

type resourcePackInstanceDocument struct {
	PackID      string `json:"pack_id"`
	Version     string `json:"version"`
	SubpackName string `json:"subpack_name"`
}

type resourcePackExperimentDocument struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

type resourcePackDataDocument struct {
	ResourceName   string `json:"resource_name"`
	ChunkSize      uint32 `json:"chunk_size"`
	NumberOfChunks uint32 `json:"number_of_chunks"`
	FileSize       uint64 `json:"file_size"`
	FileHash       string `json:"file_hash"`
	IsPremiumPack  bool   `json:"is_premium_pack"`
	PackType       uint8  `json:"pack_type"`
}

func resourcePackMetadata(infoData, stackData, dataInfoData []byte) ([]byte, error) {
	info, err := decodeResourcePacksInfo(infoData)
	if err != nil {
		return nil, fmt.Errorf("decode ResourcePacksInfoPacket with generated protocol: %w", err)
	}
	stack, err := decodeResourcePackStack(stackData)
	if err != nil {
		return nil, fmt.Errorf("decode ResourcePackStackPacket with generated protocol: %w", err)
	}

	packs := make([]resourcePackDocument, len(info.ResourcePacks))
	for i, pack := range info.ResourcePacks {
		packs[i] = resourcePackDocument{
			UUID:                pack.UUID,
			Version:             pack.Version,
			Size:                pack.Size,
			ContentKey:          pack.ContentKey,
			SubpackName:         pack.SubpackName,
			ContentIdentity:     pack.ContentIdentity,
			HasScripts:          pack.HasScripts,
			IsAddonPack:         pack.IsAddonPack,
			IsRayTracingCapable: pack.IsRayTracingCapable,
			CDNURL:              pack.CDNURL,
		}
	}
	stackPacks := make([]resourcePackInstanceDocument, len(stack.TexturePackList))
	for i, pack := range stack.TexturePackList {
		stackPacks[i] = resourcePackInstanceDocument{
			PackID:      pack.PackID,
			Version:     pack.Version,
			SubpackName: pack.SubpackName,
		}
	}
	experiments := make([]resourcePackExperimentDocument, len(stack.Experiments))
	for i, experiment := range stack.Experiments {
		experiments[i] = resourcePackExperimentDocument{Name: experiment.Name, Enabled: experiment.Enabled}
	}
	var dataInfoDocument *resourcePackDataDocument
	if len(dataInfoData) != 0 {
		dataInfo, err := decodeResourcePackDataInfo(dataInfoData)
		if err != nil {
			return nil, fmt.Errorf("decode ResourcePackDataInfoPacket with generated protocol: %w", err)
		}
		dataInfoDocument = &resourcePackDataDocument{
			ResourceName:   dataInfo.ResourceName,
			ChunkSize:      dataInfo.ChunkSize,
			NumberOfChunks: dataInfo.NumberOfChunks,
			FileSize:       dataInfo.FileSize,
			FileHash:       hex.EncodeToString(dataInfo.FileHash),
			IsPremiumPack:  dataInfo.IsPremiumPack,
			PackType:       dataInfo.PackType,
		}
	}

	document := resourcePacksDocument{
		Info: resourcePacksInfoDocument{
			ResourcePackRequired:       info.ResourcePackRequired,
			HasAddonPacks:              info.HasAddonPacks,
			HasScripts:                 info.HasScripts,
			ForceDisableVibrantVisuals: info.ForceDisableVibrantVisuals,
			WorldTemplate: resourcePackIDVersion{
				UUID:    info.WorldTemplateUUID,
				Version: info.WorldTemplateVersion,
			},
		},
		Packs: packs,
		Stack: resourcePackStackDocument{
			TexturePackRequired:    stack.TexturePackRequired,
			Packs:                  stackPacks,
			BaseGameVersion:        stack.BaseGameVersion,
			Experiments:            experiments,
			ExperimentsEverToggled: stack.ExperimentsEverToggled,
			IncludeEditorPacks:     stack.IncludeEditorPacks,
		},
		Data: dataInfoDocument,
	}
	return marshalCanonicalJSON(document)
}
