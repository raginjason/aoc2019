package day7

import (
	"github.com/raginjason/aoc2019/intcode"
	"github.com/raginjason/aoc2019/util"
	"log"
	"os"
)

func Amplifier(origProg intcode.Program, signalSeq []int) int {

	// Number of amplifiers (5) + 1
	channels := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	var amplifiers []*intcode.Computer
	for i := range signalSeq {
		p := make(intcode.Program, len(origProg))
		copy(p, origProg)
		c := intcode.NewComputer(channels[i], p, channels[i+1])
		c.Name = string(65 + i) // 65 = A, 66 = B, etc
		amplifiers = append(amplifiers, c)
	}

	// Fire up amplifiers
	for i := len(amplifiers) - 1; i >= 0; i-- {
		go amplifiers[i].Run()
	}

	// Set signal sequence on each amplifier
	for i, phase := range signalSeq {
		channels[i] <- phase
	}

	firstChan := channels[0]
	lastChan := channels[len(channels)-1]

	// Seed first channel with signal of 0
	firstChan <- 0

	var output []int
	for {
		val, ok := <-lastChan
		if !ok {
			break
		} else {
			output = append(output, val)
		}
	}

	return output[0]
}

func FeedbackAmplifier(origProg intcode.Program, signalSeq []int) int {

	// Number of amplifiers (5) + 1
	channels := []chan int{
		make(chan int, 1), // main <-> A
		make(chan int, 1), // A <-> B
		make(chan int, 1), // B <-> C
		make(chan int, 1), // C <-> D
		make(chan int, 1), // D <-> E
		make(chan int, 1), // E <-> main
	}

	var amplifiers []*intcode.Computer
	for i := range signalSeq {
		p := make(intcode.Program, len(origProg))
		copy(p, origProg)
		c := intcode.NewComputer(channels[i], p, channels[i+1])
		c.Name = string(65 + i) // 65 = A, 66 = B, etc
		amplifiers = append(amplifiers, c)
	}

	// Fire up amplifiers
	for i := len(amplifiers) - 1; i >= 0; i-- {
		go amplifiers[i].Run()
	}

	// Set signal sequence on each amplifier
	for i, phase := range signalSeq {
		channels[i] <- phase
	}

	firstChan := channels[0]
	lastChan := channels[len(channels)-1]

	// Seed first channel with signal of 0
	firstChan <- 0

	var lastSignal int
	for {
		if signal, ok := <-lastChan; ok {
			lastSignal = signal
			firstChan <- signal
		} else { // This should handle opcode 99
			break
		}
	}

	return lastSignal
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func MaxSignal(p intcode.Program) int {

	signals := permutation([]int{0, 1, 2, 3, 4})

	var maxSignal int
	for _, s := range signals {
		signal := Amplifier(p, s)
		if signal > maxSignal {
			maxSignal = signal
		}
	}

	return maxSignal
}

func MaxFeedbackSignal(p intcode.Program) int {

	signals := permutation([]int{5, 6, 7, 8, 9})

	var maxSignal int
	for _, s := range signals {
		signal := FeedbackAmplifier(p, s)
		if signal > maxSignal {
			maxSignal = signal
		}
	}

	return maxSignal
}

func scanDay7File() []int {
	fh, err := os.Open(util.InputFilePath(7))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	program := intcode.ParseIntCode(fh)

	return program
}

func Part1() int {
	program := scanDay7File()

	maxSignal := MaxSignal(program)

	return maxSignal
}

func Part2() int {
	program := scanDay7File()

	maxFeedbackSignal := MaxFeedbackSignal(program)

	return maxFeedbackSignal
}
