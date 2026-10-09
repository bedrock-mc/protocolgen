// Package entity provides generated values extracted from vanilla Bedrock entity
// definitions. It intentionally does not model events, filters, goals, or
// component-group transitions.
package entity

import "sort"

// Float is an optional floating-point value.
type Float struct {
	Value   float32
	Present bool
}

// Bool is an optional boolean value.
type Bool struct {
	Value   bool
	Present bool
}

// Number is either a fixed value or a random range from a vanilla component.
type Number struct {
	Value   float32
	Min     float32
	Max     float32
	Ranged  bool
	Present bool
}

// CollisionBox contains the base collision dimensions of an entity.
type CollisionBox struct {
	Width   float32
	Height  float32
	Present bool
}

// Health contains the initial health expression and an optional explicit
// maximum. A ranged initial value is represented by Initial.Min/Max.
type Health struct {
	Initial Number
	Maximum Float
	Present bool
}

// Entity contains stable values from an entity's base component set.
type Entity struct {
	Name              string
	SpawnCategory     string
	Experimental      bool
	Summonable        bool
	Spawnable         bool
	CollisionBox      CollisionBox
	Health            Health
	Movement          Number
	Scale             Number
	ProjectileGravity Float
	ProjectilePower   Float
	ProjectileBounce  Bool
	Pushable          Bool
	PushableByPiston  Bool
	HasPhysics        bool
	Families          []string
}

// All returns every entity specification in name order. The returned data must
// be treated as immutable.
func All() []Entity {
	return all
}

// Lookup finds an entity specification by its namespaced identifier.
func Lookup(name string) (Entity, bool) {
	index := sort.Search(len(all), func(i int) bool { return all[i].Name >= name })
	if index == len(all) || all[index].Name != name {
		return Entity{}, false
	}
	return all[index], true
}
