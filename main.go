package main

import (
	"fmt"
	"math/rand"
	"time"

	"s.mcquay.me/dm/psyfer/psyfer"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(psyfer.Transpose("hello world"))
	fmt.Println(psyfer.Transpose("hello"))
}
