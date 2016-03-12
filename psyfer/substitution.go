package psyfer

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type KeyJSON struct {
	Key map[string]string `json:"key"`
}

var k = KeyJSON{}

func ReadConfig(file string) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(dat), &k)
	k.Key[" "] = " " // keep spaces alive
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
	return output
}
