package runtimegen

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
	"testing"

	"github.com/bedrock-mc/protocolgen/data/item"
)

// runtimeFixture supplies independent registry, property and supplemental-tag inputs.
func runtimeFixture() (map[string]RegistryEntry, map[string]Properties, map[string][]string, Corrections) {
	return map[string]RegistryEntry{
			"minecraft:apple": {RuntimeID: 1, Version: 2},
			"minecraft:new":   {RuntimeID: 2, ComponentBased: true},
		}, map[string]Properties{
			"minecraft:apple": {MaxStackSize: 64},
			"minecraft:new":   {Tags: []string{"minecraft:harness"}},
		}, map[string][]string{
			"minecraft:allow_offhand": {"minecraft:apple"},
			"minecraft:head":          {"minecraft:apple"},
		}, Corrections{StackSizes: map[string]int{"minecraft:new": 16}, Offhand: []string{"minecraft:new"}, Evidence: map[string]string{"fixture": "synthetic"}}
}

func TestRuntimeCatalogCombinesSourcesDeterministically(t *testing.T) {
	registry, properties, tags, corrections := runtimeFixture()
	first, err := Generate(registry, properties, tags, corrections)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Generate(registry, properties, tags, corrections)
	if err != nil || !bytes.Equal(first, second) {
		t.Fatalf("generation is not deterministic: %v", err)
	}
	if !bytes.Contains(first, []byte(`Name: "minecraft:new", RuntimeID: 2, ComponentBased: true, Version: 0, MaxCount: 16, BodyWearable: true, AllowOffHand: true, ArmourSlot: 0, ArmourWearable: false`)) {
		t.Fatalf("registry, stack and capability inputs were not combined:\n%s", first)
	}
}

func TestRuntimeGeneratorHandlesEverySchemaField(t *testing.T) {
	r, p, tags, c := runtimeFixture()
	output, err := Generate(r, p, tags, c)
	if err != nil {
		t.Fatal(err)
	}
	file, err := parser.ParseFile(token.NewFileSet(), "runtime_generated.go", output, 0)
	if err != nil {
		t.Fatal(err)
	}
	schema := reflect.TypeFor[item.Runtime]()
	ast.Inspect(file, func(node ast.Node) bool {
		literal, ok := node.(*ast.CompositeLit)
		if !ok || literal.Type != nil {
			return true
		}
		fields := make(map[string]bool)
		for _, element := range literal.Elts {
			pair := element.(*ast.KeyValueExpr)
			fields[pair.Key.(*ast.Ident).Name] = true
		}
		for i := range schema.NumField() {
			if name := schema.Field(i).Name; !fields[name] {
				t.Errorf("generator silently omitted Runtime.%s", name)
			}
		}
		return false
	})
}

func TestRuntimeCatalogRejectsMissingAndStaleFacts(t *testing.T) {
	for _, scenario := range []struct {
		name string
		edit func(map[string]RegistryEntry, map[string]Properties, map[string][]string, *Corrections)
		want string
	}{
		{"missing size", func(_ map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, c *Corrections) {
			delete(c.StackSizes, "minecraft:new")
		}, "missing or invalid stack size"},
		{"stale size", func(_ map[string]RegistryEntry, p map[string]Properties, _ map[string][]string, _ *Corrections) {
			p["minecraft:new"] = Properties{MaxStackSize: 16}
		}, "is stale"},
		{"removed item", func(r map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, _ *Corrections) {
			delete(r, "minecraft:new")
		}, "unknown item"},
		{"stale tag", func(_ map[string]RegistryEntry, _ map[string]Properties, tags map[string][]string, _ *Corrections) {
			tags["minecraft:allow_offhand"] = append(tags["minecraft:allow_offhand"], "minecraft:new")
		}, "stale or duplicated"},
		{"missing tag", func(_ map[string]RegistryEntry, _ map[string]Properties, tags map[string][]string, _ *Corrections) {
			delete(tags, "minecraft:head")
		}, "missing or empty"},
		{"duplicate ID", func(r map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, _ *Corrections) {
			r["minecraft:apple"] = r["minecraft:new"]
		}, "share runtime ID"},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			r, p, tags, c := runtimeFixture()
			scenario.edit(r, p, tags, &c)
			if _, err := Generate(r, p, tags, c); err == nil || !strings.Contains(err.Error(), scenario.want) {
				t.Fatalf("error = %v, want %q", err, scenario.want)
			}
		})
	}
}
