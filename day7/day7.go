package day7

import "github.com/raginjason/aoc2019/intcode"

func Amplifier(origProg intcode.Program, signalSeq []int) int {

	var signal int

	for _, phase := range signalSeq {
		p := make(intcode.Program, len(origProg))
		copy(p, origProg)

		c := intcode.NewComputer([]int{phase, signal}, p)
		o := c.Run()
		signal = o[0]
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
