package main

func GravityComputer(opCodes []int) []int {
	skip := 0
	for i, op := range opCodes {
		if skip > 1 {
			skip = skip - 1
			continue
		}

		switch op {
		case 1: // Add
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] + opCodes[opCodes[i+2]]
			skip = 4
		case 2: // Multiply
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] * opCodes[opCodes[i+2]]
			skip = 4
		case 99: // Terminate
			return opCodes
		default:
			continue
		}
	}
	return opCodes
}
