// Package voxelshape provides generated shapes from the vanilla Bedrock shape pack.
package voxelshape

import "sort"

// Box is an axis-aligned box in block-local coordinates.
type Box [6]float32

// Shape contains one named vanilla voxel shape.
type Shape struct {
	Name  string
	Boxes []Box
}

// All returns every shape in name order. The returned data must be treated as
// immutable.
func All() []Shape {
	return all
}

// Lookup finds a named shape.
func Lookup(name string) (Shape, bool) {
	index := sort.Search(len(all), func(i int) bool { return all[i].Name >= name })
	if index == len(all) || all[index].Name != name {
		return Shape{}, false
	}
	return all[index], true
}
