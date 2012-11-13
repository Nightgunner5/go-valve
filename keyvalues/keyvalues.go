package keyvalues

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyValues struct {
	name         string
	simpleValue  string // Only used if complexValue is nil
	complexValue []KeyValues
}

func (kv KeyValues) Name() string {
	return kv.name
}

// Returns the string value of this node. If the node is nonexistent or complex
// (one that has subnodes) the default (def) will be returned.
func (kv *KeyValues) String(def string) string {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	return kv.simpleValue
}

// Returns the integer value of this node. If the node is nonexistent, complex,
// or unable to be parsed as an integer, the default (def) will be returned.
func (kv *KeyValues) Int(def int64) int64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if i, err := strconv.ParseInt(kv.simpleValue, 0, 64); err == nil {
		return i
	}
	return def
}

// Returns the uint64 value of this node. If the node is nonexistent, complex,
// or unable to be parsed as a uint64, the default (def) will be returned.
func (kv *KeyValues) Uint64(def uint64) uint64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if i, err := strconv.ParseUint(kv.simpleValue, 0, 64); err == nil {
		return i
	}
	return def
}

// Returns the floating-point value of this node. If the node is nonexistent,
// complex, or unable to be parsed as a float, the default (def) will be returned.
func (kv *KeyValues) Float(def float64) float64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if f, err := strconv.ParseFloat(kv.simpleValue, 64); err == nil {
		return f
	}
	return def
}

// Returns the boolean value of this node. The boolean value is equivelent to
// false iff the node has the integer value 0.
func (kv *KeyValues) Bool(def bool) bool {
	if def {
		return kv.Int(1) != 0
	}
	return kv.Int(0) != 0
}

// Sets the value of this node. If this node is complex or nonexistent, this
// method will panic.
func (kv *KeyValues) SetValueString(v string) {
	if kv == nil {
		panic("SetValueString on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueString on a complex *KeyValues")
	}

	kv.simpleValue = v
}

// Sets the value of this node. If this node is complex or nonexistent, this
// method will panic. The integer will be formatted in base 10.
func (kv *KeyValues) SetValueInt(v int64) {
	if kv == nil {
		panic("SetValueInt on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueInt on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprint(v)
}

// Sets the value of this node. If this node is complex or nonexistent, this
// method will panic. The uint64 will be formatted as a hexadecimal number
// prefixed by "0x".
func (kv *KeyValues) SetValueUint64(v uint64) {
	if kv == nil {
		panic("SetValueUint64 on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueUint64 on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprintf("0x%x", v)
}

// Sets the value of this node. If this node is complex or nonexistent, this
// method will panic.
func (kv *KeyValues) SetValueFloat(v float64) {
	if kv == nil {
		panic("SetValueFloat on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueFloat on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprint(v)
}

// Sets the value of this node. If this node is complex or nonexistent, this
// method will panic. The value of the node will be "1" if v is true and "0"
// otherwise.
func (kv *KeyValues) SetValueBool(v bool) {
	if kv == nil {
		panic("SetValueBool on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueBool on a complex *KeyValues")
	}

	if v {
		kv.simpleValue = "1"
	} else {
		kv.simpleValue = "0"
	}
}

// Returns the first subkey (if any) that has a name equal to the argument under
// Unicode case-folding.
func (kv *KeyValues) SubKey(name string) *KeyValues {
	if kv == nil || kv.complexValue == nil {
		return nil
	}
	for i := range kv.complexValue {
		if strings.EqualFold(kv.complexValue[i].name, name) {
			return &kv.complexValue[i]
		}
	}
	return nil
}

// Creates, appends, and returns a new subkey. If the current node (the subkey's
// parent) is nil, this method will panic.
func (kv *KeyValues) NewSubKey(name string) *KeyValues {
	if kv == nil {
		panic("Call to NewSubKey on a nil *KeyValues")
	}

	kv.complexValue = append(kv.complexValue, KeyValues{name: name})
	return &kv.complexValue[len(kv.complexValue)-1]
}

// Appends the given child node to this node. If the current node (the subkey's
// parent) is nil, this method will panic. This method is a no-op on a nil child
// if the parent is valid.
func (kv *KeyValues) Append(child *KeyValues) {
	if kv == nil {
		panic("Call to Append on a nil *KeyValues")
	}

	if child == nil {
		return
	}

	kv.complexValue = append(kv.complexValue, *child)
}

// Returns a readable channel that will recieve each subkey of this node by
// reference. The channel is closed after the last node is sent. The behavior
// of this method if this node is modified while the channel is open.
//
// Example:
//      for subkey := range node.Each() {
func (kv *KeyValues) Each() <-chan *KeyValues {
	if kv == nil {
		panic("Call to Each on a nil *KeyValues")
	}

	ch := make(chan *KeyValues)
	if kv.complexValue == nil {
		// Skip spawning an extra goroutine and just close the channel.
		close(ch)
		return ch
	}
	go func() {
		for i := range kv.complexValue {
			ch <- &kv.complexValue[i]
		}
		close(ch)
	}()
	return ch
}
