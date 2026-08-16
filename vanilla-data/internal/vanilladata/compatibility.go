package vanilladata

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

// These small value types keep the derived-data normalizer independent of a
// particular generated snapshot. The build-tagged codec adapters populate
// them from the exact generated packet package selected for the capture.
type generatedItemData struct {
	ItemName          string
	ItemID            int16
	IsComponentBased  bool
	ItemVersion       int32
	ItemComponentData []byte
}

type generatedBiomeEntry struct {
	Key           uint16
	ID            uint16
	Temperature   float32
	Downfall      float32
	FoliageSnow   float32
	Depth         float32
	Scale         float32
	MapWaterColor int32
	Rain          bool
	Tags          []uint16
}

type generatedBiomeDefinitionList struct {
	Entries []generatedBiomeEntry
	Strings []string
}

// ValidateGeneratedTarget ensures the derived-data exporter is compiled
// against the generated packet set selected by the capture manifest.
func ValidateGeneratedTarget(target Target) error {
	generatedVersion, generatedProtocol := generatedProtocolTarget()
	if target.MinecraftVersion != generatedVersion || target.ProtocolVersion != generatedProtocol {
		return fmt.Errorf("derived-data exporter uses generated protocol %s/%d, capture target is %s/%d", generatedVersion, generatedProtocol, target.MinecraftVersion, target.ProtocolVersion)
	}
	return nil
}

// BuildDerivedArtifacts decodes captured wire bodies with protocolgen's
// generated packet definitions and emits established vanilla-data formats
// that can be derived exactly from login packets alone.
func BuildDerivedArtifacts(payloads map[string][]byte) (map[string][]byte, error) {
	files := make(map[string][]byte)
	if data, ok := payloads["ItemRegistryPacket"]; ok {
		items, err := decodeItemRegistry(data)
		if err != nil {
			return nil, err
		}
		required, err := requiredItemList(items)
		if err != nil {
			return nil, err
		}
		files["required_item_list.json"] = required
	}
	if data, ok := payloads["AvailableActorIdentifiersPacket"]; ok {
		identifierList, err := decodeAvailableActorIdentifiers(data)
		if err != nil {
			return nil, err
		}
		entityMap, err := entityIDMap(identifierList)
		if err != nil {
			return nil, err
		}
		files["entity_id_map.json"] = entityMap
		files["entity_identifiers.nbt"] = append([]byte(nil), identifierList...)
	}
	if data, ok := payloads["BiomeDefinitionListPacket"]; ok {
		biomes, err := decodeBiomeDefinitionList(data)
		if err != nil {
			return nil, err
		}
		definitions, err := biomeDefinitions(&biomes)
		if err != nil {
			return nil, err
		}
		files["biome_definitions.json"] = definitions
	}
	return files, nil
}

type requiredItemEntry struct {
	RuntimeID      int16  `json:"runtime_id"`
	ComponentBased bool   `json:"component_based"`
	Version        int32  `json:"version"`
	ComponentNBT   string `json:"component_nbt,omitempty"`
}

func requiredItemList(items []generatedItemData) ([]byte, error) {
	result := make(map[string]requiredItemEntry, len(items))
	for _, item := range items {
		entry := requiredItemEntry{RuntimeID: item.ItemID, ComponentBased: item.IsComponentBased, Version: int32(item.ItemVersion)}
		if len(item.ItemComponentData) != 0 {
			var value map[string]any
			if err := nbt.UnmarshalEncoding(item.ItemComponentData, &value, nbt.NetworkLittleEndian); err != nil {
				return nil, fmt.Errorf("decode component NBT for %s: %w", item.ItemName, err)
			}
			if len(value) != 0 {
				persistent, err := marshalPersistentNBT(value)
				if err != nil {
					return nil, fmt.Errorf("encode component NBT for %s: %w", item.ItemName, err)
				}
				entry.ComponentNBT = base64.StdEncoding.EncodeToString(persistent)
			}
		}
		result[item.ItemName] = entry
	}
	return marshalCanonicalJSON(result)
}

// marshalPersistentNBT uses sorted compound keys. The gophertunnel NBT
// encoder accepts maps but deliberately follows Go's random map iteration,
// which would make checked-in component_nbt strings change between runs.
func marshalPersistentNBT(value map[string]any) ([]byte, error) {
	var out bytes.Buffer
	if err := writeNBTTag(&out, 10, "", value); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func writeNBTTag(out *bytes.Buffer, tag byte, name string, value any) error {
	out.WriteByte(tag)
	if err := writeNBTString(out, name); err != nil {
		return err
	}
	return writeNBTPayload(out, tag, reflect.ValueOf(value))
}

func writeNBTPayload(out *bytes.Buffer, tag byte, value reflect.Value) error {
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	switch tag {
	case 1:
		out.WriteByte(byte(value.Uint()))
	case 2:
		writeLE(out, uint16(value.Int()))
	case 3:
		writeLE(out, uint32(value.Int()))
	case 4:
		writeLE(out, uint64(value.Int()))
	case 5:
		writeLE(out, math.Float32bits(float32(value.Float())))
	case 6:
		writeLE(out, math.Float64bits(value.Float()))
	case 7:
		writeLE(out, uint32(value.Len()))
		for i := 0; i < value.Len(); i++ {
			out.WriteByte(byte(value.Index(i).Uint()))
		}
	case 8:
		return writeNBTString(out, value.String())
	case 9:
		elementTag := byte(0)
		if value.Len() != 0 {
			var err error
			elementTag, err = nbtTagOf(value.Index(0))
			if err != nil {
				return err
			}
		}
		out.WriteByte(elementTag)
		writeLE(out, uint32(value.Len()))
		for i := 0; i < value.Len(); i++ {
			currentTag, err := nbtTagOf(value.Index(i))
			if err != nil {
				return err
			}
			if currentTag != elementTag {
				return fmt.Errorf("NBT list contains tag %d after tag %d", currentTag, elementTag)
			}
			if err := writeNBTPayload(out, elementTag, value.Index(i)); err != nil {
				return err
			}
		}
	case 10:
		keys := value.MapKeys()
		sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
		for _, key := range keys {
			entry := value.MapIndex(key)
			entryTag, err := nbtTagOf(entry)
			if err != nil {
				return fmt.Errorf("NBT key %s: %w", key.String(), err)
			}
			if err := writeNBTTag(out, entryTag, key.String(), entry.Interface()); err != nil {
				return err
			}
		}
		out.WriteByte(0)
	case 11:
		writeLE(out, uint32(value.Len()))
		for i := 0; i < value.Len(); i++ {
			writeLE(out, uint32(value.Index(i).Int()))
		}
	case 12:
		writeLE(out, uint32(value.Len()))
		for i := 0; i < value.Len(); i++ {
			writeLE(out, uint64(value.Index(i).Int()))
		}
	default:
		return fmt.Errorf("unsupported NBT tag %d", tag)
	}
	return nil
}

func nbtTagOf(value reflect.Value) (byte, error) {
	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			return 0, fmt.Errorf("nil NBT value")
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Uint8:
		return 1, nil
	case reflect.Int16:
		return 2, nil
	case reflect.Int32:
		return 3, nil
	case reflect.Int64:
		return 4, nil
	case reflect.Float32:
		return 5, nil
	case reflect.Float64:
		return 6, nil
	case reflect.String:
		return 8, nil
	case reflect.Slice:
		return 9, nil
	case reflect.Map:
		return 10, nil
	case reflect.Array:
		switch value.Type().Elem().Kind() {
		case reflect.Uint8:
			return 7, nil
		case reflect.Int32:
			return 11, nil
		case reflect.Int64:
			return 12, nil
		}
	}
	return 0, fmt.Errorf("unsupported NBT Go type %s", value.Type())
}

func writeNBTString(out *bytes.Buffer, value string) error {
	if len(value) > math.MaxUint16 {
		return fmt.Errorf("NBT string is too long: %d", len(value))
	}
	writeLE(out, uint16(len(value)))
	out.WriteString(value)
	return nil
}

func writeLE(out *bytes.Buffer, value any) { _ = binary.Write(out, binary.LittleEndian, value) }

type actorIdentifierList struct {
	IDList []actorIdentifier `nbt:"idlist"`
}

type actorIdentifier struct {
	BID         string `nbt:"bid"`
	HasSpawnEgg bool   `nbt:"hasspawnegg"`
	ID          string `nbt:"id"`
	RuntimeID   int32  `nbt:"rid"`
	Summonable  bool   `nbt:"summonable"`
}

func entityIDMap(encoded []byte) ([]byte, error) {
	var identifiers actorIdentifierList
	if err := nbt.UnmarshalEncoding(encoded, &identifiers, nbt.NetworkLittleEndian); err != nil {
		return nil, fmt.Errorf("decode actor identifiers NBT: %w", err)
	}
	sort.SliceStable(identifiers.IDList, func(i, j int) bool { return identifiers.IDList[i].RuntimeID < identifiers.IDList[j].RuntimeID })
	var out bytes.Buffer
	out.WriteString("{\n")
	for i, entry := range identifiers.IDList {
		name, _ := json.Marshal(entry.ID)
		fmt.Fprintf(&out, "    %s: %d", name, entry.RuntimeID)
		if i+1 != len(identifiers.IDList) {
			out.WriteByte(',')
		}
		out.WriteByte('\n')
	}
	out.WriteString("}\n")
	return out.Bytes(), nil
}

type canonicalColour struct {
	A uint8 `json:"a"`
	B uint8 `json:"b"`
	G uint8 `json:"g"`
	R uint8 `json:"r"`
}

type canonicalBiomeDefinition struct {
	Depth          float64         `json:"depth"`
	Downfall       float64         `json:"downfall"`
	FoliageSnow    float64         `json:"foliageSnow"`
	ID             uint16          `json:"id"`
	MapWaterColour canonicalColour `json:"mapWaterColour"`
	Rain           bool            `json:"rain"`
	Scale          float64         `json:"scale"`
	Tags           []string        `json:"tags"`
	Temperature    float64         `json:"temperature"`
}

func biomeDefinitions(pk *generatedBiomeDefinitionList) ([]byte, error) {
	strings := pk.Strings
	result := make(map[string]canonicalBiomeDefinition, len(pk.Entries))
	for _, entry := range pk.Entries {
		if int(entry.Key) >= len(strings) {
			return nil, fmt.Errorf("biome name index %d exceeds string list", entry.Key)
		}
		tags := []string{}
		for _, index := range entry.Tags {
			if int(index) >= len(strings) {
				return nil, fmt.Errorf("biome tag index %d exceeds string list", index)
			}
			tags = append(tags, strings[index])
		}
		colour := uint32(entry.MapWaterColor)
		result[strings[entry.Key]] = canonicalBiomeDefinition{
			Depth: round3(entry.Depth), Downfall: round3(entry.Downfall), FoliageSnow: round3(entry.FoliageSnow), ID: entry.ID,
			MapWaterColour: canonicalColour{A: uint8(colour >> 24), R: uint8(colour >> 16), G: uint8(colour >> 8), B: uint8(colour)},
			Rain:           entry.Rain, Scale: round3(entry.Scale), Tags: tags, Temperature: round3(entry.Temperature),
		}
	}
	return marshalCanonicalJSON(result)
}

func round3(value float32) float64 { return math.Round(float64(value)*1000) / 1000 }

func marshalCanonicalJSON(value any) ([]byte, error) {
	data, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}
