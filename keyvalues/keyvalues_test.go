package keyvalues

import (
	"bytes"
	"testing"
)

func assertEquals(found, expected interface{}, message string, t *testing.T) {
	if found != expected {
		t.Logf("Found:    %v", found)
		t.Logf("Expected: %v", expected)
		t.Error(message)
	}
}

func TestSimpleValues(t *testing.T) {
	var kv KeyValues

	assertEquals(kv.Int(57), int64(57), "Zero value of KeyValues should return defaults", t)
	kv.SetValueString("10")
	assertEquals(kv.Uint64(57), uint64(10), "KeyValues with value \"10\" should have an uint64 value of 10", t)
	assertEquals(kv.Int(57), int64(10), "KeyValues with value \"10\" should have an integer value of 10", t)
	assertEquals(kv.String("Bacon"), "10", "KeyValues with value \"10\" should have a string value of \"10\"", t)
	assertEquals(kv.Float(57), float64(10), "KeyValues with value \"10\" should have a float value of 10.0", t)
	assertEquals(kv.Bool(false), true, "KeyValues with value \"10\" should have a boolean value of true", t)
}

func TestComplexSerialize(t *testing.T) {
	var kv KeyValues

	lightmappedGeneric := kv.NewSubKey("LightmappedGeneric")
	lightmappedGeneric.NewSubKey("$basetexture").SetValueString("nature/dirtfloor001a")
	lightmappedGeneric.NewSubKey("$surfaceprop").SetValueString("dirt")
	lightmappedGeneric.NewSubKey("Proxy").NewSubKey("Test").NewSubKey("$key").SetValueString("Value")

	var buf bytes.Buffer
	kv.WriteTo(&buf)

	assertEquals(string(buf.Bytes()), `"LightmappedGeneric" {
	"$basetexture" "nature/dirtfloor001a"
	"$surfaceprop" "dirt"
	"Proxy" {
		"Test" {
			"$key" "Value"
		}
	}
}
`, "KeyValues serialization differs from expected", t)
}
