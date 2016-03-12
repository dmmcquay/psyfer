package psyfer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type KeyJSON struct {
	Key map[string]string `json:"key"`
}

var k = KeyJSON{}

func ReadConfig() {
	dat, err := ioutil.ReadFile("key.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(dat), &k)
	fmt.Println(k.Key["a"])
}

func Substitution(input string) string {
	inputSlice := []string{}
	output := ""
	for _, char := range input {
		inputSlice = append(inputSlice, string(char))
	}
	for i := range inputSlice {
		output += k.Key[inputSlice[i]]
	}
	fmt.Println(output)
	return output
}
