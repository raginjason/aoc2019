package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/bits"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Segment struct {
	horizontal bool
	vertical   bool
	begin      Point
	end        Point
}

func (s Segment) intersect(b Segment) Point {
	var p Point

	var verticalSeg Segment
	var horizontalSeg Segment

	if s.vertical && b.horizontal {
		verticalSeg = s
		horizontalSeg = b
	} else if s.horizontal && b.vertical {
		verticalSeg = b
		horizontalSeg = s
	} else {
		return p
	}

	maxX := Max(horizontalSeg.begin.x, horizontalSeg.end.x)
	minX := Min(horizontalSeg.begin.x, horizontalSeg.end.x)
	maxY := Max(verticalSeg.begin.y, verticalSeg.end.y)
	minY := Min(verticalSeg.begin.y, verticalSeg.end.y)
	if minX <= verticalSeg.begin.x && verticalSeg.begin.x <= maxX && minY <= horizontalSeg.begin.y && horizontalSeg.begin.y <= maxY {
		p.x = verticalSeg.begin.x
		p.y = horizontalSeg.begin.y
	}
	return p
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s Segment) String() string {
	return fmt.Sprintf("%s-%s", s.begin, s.end)
}

func CrossedWires(inputA []string, inputB []string) int {

	aSegments := ComputeSegments(inputA)
	bSegments := ComputeSegments(inputB)
	zeroPoint := Point{0, 0}
	intersections := make([]Point, 0)

	for _, aSeg := range aSegments {
		for _, bSeg := range bSegments {
			point := aSeg.intersect(bSeg)
			if point != zeroPoint {
				intersections = append(intersections, point)
			}
		}
	}

	closestDistance := 1<<(bits.UintSize-1) - 1 // Set to something big
	for _, intersection := range intersections {
		distance := Abs(intersection.x) + Abs(intersection.y)
		if distance < closestDistance {
			closestDistance = distance
		}
	}

	return closestDistance
}

func CalculateIntersection(a, b Segment) Point {
	if a.vertical && b.vertical { // Segments are parallel
		return Point{0, 0}
	}

	if a.horizontal && b.horizontal { // Segments are parallel
		return Point{0, 0}
	}

	return a.intersect(b)
}

func ComputeSegments(input []string) []Segment {
	segs := make([]Segment, len(input))

	seg := Segment{begin: Point{0, 0}}
	for i, path := range input {

		dir := string(path[0])
		distance, err := strconv.Atoi(string(path[1:]))
		if err != nil {
			log.Fatalf("could not convert %s, %s\n", string(path[1:]), err)
		}

		switch dir {
		case "U": // +y
			seg.end.y = seg.begin.y + distance
			seg.end.x = seg.begin.x
			seg.vertical = true
		case "D": // -y
			seg.end.y = seg.begin.y - distance
			seg.end.x = seg.begin.x
			seg.vertical = true
		case "L": // -x
			seg.end.x = seg.begin.x - distance
			seg.end.y = seg.begin.y
			seg.horizontal = true
		case "R": // +x
			seg.end.x = seg.begin.x + distance
			seg.end.y = seg.begin.y
			seg.horizontal = true
		}

		segs[i] = seg

		// Reset for next iteration after adding to segs
		seg.begin.x = seg.end.x
		seg.begin.y = seg.end.y
		seg.vertical = false
		seg.horizontal = false
	}
	return segs
}

func scanDay3File() [][]string {
	lines, err := FileInput(3)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	var paths [][]string

	for _, line := range lines {
		r := csv.NewReader(strings.NewReader(line))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			paths = append(paths, record)
		}
	}
	return paths
}

func day3() int {
	paths := scanDay3File()

	res := CrossedWires(paths[0], paths[1])

	return res
}
