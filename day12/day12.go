package day12

import (
	"encoding/csv"
	"fmt"
	"github.com/raginjason/aoc2019/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y, Z int
}

type Velocity struct {
	X, Y, Z int
}

type Moon struct {
	Pos Position
	Vel Velocity
}

func CalculateVelocityDelta(a, b int) (int, int) { // Test works
	bRes, aRes := 0, 0
	if a < b {
		aRes++
		bRes--
	} else if a > b {
		aRes--
		bRes++
	}
	return aRes, bRes
}

func CalculateGravity(a, b *Moon, av, bv *Velocity) {
	axRes, bxRes := CalculateVelocityDelta(a.Pos.X, b.Pos.X)
	ayRes, byRes := CalculateVelocityDelta(a.Pos.Y, b.Pos.Y)
	azRes, bzRes := CalculateVelocityDelta(a.Pos.Z, b.Pos.Z)

	av.X = av.X + axRes
	av.Y = av.Y + ayRes
	av.Z = av.Z + azRes
	bv.X = bv.X + bxRes
	bv.Y = bv.Y + byRes
	bv.Z = bv.Z + bzRes
}

// Applying gravity mutates velocity
func (a *Moon) ApplyGravity(v *Velocity) {
	a.Vel.X = a.Vel.X + v.X
	a.Vel.Y = a.Vel.Y + v.Y
	a.Vel.Z = a.Vel.Z + v.Z
}

// Applying velocity mutates position
func (m *Moon) ApplyVelocity() { // Test works
	m.Pos.X = m.Pos.X + m.Vel.X
	m.Pos.Y = m.Pos.Y + m.Vel.Y
	m.Pos.Z = m.Pos.Z + m.Vel.Z
}

func (m *Moon) PotentialEnergy() int { // Test works
	return Abs(m.Pos.X) + Abs(m.Pos.Y) + Abs(m.Pos.Z)
}

func (m *Moon) KineticEnergy() int { // Test works
	return Abs(m.Vel.X) + Abs(m.Vel.Y) + Abs(m.Vel.Z)
}

func (m *Moon) TotalEnergy() int { // Test works
	return m.PotentialEnergy() * m.KineticEnergy()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m Moon) String() string {
	//pos=<x=-1, y=  0, z= 2>, vel=<x= 0, y= 0, z= 0>
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.Pos.X, m.Pos.Y, m.Pos.Z, m.Vel.X, m.Vel.Y, m.Vel.Z)
}

func scanFile() []Moon {
	fh, err := os.Open(util.InputFilePath(12))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	moons := parseMoons(fh)

	return moons
}

func parseMoons(reader io.Reader) []Moon {
	var moons []Moon

	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read from file: %s", err)
		}

		// Clean up record some
		record[0] = strings.TrimPrefix(record[0], "<")
		record[2] = strings.TrimSuffix(record[2], ">")

		var location [3]int
		for i := range record {
			parts := strings.Split(record[i], "=")
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("could not convert %s, %s\n", parts[1], err)
			}
			location[i] = num
		}
		moon := Moon{
			Position{location[0], location[1], location[2]},
			Velocity{0, 0, 0},
		}
		moons = append(moons, moon)
	}

	return moons
}

func Simulate(moons []Moon, maxSteps int) {
	var step int
	for {
		if step == maxSteps {
			break
		}
		step++
		var delta [4]Velocity
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				CalculateGravity(&moons[i], &moons[j], &delta[i], &delta[j])
			}
		}
		for i := range moons {
			moons[i].ApplyGravity(&delta[i])
		}
		for i := range moons {
			moons[i].ApplyVelocity()
		}
	}
}

func Part1() int {
	moons := scanFile()

	Simulate(moons, 1000)

	var totalEnergy int
	for _, moon := range moons {
		totalEnergy = totalEnergy + moon.TotalEnergy()
	}

	return totalEnergy
}
