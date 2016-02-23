package psyfer

import "math/rand"

func TransposeRandom(input string) string {
	shuffle := ""
	list := rand.Perm(len(input))
	for _, i := range list {
		shuffle += string(input[i])
	}
	return shuffle
}

func TransposeRailFence(input string) string {
	rf := ""
	for i := 0; i < len(input); i += 2 {
		rf += string(input[i])
	}
	for i := 1; i < len(input); i += 2 {
		rf += string(input[i])
	}
	return rf
}

func TransposeSplit(input string) string {
	split := ""
	length := len(input)
	first := input[:length/2]
	second := input[length/2:]
	if length%2 == 0 {
		for i := range first {
			split += string(first[i]) + string(second[i])
		}
	} else {
		for i := range first {
			split += string(first[i]) + string(second[i])
		}
		split += string(second[len(second)-1])
	}
	return split
}

func DeTransposeRailFence(input string) string {
	derf := ""
	length := len(input)
	first := input[:length/2]
	second := input[length/2:]
	if length%2 == 0 {
		for i := range first {
			derf += string(first[i]) + string(second[i])
		}
	} else {
		for i := range first {
			derf += string(first[i]) + string(second[i+1])
		}
		derf += string(second[0])
	}
	return derf
}

func DeTransposeSplit(input string) string {
	desplit := ""
	if len(input)%2 == 0 {
		for i := 0; i < len(input); i += 2 {
			desplit += string(input[i])
		}
		for i := 1; i < len(input); i += 2 {
			desplit += string(input[i])
		}
	} else {
		for i := 0; i < len(input)-2; i += 2 {
			desplit += string(input[i])
		}
		for i := 1; i < len(input)-2; i += 2 {
			desplit += string(input[i])
		}
		desplit += string(input[len(input)-2]) + string(input[len(input)-1])
	}
	return desplit
}
