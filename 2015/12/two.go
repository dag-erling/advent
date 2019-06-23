package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// As far as I can tell, Go's built-in JSON implementation requires
// prior knowledge of the structure of the data being read, so we have
// to implement our own.

type Input struct {
	r *bufio.Reader
	pos int
}

func NewInput(ior io.Reader) *Input {
	in := Input{ r: bufio.NewReader(ior) }
	return &in
}

func (in *Input) panic(msg string) {
	panic(fmt.Sprintf("%s at position %d", msg, in.pos))
}

func (in *Input) getb(bp *byte) bool {
	b, err := in.r.ReadByte()
	if err != nil {
		return false
	}
	*bp = b
	in.pos++
	return true
}

func (in *Input) ungetb() bool {
	err := in.r.UnreadByte()
	in.pos--
	return err == nil
}

func (in *Input) gets(sp *string, delim byte) bool {
	s, err := in.r.ReadString(delim)
	if err != nil {
		return false
	}
	*sp = s
	in.pos += len(s)
	return true
}

func (in *Input) gettok(sp *string) bool {
	var b byte
	var s string
	for in.getb(&b) {
		switch {
		case b == '"':
			*sp = string(b)
			if ! in.gets(&s, '"') {
				in.panic("Unterminated string")
			}
			*sp += s
			return true
		case b == '-' || (b >= '0' && b <= '9'):
			// XXX should reject a naked minus sign
			*sp = string(b)
			for in.getb(&b) {
				if b < '0' || b > '9' {
					in.ungetb()
					break
				}
				*sp += string(b)
			}
			return true
		case b == '[' || b == ']' || b == '{' || b == '}' || b == ':' || b == ',':
			*sp = string(b)
			return true
		case b == ' ' || b == '\t' || b == '\r' || b == '\n':
			continue
		default:
			in.panic(fmt.Sprintf("Invalid character %02x", b))
		}
	}
	return false
}

func sum_object(in *Input) int {
	var tok string
	var red bool
	var sum int
	for {
		if ! in.gettok(&tok) || tok[0] != '"' {
			in.panic("Expected identifier")
		}
		if ! in.gettok(&tok) || tok != ":" {
			in.panic("Expected colon")
		}
		if ! in.gettok(&tok) {
			in.panic("Expected value")
		}
		switch {
		case tok == "[":
			sum += sum_array(in)
		case tok == "{":
			sum += sum_object(in)
		case tok == "\"red\"":
			red = true
		case tok[0] == '"':
			// some string, we don't care
		case tok[0] == '-' || (tok[0] >= '0' && tok[0] <= '9'):
			n, _ := strconv.Atoi(tok)
			sum += n
		default:
			in.panic(fmt.Sprintf("Unexpected token '%s'", tok))
		}
		if ! in.gettok(&tok) || (tok != "," && tok != "}") {
			in.panic("Expected comma or right brace")
		}
		if tok == "}" {
			break
		}
	}
	if red {
		sum = 0
	}
	return sum
}

func sum_array(in *Input) int {
	var tok string
	var sum int
	for {
		if ! in.gettok(&tok) {
			in.panic("Expected value")
		}
		switch {
		case tok == "[":
			sum += sum_array(in)
		case tok == "{":
			sum += sum_object(in)
		case tok[0] == '"':
			// some string, we don't care
		case tok[0] == '-' || (tok[0] >= '0' && tok[0] <= '9'):
			n, _ := strconv.Atoi(tok)
			sum += n
		default:
			in.panic(fmt.Sprintf("Unexpected token '%s'", tok))
		}
		if ! in.gettok(&tok) || (tok != "," && tok != "]") {
			in.panic("Expected comma or right bracket")
		}
		if tok == "]" {
			break
		}
	}
	return sum
}

func sum_json(in *Input) int {
	var tok string
	var sum int
	if in.gettok(&tok) {
		switch {
		case tok == "[":
			sum += sum_array(in)
		case tok == "{":
			sum += sum_object(in)
		case tok[0] == '"':
			// some string, we don't care
		case tok[0] == '-' || (tok[0] >= '0' && tok[0] <= '9'):
			n, _ := strconv.Atoi(tok)
			sum += n
		default:
			in.panic(fmt.Sprintf("Unexpected token '%s'", tok))
		}
	}
	return sum
}

func main() {
	in := NewInput(os.Stdin)
	sum := sum_json(in)
	fmt.Println(sum)
}
