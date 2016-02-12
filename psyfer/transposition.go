package psyfer

import (
	"fmt"
	"math/rand"
)

func TransposeRandom(input string) string {
	shuffle := ""
	list := rand.Perm(len(input))
	fmt.Println(list)
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
		for i, _ := range first {
			split += string(first[i]) + string(second[i])
		}
	} else {
		for i, _ := range first {
			split += string(first[i]) + string(second[i])
		}
		split += string(second[len(second)-1])
	}
	return split
}
