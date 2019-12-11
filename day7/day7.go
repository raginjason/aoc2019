package day7

import (
	"github.com/raginjason/aoc2019/intcode"
	"github.com/raginjason/aoc2019/util"
	"log"
	"os"
)

func Amplifier(origProg intcode.Program, signalSeq []int) int {

	var signal int

	for _, phase := range signalSeq {
		p := make(intcode.Program, len(origProg))
		copy(p, origProg)

		in := make(chan int)
		out := make(chan int)
		c := intcode.NewComputer(in, p, out)
		go c.Run()
		in <- phase
		in <- signal

		var output []int
		for {
			val, ok := <-out
			if ok == false {
				break
			} else {
				output = append(output, val)
			}
		}
		signal = output[0]
	}

	return signal
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
