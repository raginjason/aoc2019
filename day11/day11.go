package day11

import (
	"github.com/raginjason/aoc2019/intcode"
	"github.com/raginjason/aoc2019/util"
	"log"
	"os"
	"strconv"
)

func scanDay11File() []int {
	fh, err := os.Open(util.InputFilePath(11))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	program := intcode.ParseIntCode(fh)

	return program
}

func Part1() int {
	program := scanDay11File()

	in := make(chan int, 10)
	out := make(chan int, 10)
	c := intcode.NewComputer(in, program, out)
	go c.Run()

	h := NewHull()
	r := Robot{Up, Coordinate{0, 0}}

	for {
		in <- int(r.ScanHull(h))

		// 0 = black, 1 = white
		val, ok := <-out
		if !ok {
			break
		}
		newColor := Color(val)
		r.PaintPanel(h, newColor)

		// 0 = left, 1 = right
		val, ok = <-out
		if !ok {
			break
		}
		direction := Turn(val)
		r.Turn(direction)

		r.Move()
	}
	return len(h.PaintCount)
}

type Robot struct {
	Orientation Direction
	Location    Coordinate
}

func (r *Robot) Turn(t Turn) {
	switch t {
	case TurnRight:
		r.Orientation++
	case TurnLeft:
		r.Orientation--
	}

	if r.Orientation > Left {
		r.Orientation = Up
	}

	if r.Orientation < Up {
		r.Orientation = Left
	}
}

func (r *Robot) Move() {
	switch r.Orientation {
	case Up:
		r.Location.Y++
	case Right:
		r.Location.X++
	case Down:
		r.Location.Y--
	case Left:
		r.Location.X--
	}
}

func (r *Robot) ScanHull(h *Hull) Color {
	return h.Panels[r.Location]
}

func (r *Robot) PaintPanel(h *Hull, c Color) {
	h.Paint(r.Location, c)
}

// Direction
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}

// Hull
type Hull struct {
	Panels     map[Coordinate]Color
	PaintCount map[Coordinate]int
}

func NewHull() *Hull {
	h := new(Hull)
	h.Panels = make(map[Coordinate]Color)
	h.PaintCount = make(map[Coordinate]int)
	return h
}

func (h *Hull) Color(c Coordinate) Color {
	if val, ok := h.Panels[c]; ok {
		return val
	}
	return Black
}

func (s *Hull) Paint(c Coordinate, color Color) {
	s.PaintCount[c]++
	s.Panels[c] = color
}

// Coordinate
type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

// Color
type Color int

const (
	Black Color = iota
	White
)

func (c Color) String() string {
	return [...]string{"Black", "White"}[c]
}

// Turn
type Turn int

const (
	TurnLeft Turn = iota
	TurnRight
)

func (d Turn) String() string {
	return [...]string{"Turn Left", "Turn Right"}[d]
}
