package psyfer

import (
	"fmt"
	"math/rand"
)

func Transpose(input string) string {
	shuffle := ""
	list := rand.Perm(len(input))
	fmt.Println(list)
	for _, i := range list {
		shuffle += string(input[i])
	}
	return shuffle
}
