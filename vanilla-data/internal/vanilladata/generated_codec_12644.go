//go:build !protocolgen_12640

package vanilladata

import (
	genprotocol "protocolgen/generated/1.26.44/go/protocol"
	genpacket "protocolgen/generated/1.26.44/go/protocol/packet"
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
