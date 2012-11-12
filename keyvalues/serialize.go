package keyvalues

import (
	"fmt"
	"io"
	"strings"
)

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
			c, err = kv.complexValue[i].writeIndented(w, indent+1)
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
