package keyvalues

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseMaterial(t *testing.T) {
	var kv KeyValues
	if _, err := kv.ReadFrom(strings.NewReader(`LightmappedGeneric{$basetexture"nature/dirtfloor001a""$surfaceprop"dirt}`)); err != nil {
		t.Error(err)
	}

	var buf bytes.Buffer
	kv.WriteTo(&buf)

	assertEquals(string(buf.Bytes()), `"LightmappedGeneric" {
	"$basetexture" "nature/dirtfloor001a"
	"$surfaceprop" "dirt"
}
`, "serialize mismatch", t)
}

func TestParseComments(t *testing.T) {
	var kv KeyValues
	if _, err := kv.ReadFrom(strings.NewReader(`// Comment
//* comment
Key {
//*/
	subkey // comment
	value /*comment*/
	// comment
}
`)); err != nil {
		t.Error(err)
	}

	var buf bytes.Buffer
	kv.WriteTo(&buf)

	assertEquals(string(buf.Bytes()), `"Key" {
	"subkey" "value"
}
`, "serialize mismatch", t)
}
