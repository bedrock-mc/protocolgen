// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BiomeDefinitionList struct {
	MapOfBiomeNamesToData []OrderedEntry[uint16, BiomeDefinitionData]
	StringList            BiomeStringList
}
