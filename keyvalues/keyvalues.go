package keyvalues

import (
	"fmt"
	"io"
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

func (kv *KeyValues) String(def string) string {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	return kv.simpleValue
}

func (kv *KeyValues) Int(def int64) int64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if i, err := strconv.ParseInt(kv.simpleValue, 0, 64); err == nil {
		return i
	}
	return def
}

func (kv *KeyValues) Uint64(def uint64) uint64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if i, err := strconv.ParseUint(kv.simpleValue, 0, 64); err == nil {
		return i
	}
	return def
}

func (kv *KeyValues) Float(def float64) float64 {
	if kv == nil || kv.complexValue != nil {
		return def
	}
	if f, err := strconv.ParseFloat(kv.simpleValue, 64); err == nil {
		return f
	}
	return def
}

func (kv *KeyValues) Bool(def bool) bool {
	if def {
		return kv.Int(1) != 0
	}
	return kv.Int(0) != 0
}

func (kv *KeyValues) SetValueString(v string) {
	if kv == nil {
		panic("SetValueString on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueString on a complex *KeyValues")
	}

	kv.simpleValue = v
}

func (kv *KeyValues) SetValueInt(v int64) {
	if kv == nil {
		panic("SetValueInt on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueInt on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprint(v)
}

func (kv *KeyValues) SetValueUint64(v uint64) {
	if kv == nil {
		panic("SetValueUint64 on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueUint64 on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprintf("0x%x", v)
}

func (kv *KeyValues) SetValueFloat(v float64) {
	if kv == nil {
		panic("SetValueFloat on a nil *KeyValues")
	}
	if kv.complexValue != nil {
		panic("SetValueFloat on a complex *KeyValues")
	}

	kv.simpleValue = fmt.Sprint(v)
}

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

func (kv *KeyValues) NewSubKey(name string) *KeyValues {
	if kv == nil {
		panic("Call to NewSubKey on a nil *KeyValues")
	}

	kv.complexValue = append(kv.complexValue, KeyValues{name: name})
	return &kv.complexValue[len(kv.complexValue)-1]
}

func (kv *KeyValues) WriteTo(w io.Writer) (n int, err error) {
	for i := range kv.complexValue {
		var c int
		c, err = kv.complexValue[i].writeIndented(w, 0)
		n += c
		if err != nil {
			return
		}
	}
	return
}

func (kv *KeyValues) writeIndented(w io.Writer, indent int) (n int, err error) {
	indentString := strings.Repeat("\t", indent)
	n, err = fmt.Fprintf(w, "%s%q ", indentString, kv.name)
	if err != nil {
		return
	}

	if kv.complexValue == nil {
		var c int

		c, err = fmt.Fprintf(w, "%q\n", kv.simpleValue)
		n += c
		if err != nil {
			return
		}
	} else {
		var c int

		c, err = fmt.Fprintf(w, "{\n")
		n += c
		if err != nil {
			return
		}

		for i := range kv.complexValue {
			c, err = kv.complexValue[i].writeIndented(w, indent + 1)
			n += c
			if err != nil {
				return
			}
		}

		c, err = fmt.Fprintf(w, "%s}\n", indentString)
		n += c
		if err != nil {
			return
		}
	}
	return
}