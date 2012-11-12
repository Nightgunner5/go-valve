package keyvalues

import (
	"strings"
	"testing"
)

func TestParseMaterial(t *testing.T) {
	var kv KeyValues
	if _, err := kv.ReadFrom(strings.NewReader(`LightmappedGeneric{$basetexture"nature/dirtfloor001a""$surfaceprop"dirt}`)); err != nil {
		t.Error(err)
	}

	first := true
	for key := range kv.Each() {
		if first {
			first = false
		} else {
			t.Errorf("Unexpected key %q", key.Name())
			continue
		}

		assertEquals(key.Name(), "LightmappedGeneric", "root key name", t)
		count := 0
		for subkey := range key.Each() {
			switch count {
			case 0:
				assertEquals(subkey.Name(), "$basetexture", "first subkey name", t)
				assertEquals(subkey.String("UNDEFINED"), "nature/dirtfloor001a", "first subkey value", t)
			case 1:
				assertEquals(subkey.Name(), "$surfaceprop", "second subkey name", t)
				assertEquals(subkey.String("UNDEFINED"), "dirt", "second subkey value", t)
			default:
				t.Errorf("Unexpected subkey %q", subkey.Name())
			}
			count++
		}
	}
	if first {
		t.Error("Missing LightmappedGeneric key!")
	}
}
