package keyvalues

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

func consumeSpaces(in *bufio.Reader) (n int64, err error) {
	r := ' '
	var c int
	for unicode.IsSpace(r) {
		n += int64(c)
		r, c, err = in.ReadRune()
		if err != nil {
			return
		}
	}
	err = in.UnreadRune()
	return
}

func readString(in *bufio.Reader) (s string, n int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
				return
			}
			panic(r)
		}
	}()
	var (
		buf    []rune
		quoted bool

		rr = func() rune {
			r, c, err := in.ReadRune()
			n += int64(c)
			if err != nil {
				panic(err)
			}
			return r
		}
	)

	r := rr()

	if r == '"' {
		quoted = true
	} else {
		buf = append(buf, r)
	}

	for {
		r = rr()
		if r == '"' && quoted {
			s = string(buf)
			return
		}
		if unicode.IsSpace(r) && !quoted {
			s = string(buf)
			return
		}
		if (r == '"' || r == '{' || r == '}') && !quoted {
			err = in.UnreadRune()
			s = string(buf)
			return
		}
		if r == '\\' && quoted {
			next := rr()
			switch next {
			case 'n':
				r = '\n'
			case 'r':
				r = '\r'
			case 't':
				r = '\t'
			case '"':
				r = '"'
			default:
				buf = append(buf, r)
				r = next
			}
		}
		buf = append(buf, r)
	}
	panic("unreachable")
}

type stack []*KeyValues

func (s *stack) push(kv *KeyValues) {
	*s = append(*s, kv)
}

func (s stack) peek() *KeyValues {
	if len(s) == 0 {
		panic("stack underflow")
	}
	return s[len(s)-1]
}

func (s *stack) pop() *KeyValues {
	if len(*s) == 0 {
		panic("stack underflow")
	}
	kv := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return kv
}

func (kv *KeyValues) ReadFrom(r io.Reader) (n int64, err error) {
	in := bufio.NewReader(r)

	var s = stack{kv}

	for err == nil {
		var c int64

		c, err = consumeSpaces(in)
		n += c
		if err != nil {
			if err == io.EOF && s.pop() == kv {
				err = nil
			}
			return
		}

		var r rune
		var c_ int
		r, c_, err = in.ReadRune()
		n += int64(c_)
		if err != nil {
			if err == io.EOF && s.pop() == kv {
				err = nil
			}
			return
		}
		switch r {
		case '{':
			err = errors.New("Unexpected '{': expecting '}' or a key")
			return
		case '}':
			s.pop()
			if len(s) == 0 {
				err = errors.New("Unexpected '}': expecting a key")
				return
			}
			continue
		default:
			err = in.UnreadRune()
			if err != nil {
				return
			}
			n -= int64(c_)
		}

		var key string
		key, c, err = readString(in)
		n += c
		if err != nil {
			return
		}
		s.push(s.peek().NewSubKey(key))

		c, err = consumeSpaces(in)
		n += c
		if err != nil {
			return
		}

		r, c_, err = in.ReadRune()
		n += int64(c_)
		if err != nil {
			return
		}
		switch r {
		case '{':
			continue
		case '}':
			err = errors.New("Unexpected '}': expecting '{' or a value")
			return
		default:
			err = in.UnreadRune()
			if err != nil {
				return
			}
			n -= int64(c_)
		}

		var value string
		value, c, err = readString(in)
		n += c
		if err != nil {
			return
		}
		s.pop().SetValueString(value)
	}

	return
}

var _ io.ReaderFrom = new(KeyValues)
