package codegen

import (
	"fmt"
	"strings"
)

// genField is a documented packet field with everything the generator needs.
type genField struct {
	DocName   string
	GoName    string
	Ordinal   int
	Under     string // x-underlying-type (flattened from a single-scalar ref)
	Options   []string
	JSONType  string
	Enum      bool
	Required  bool
	Ref       string
	RefTitle  string
	Composite bool // a $ref to a multi-field definition
	IsArray   bool
	Elem      *genField // element shape for arrays
	OneOf     bool
}

// fieldShape is the resolved Go representation of a field.
type fieldShape struct {
	GoType  string
	Marshal string   // fully-rendered Marshal statement (receiver "pk", io var "io")
	Todo    string   // non-empty when hand-finishing is needed
	Imports []string // extra import paths required
}

// compositeMapping maps a known Mojang composite title to its gophertunnel Go
// type, the io method that encodes it, and any extra import needed.
type compositeMapping struct {
	goType string
	ioOp   string
	imp    string
}

// knownComposites maps a (cleaned) Mojang composite title to a gophertunnel type.
// An entry with a non-empty ioOp is encoded via that io method (io.Vec3); an
// entry with an empty ioOp is a protocol.Marshaler type encoded via
// protocol.Single (one) / protocol.Slice (array). See compositeProtocolTypes for
// the verified protocol.* mappings appended in resolve.go's init.
var knownComposites = map[string]compositeMapping{
	"Vec3":                 {"mgl32.Vec3", "Vec3", "github.com/go-gl/mathgl/mgl32"},
	"Vec2":                 {"mgl32.Vec2", "Vec2", "github.com/go-gl/mathgl/mgl32"},
	"BlockPos":             {"protocol.BlockPos", "BlockPos", ""},
	"NetworkBlockPosition": {"protocol.BlockPos", "BlockPos", ""},
	"ChunkPos":             {"protocol.ChunkPos", "ChunkPos", ""},
	"SubChunkPos":          {"protocol.SubChunkPos", "SubChunkPos", ""},
	"UUID":                 {"uuid.UUID", "UUID", "github.com/google/uuid"},
	// NetworkItemStackDescriptor is the network item-with-stack-id type, encoded
	// via the dedicated io.ItemInstance method (not a protocol.Marshaler slice).
	"NetworkItemStackDescriptor": {"protocol.ItemInstance", "ItemInstance", ""},
}

// compositeProtocolTypes maps a doc composite title to an existing gophertunnel
// protocol.Marshaler type, verified by structural (field-by-field wire) comparison
// against the gopher type's Marshal method. Encoded via protocol.Single (one) /
// protocol.Slice (array). Keys are cleaned with cleanCompositeTitle in init, so
// raw doc titles (with spaces/namespaces/casing) may be written here directly.
var compositeProtocolTypes = map[string]string{
	"MissingBlobData":            "CacheBlob",
	"ActorLink":                  "EntityLink",
	"AimAssistActorPriorityData": "CameraAimAssistActorPriorityData",
	"AttributeData":              "Attribute",
	"AvailableCommandsPacketChainedSubcommandData": "ChainedSubcommand",
	"AvailableCommandsPacketConstrainedValueData":  "CommandEnumConstraint",
	"AvailableCommandsPacketSoftEnumData":          "DynamicEnum",
	"CameraAimAssistCategoryDefinition":            "CameraAimAssistCategory",
	"CameraAimAssistPresetDefinition":              "CameraAimAssistPreset",
	"CameraSplineDefinition":                       "CameraSplineDefinition",
	"clientStoreEntryPointConfig":                  "StoreEntryPointInfo",
	"ContainerMixDataEntry":                        "PotionContainerChangeRecipe",
	"CreativeGroupInfoPayload":                     "CreativeGroup",
	"CreativeItemEntryPayload":                     "CreativeItem",
	"Data Store Update":                            "DataStoreUpdate",
	"EduSharedUriResource":                         "EducationSharedResourceURI",
	"EntityDiagnosticTimingInfo":                   "EntityDiagnosticTimingInfo",
	"gatheringsConfig":                             "GatheringJoinInfo",
	"ItemData":                                     "ItemEntry",
	"ItemEnchantOption":                            "EnchantmentOption",
	"ItemStackRequest":                             "ItemStackRequest",
	"ItemStackResponseInfo":                        "ItemStackResponse",
	"LegacySetSlot":                                "LegacySetItemSlot",
	"LocatorBarWaypointPayload":                    "LocatorBarWaypoint",
	"MapDecoration":                                "MapDecoration",
	"MapItemTrackedActor UniqueId":                 "MapTrackedObject",
	"MemoryCategoryCounter":                        "MemoryCategoryCounter",
	"PackInstanceId":                               "StackResourcePack",
	"PlayerBlockActionData":                        "PlayerBlockAction",
	"PotionMixDataEntry":                           "PotionRecipe",
	"PropertySyncData":                             "EntityProperties",
	"ScoreboardIdentityPacketInfo":                 "ScoreboardIdentityEntry",
	"SerializedAbilitiesData":                      "AbilityData",
	"ServerBlockProperty":                          "BlockEntry",
	"ShapedRecipePayload":                          "ShapedRecipe",
	"ShapelessRecipePayload":                       "ShapelessRecipe",
	"SmithingTransformRecipePayload":               "SmithingTransformRecipe",
	"SmithingTrimRecipePayload":                    "SmithingTrimRecipe",
	"StructureSettings":                            "StructureSettings",
	"SubChunkPosOffset":                            "SubChunkOffset",
	"SyncedAttribute":                              "AttributeValue",
	"SystemDiagnosticTimingInfo":                   "SystemDiagnosticTimingInfo",
	"TrimMaterial":                                 "TrimMaterial",
	"TrimPattern":                                  "TrimPattern",
	"WhiskerScopeDataSummary":                      "WhiskerScopeDataSummary",
}

func init() {
	for title, typ := range compositeProtocolTypes {
		knownComposites[cleanCompositeTitle(title)] = compositeMapping{goType: "protocol." + typ}
	}
}

// resolve maps a documented field onto its Go type and Marshal statement.
func resolve(f genField) fieldShape {
	ref := "&pk." + f.GoName

	if f.OneOf {
		return fieldShape{
			GoType:  "any",
			Marshal: fmt.Sprintf("// TODO(protocoldrift): oneOf union %q — hand-write the discriminated marshal", f.GoName),
			Todo:    "oneOf union",
		}
	}

	if f.IsArray {
		return resolveArray(f)
	}

	if f.Composite {
		title := cleanCompositeTitle(f.RefTitle)
		if m, ok := knownComposites[title]; ok {
			if m.ioOp == "" { // protocol.Marshaler type
				return fieldShape{GoType: m.goType, Marshal: fmt.Sprintf("protocol.Single(io, %s)", ref), Imports: impList(m.imp)}
			}
			return fieldShape{GoType: m.goType, Marshal: fmt.Sprintf("io.%s(%s)", m.ioOp, ref), Imports: impList(m.imp)}
		}
		return fieldShape{
			GoType:  title,
			Marshal: fmt.Sprintf("protocol.Single(io, %s) // TODO(protocoldrift): composite %q — map to existing protocol type or generate it", ref, f.RefTitle),
			Todo:    "composite " + f.RefTitle,
		}
	}

	if isNBT(f) {
		return fieldShape{
			GoType:  "map[string]any",
			Marshal: fmt.Sprintf("io.NBT(%s, nbt.NetworkLittleEndian) // TODO(protocoldrift): docs model NBT as a string blob — verify encoding", ref),
			Todo:    "nbt blob",
			Imports: []string{"github.com/sandertv/gophertunnel/minecraft/nbt"},
		}
	}

	// Enum serialized as its string name (no Enum-as-Value option).
	if f.Enum && !hasOption(f.Options, "Enum-as-Value") {
		return fieldShape{
			GoType:  "string",
			Marshal: fmt.Sprintf("io.String(%s) // TODO(protocoldrift): enum serialized as string name — add value<->string conversion", ref),
			Todo:    "enum-as-string",
		}
	}

	goType, ioOp, ok := scalarIO(f.Under, f.JSONType, f.Options)
	if !ok {
		return fieldShape{
			GoType:  "any",
			Marshal: fmt.Sprintf("// TODO(protocoldrift): unmapped type %q for %q", f.Under, f.GoName),
			Todo:    "unmapped type " + f.Under,
		}
	}
	shape := fieldShape{GoType: goType, Marshal: fmt.Sprintf("io.%s(%s)", ioOp, ref)}
	if !f.Required {
		shape.Marshal += " // TODO(protocoldrift): not in required — may be a protocol.Optional on the wire"
		shape.Todo = "optional"
	}
	return shape
}

// resolveArray resolves an array field into a slice type and a slice marshal call.
func resolveArray(f genField) fieldShape {
	elem := *f.Elem
	if elem.Composite {
		title := cleanCompositeTitle(elem.RefTitle)
		if m, ok := knownComposites[title]; ok {
			if m.ioOp == "" { // protocol.Marshaler element
				return fieldShape{
					GoType:  "[]" + m.goType,
					Marshal: fmt.Sprintf("protocol.Slice(io, &pk.%s)", f.GoName),
					Imports: impList(m.imp),
				}
			}
			return fieldShape{
				GoType:  "[]" + m.goType,
				Marshal: fmt.Sprintf("protocol.FuncSlice(io, &pk.%s, io.%s)", f.GoName, m.ioOp),
				Imports: impList(m.imp),
			}
		}
		return fieldShape{
			GoType:  "[]" + title,
			Marshal: fmt.Sprintf("protocol.Slice(io, &pk.%s) // TODO(protocoldrift): element composite %q — ensure it implements Marshaler", f.GoName, elem.RefTitle),
			Todo:    "array of composite " + elem.RefTitle,
		}
	}
	goType, ioOp, ok := scalarIO(elem.Under, elem.JSONType, elem.Options)
	if !ok {
		return fieldShape{
			GoType:  "[]any",
			Marshal: fmt.Sprintf("// TODO(protocoldrift): array of unmapped type %q for %q", elem.Under, f.GoName),
			Todo:    "array of unmapped type",
		}
	}
	return fieldShape{
		GoType:  "[]" + goType,
		Marshal: fmt.Sprintf("protocol.FuncSlice(io, &pk.%s, io.%s)", f.GoName, ioOp),
	}
}

// scalarIO returns the Go type, io method, and ok for a documented scalar. When
// x-underlying-type is absent (common for plain strings/bools), it falls back to
// the JSON Schema type.
func scalarIO(under, jsonType string, options []string) (goType, ioOp string, ok bool) {
	if under == "" {
		switch jsonType {
		case "string":
			under = "string"
		case "boolean":
			under = "boolean"
		case "number":
			under = "float"
			// "integer" without an underlying type is ambiguous (varint vs fixed,
			// width unknown) — leave it unmapped for a human to resolve.
		}
	}
	comp := hasOption(options, "Compression")
	switch under {
	case "bool", "boolean":
		return "bool", "Bool", true
	case "int8":
		return "int8", "Int8", true
	case "uint8":
		return "uint8", "Uint8", true
	case "int16":
		return "int16", "Int16", true
	case "uint16":
		return "uint16", "Uint16", true
	case "int32":
		if comp {
			return "int32", "Varint32", true
		}
		return "int32", "Int32", true
	case "uint32":
		if comp {
			return "uint32", "Varuint32", true
		}
		return "uint32", "Uint32", true
	case "int64":
		if comp {
			return "int64", "Varint64", true
		}
		return "int64", "Int64", true
	case "uint64":
		if comp {
			return "uint64", "Varuint64", true
		}
		return "uint64", "Uint64", true
	case "float":
		return "float32", "Float32", true
	case "double":
		return "float64", "Float64", true
	case "string":
		return "string", "String", true
	default:
		return "", "", false
	}
}

// isNBT heuristically detects an NBT blob (docs model these as string fields).
func isNBT(f genField) bool {
	if f.Under != "string" && f.JSONType != "string" {
		return false
	}
	name := strings.ToLower(f.DocName + " " + f.GoName)
	return strings.Contains(name, "nbt") || strings.Contains(name, "tag")
}

func hasOption(options []string, opt string) bool {
	for _, o := range options {
		if o == opt {
			return true
		}
	}
	return false
}

// impList returns a single-element import slice, or nil for the empty path.
func impList(path string) []string {
	if path == "" {
		return nil
	}
	return []string{path}
}
