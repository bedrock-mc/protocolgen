//go:build protocolgen_12640

package vanilladata

import (
	genprotocol "protocolgen/generated/1.26.40/go/protocol"
	genpacket "protocolgen/generated/1.26.40/go/protocol/packet"
)

func generatedProtocolTarget() (string, int) {
	return genprotocol.GAME_VERSION, genprotocol.PROTOCOL_VERSION
}

func decodeItemRegistry(data []byte) ([]generatedItemData, error) {
	var packet genpacket.ItemRegistry
	if err := genpacket.Decode(data, &packet); err != nil {
		return nil, err
	}
	items := make([]generatedItemData, len(packet.ItemData))
	for i, item := range packet.ItemData {
		items[i] = generatedItemData{
			ItemName:          item.ItemName,
			ItemID:            item.ItemID,
			IsComponentBased:  item.IsComponentBased,
			ItemVersion:       int32(item.ItemVersion),
			ItemComponentData: append([]byte(nil), item.ItemComponentData...),
		}
	}
	return items, nil
}

func decodeAvailableActorIdentifiers(data []byte) ([]byte, error) {
	var packet genpacket.AvailableActorIdentifiers
	if err := genpacket.Decode(data, &packet); err != nil {
		return nil, err
	}
	return append([]byte(nil), packet.IdentifierList...), nil
}

func decodeBiomeDefinitionList(data []byte) (generatedBiomeDefinitionList, error) {
	var packet genpacket.BiomeDefinitionList
	if err := genpacket.Decode(data, &packet); err != nil {
		return generatedBiomeDefinitionList{}, err
	}
	result := generatedBiomeDefinitionList{
		Entries: make([]generatedBiomeEntry, len(packet.MapOfBiomeNamesToData)),
		Strings: append([]string(nil), packet.StringList.Strings...),
	}
	for i, entry := range packet.MapOfBiomeNamesToData {
		value := entry.Value
		converted := generatedBiomeEntry{
			Key:           entry.Key,
			ID:            value.ID,
			Temperature:   value.Temperature,
			Downfall:      value.Downfall,
			FoliageSnow:   value.FoliageSnow,
			Depth:         value.Depth,
			Scale:         value.Scale,
			MapWaterColor: value.MapWaterColorARGB,
			Rain:          value.Rain,
		}
		if tags, ok := value.Tags.Value(); ok {
			converted.Tags = append([]uint16(nil), tags.Tags...)
		}
		result.Entries[i] = converted
	}
	return result, nil
}

func decodeResourcePacksInfo(data []byte) (generatedResourcePacksInfo, error) {
	var packet genpacket.ResourcePacksInfo
	if err := genpacket.Decode(data, &packet); err != nil {
		return generatedResourcePacksInfo{}, err
	}
	result := generatedResourcePacksInfo{
		ResourcePackRequired:       packet.ResourcePackRequired,
		HasAddonPacks:              packet.HasAddonPacks,
		HasScripts:                 packet.HasScripts,
		ForceDisableVibrantVisuals: packet.ForceDisableVibrantVisuals,
		WorldTemplateUUID:          packet.WorldTemplateIDAndVersion.PackUUID.String(),
		WorldTemplateVersion:       packet.WorldTemplateIDAndVersion.PackVersion.Version,
		ResourcePacks:              make([]generatedResourcePack, len(packet.ResourcePacks)),
	}
	for i, pack := range packet.ResourcePacks {
		result.ResourcePacks[i] = generatedResourcePack{
			UUID:                pack.PackIDVersion.PackUUID.String(),
			Version:             pack.PackIDVersion.PackVersion.Version,
			Size:                pack.PackSize,
			ContentKey:          pack.ContentKey,
			SubpackName:         pack.SubpackName,
			ContentIdentity:     pack.ContentIdentity.Identity,
			HasScripts:          pack.HasScripts,
			IsAddonPack:         pack.IsAddonPack,
			IsRayTracingCapable: pack.IsRayTracingCapable,
			CDNURL:              pack.CDNURL,
		}
	}
	return result, nil
}

func decodeResourcePackStack(data []byte) (generatedResourcePackStack, error) {
	var packet genpacket.ResourcePackStack
	if err := genpacket.Decode(data, &packet); err != nil {
		return generatedResourcePackStack{}, err
	}
	result := generatedResourcePackStack{
		TexturePackRequired:    packet.TexturePackRequired,
		TexturePackList:        make([]generatedResourcePackInstance, len(packet.TexturePackList)),
		BaseGameVersion:        packet.BaseGameVersion,
		Experiments:            make([]generatedResourcePackExperiment, len(packet.Experiments.Toggles)),
		ExperimentsEverToggled: packet.Experiments.ExperimentsEverToggled,
		IncludeEditorPacks:     packet.IncludeEditorPacks,
	}
	for i, pack := range packet.TexturePackList {
		result.TexturePackList[i] = generatedResourcePackInstance{PackID: pack.PackID, Version: pack.Version, SubpackName: pack.SubPackName}
	}
	for i, experiment := range packet.Experiments.Toggles {
		result.Experiments[i] = generatedResourcePackExperiment{Name: experiment.Name, Enabled: experiment.Enabled}
	}
	return result, nil
}

func decodeResourcePackDataInfo(data []byte) (generatedResourcePackDataInfo, error) {
	var packet genpacket.ResourcePackDataInfo
	if err := genpacket.Decode(data, &packet); err != nil {
		return generatedResourcePackDataInfo{}, err
	}
	return generatedResourcePackDataInfo{
		ResourceName:   packet.ResourceName,
		ChunkSize:      packet.ChunkSize,
		NumberOfChunks: packet.NumberOfChunks,
		FileSize:       packet.FileSize,
		FileHash:       append([]byte(nil), packet.FileHash...),
		IsPremiumPack:  packet.IsPremiumPack,
		PackType:       packet.PackType,
	}, nil
}
