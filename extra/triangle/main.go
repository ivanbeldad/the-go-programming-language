package main

import (
	"fmt"
	"sort"
	"strings"
)

type Tree struct {
	Units map[int]map[int]*Unit
}

func (t Tree) unitVal(pos Position) int {
	if t.Units[pos.y] == nil {
		return 0
	}
	if t.Units[pos.y][pos.x] == nil {
		return 0
	}
	return t.Units[pos.y][pos.x].value
}

func (t Tree) sortYKeys() []int {
	var ykeys []int
	for k := range t.Units {
		ykeys = append(ykeys, k)
	}
	sort.Ints(ykeys)
	return ykeys
}

func (t Tree) sortXKeys(ykey int) []int {
	var xkeys []int
	for k := range t.Units[ykey] {
		xkeys = append(xkeys, k)
	}
	sort.Ints(xkeys)
	return xkeys
}

func (t Tree) String() string {
	s := strings.Builder{}
	ykeys := t.sortYKeys()
	for row, ykey := range ykeys {
		xkeys := t.sortXKeys(ykey)
		spaces := ""
		for i := len(ykeys) - 1; i > row; i-- {
			spaces += "   "
		}
		s.Write([]byte(spaces))
		for _, xkey := range xkeys {
			s.Write([]byte(fmt.Sprintf(" %5d", t.Units[ykey][xkey].value)))
		}
		s.Write([]byte("\n\n"))
	}
	return s.String()
}

type Position struct {
	y int
	x int
}

func (p Position) String() string {
	return fmt.Sprintf("y: %4d    x: %d", p.y, p.x)
}

type Unit struct {
	value int
	pos   Position
}

func (u *Unit) calc(t Tree) {
	if u.pos.y == 0 {
		return
	}
	u.value = t.unitVal(u.left()) + t.unitVal(u.right())
}

func (u *Unit) left() Position {
	return Position{
		x: u.pos.x - 1,
		y: u.pos.y - 1,
	}
}

func (u *Unit) right() Position {
	return Position{
		x: u.pos.x + 1,
		y: u.pos.y - 1,
	}
}

func (u Unit) String() string {
	return fmt.Sprintf("Position{y: %3d x: %3d value: %3d}", u.pos.y, u.pos.x, u.value)
}

func main() {
	t := Tree{}
	genRows(20, &t)
	fmt.Print(t)
}

func genRows(c int, t *Tree) {
	if t.Units == nil {
		t.Units = make(map[int]map[int]*Unit)
	}
	for i := 0; i < c; i++ {
		genRow(i, t)
	}
}

func genRow(r int, t *Tree) {
	var u *Unit
	t.Units[r] = make(map[int]*Unit)
	if r == 0 {
		u = &Unit{value: 1, pos: Position{x: 0, y: 0}}
		t.Units[r][0] = u
		return
	}
	pairs := (r + 1) / 2
	center := 1
	if r%2 == 0 {
		center = 0
		pairs++
	}
	for i := 0; i < pairs; i++ {
		pos := i*2 + center
		u = &Unit{pos: Position{x: pos, y: r}}
		u.calc(*t)
		t.Units[r][pos] = u
		if pos == 0 {
			continue
		}
		u = &Unit{pos: Position{x: pos * -1, y: r}}
		u.calc(*t)
		t.Units[r][pos*-1] = u
	}
}
