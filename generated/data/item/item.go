// Package item provides generated values extracted from vanilla Bedrock item
// definitions.
package item

import "sort"

// Effect describes an effect granted when a food item is consumed.
type Effect struct {
	Name      string
	Chance    float32
	Duration  float32
	Amplifier int32
}

// Food contains static food and use-duration properties.
type Food struct {
	Name               string
	Nutrition          int32
	SaturationModifier string
	UseDuration        float32
	CanAlwaysEat       bool
	UsingConvertsTo    string
	Effects            []Effect
}

// Foods returns every food specification in name order. The returned data must
// be treated as immutable.
func Foods() []Food {
	return all
}

// LookupFood finds food properties by a namespaced item identifier.
func LookupFood(name string) (Food, bool) {
	index := sort.Search(len(all), func(i int) bool { return all[i].Name >= name })
	if index == len(all) || all[index].Name != name {
		return Food{}, false
	}
	return all[index], true
}
