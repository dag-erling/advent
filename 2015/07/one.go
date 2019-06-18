package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Components have a value which can be calculated
type Component interface {
	Calculate() uint16
	Text() string
}

// Gates are a subtype of components
type Gate interface {
	Calculate() uint16
	Text() string
}

// Wires have an input, an output and a value
type Wire struct {
	label string
	value uint16
	busy, done bool
	in, out *Gate
}

func (w *Wire) Calculate() uint16 {
	if w.busy {
		panic("recursion in " + w.label)
	}
	if ! w.done && w.in != nil {
		w.busy = true
		// fmt.Printf("calculating for %s: %s\n", w.label, (*w.in).Text())
		w.value = (*w.in).Calculate()
		w.done = true
		w.busy = false
	}
	return w.value
}

func (w Wire) Text() string {
	if w.in == nil {
		return ""
	}
	return fmt.Sprintf("%s -> %s", (*w.in).Text(), w.label)
}

// Identity gate
type IdentGate struct {
	in *Wire
	out *Wire
}

func (g IdentGate) Calculate() uint16 {
	return g.in.Calculate()
}

func (g IdentGate) Text() string {
	return g.in.label
}

// Bitwise NOT gate
type NotGate struct {
	in *Wire
	out *Wire
}

func (g NotGate) Calculate() uint16 {
	return ^g.in.Calculate()
}

func (g NotGate) Text() string {
	return fmt.Sprintf("NOT %s", g.in.label)
}

// Bitwise AND gate
type AndGate struct {
	in1, in2 *Wire
	out *Wire
}

func (g AndGate) Calculate() uint16 {
	return g.in1.Calculate() & g.in2.Calculate()
}

func (g AndGate) Text() string {
	return fmt.Sprintf("%s AND %s", g.in1.label, g.in2.label)
}

// Bitwise OR gate
type OrGate struct {
	in1, in2 *Wire
	out *Wire
}

func (g OrGate) Calculate() uint16 {
	return g.in1.Calculate() | g.in2.Calculate()
}

func (g OrGate) Text() string {
	return fmt.Sprintf("%s OR %s", g.in1.label, g.in2.label)
}

// Left-shift gate
type LeftShiftGate struct {
	in *Wire
	bits uint16
	out *Wire
}

func (g LeftShiftGate) Calculate() uint16 {
	return g.in.Calculate() << g.bits
}

func (g LeftShiftGate) Text() string {
	return fmt.Sprintf("%s LSHIFT %d", g.in.label, g.bits)
}

// Right-shift gate
type RightShiftGate struct {
	in *Wire
	bits uint16
	out *Wire
}

func (g RightShiftGate) Calculate() uint16 {
	return g.in.Calculate() >> g.bits
}

func (g RightShiftGate) Text() string {
	return fmt.Sprintf("%s RSHIFT %d", g.in.label, g.bits)
}

// Map of labeled wires
var wires = map[string]*Wire{}

// Return a pointer to the wire with the given label, creating it if
// necessary.  If the label is an integer, assign that value to the
// wire.
func GetWire(label string) *Wire {
	if wires[label] == nil {
		wires[label] = &Wire{ label: label }
		if value, err := strconv.ParseUint(label, 10, 16); err == nil {
			wires[label].value = uint16(value)
			wires[label].done = true
		}
	}
	return wires[label]
}

func main() {
	lno := 0
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		lno++
		// fmt.Printf("%4d %s\n", lno, line)
		sides := strings.Split(line, " -> ")
		if len(sides) != 2 {
			panic("syntax error")
		}
		lhs := strings.Split(sides[0], " ")
		rhs := sides[1]
		wout := GetWire(rhs)
		var gate Gate
		switch {
		case len(lhs) == 1:
			// The text of the assignment doesn't mention
			// that the lhs here can be not only an
			// integer literal but also a wire.  We
			// generalize this to a gate which copies its
			// input to its output.
			win := GetWire(lhs[0])
			gate = IdentGate{ win, wout }
			win.out = &gate
			wout.in = &gate
		case len(lhs) == 2 && lhs[0] == "NOT":
			win := GetWire(lhs[1])
			gate = NotGate{ win, wout }
			win.out = &gate
			wout.in = &gate
		case len(lhs) == 3 && lhs[1] == "AND":
			win1, win2 := GetWire(lhs[0]), GetWire(lhs[2])
			gate = AndGate{ win1, win2, wout }
			win1.out = &gate
			win2.out = &gate
			wout.in = &gate
		case len(lhs) == 3 && lhs[1] == "OR":
			win1, win2 := GetWire(lhs[0]), GetWire(lhs[2])
			gate = OrGate{ win1, win2, wout }
			win1.out = &gate
			win2.out = &gate
			wout.in = &gate
		case len(lhs) == 3 && lhs[1] == "LSHIFT":
			win := GetWire(lhs[0])
			bits, _ := strconv.ParseUint(lhs[2], 10, 16)
			gate = LeftShiftGate{ win, uint16(bits), wout }
			win.out = &gate
			wout.in = &gate
		case len(lhs) == 3 && lhs[1] == "RSHIFT":
			win := GetWire(lhs[0])
			bits, _ := strconv.ParseUint(lhs[2], 10, 16)
			gate = RightShiftGate{ win, uint16(bits), wout }
			win.out = &gate
			wout.in = &gate
		default:
			panic("syntax error");
		}
	}
//	fmt.Println(len(wires), "wires")
//	for _, w := range wires {
//		if w.in != nil { // XXX
//			fmt.Println(w.Text())
//		}
//	}
	a := GetWire("a")
	a.Calculate()
	fmt.Println(a.value)
}
