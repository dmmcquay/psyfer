package main

import (
	"fmt"
	"math/rand"
	"time"

	"s.mcquay.me/dm/psyfer/psyfer"

	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	psyfer.ReadConfig()
	psyfer.Substitution("hello")
	var input string
	var cipher string
	var cipType string
	var key string

	var encrypt = &cobra.Command{
		Use:   "encrypt -c [cipher] -k [key] -i [input]",
		Short: "encrypt string",
		Long:  `encrypt given string`,
		Run: func(cmd *cobra.Command, args []string) {
			if input != "" {
				switch {
				case cipher == "transposition":
					switch {
					case cipType == "random":
						fmt.Println(psyfer.TransposeRandom(input))
					case cipType == "rail":
						fmt.Println(psyfer.TransposeRailFence(input))
					case cipType == "split":
						fmt.Println(psyfer.TransposeSplit(input))
					default:
						fmt.Println("Missing cipher sub type (random, rail, split)")
					}
				case cipher == "substitution":
					fmt.Println("substitution")
				case cipher == "vigenere":
					fmt.Println("vigenere")
				default:
					fmt.Println("Must choose transposition, substitution, or vigenere")
				}
			} else {
				fmt.Println("Missing input")
			}
		},
	}

	var decrypt = &cobra.Command{
		Use:   "decrypt -c [cipher] -k [key] -i [input]",
		Short: "decrypt string",
		Long:  `decrypt given string`,
		Run: func(cmd *cobra.Command, args []string) {
			if input != "" {
				switch {
				case cipher == "transposition":
					switch {
					case cipType == "random":
						fmt.Println("random doesn't have a decryptable solution")
					case cipType == "rail":
						fmt.Println(psyfer.DeTransposeRailFence(input))
					case cipType == "split":
						fmt.Println(psyfer.DeTransposeSplit(input))
					default:
						fmt.Println("Missing cipher sub type (random, rail, split)")
					}
				case cipher == "substitution":
					fmt.Println("substitution")
				case cipher == "vigenere":
					fmt.Println("vigenere")
				default:
					fmt.Println("Must choose transposition, substitution, or vigenere")
				}
			} else {
				fmt.Println("Missing input")
			}
		},
	}

	var crack = &cobra.Command{
		Use:   "times [# times] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long:  `echo things multiple times back to the user by providing a count and a string.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("crack")
		},
	}

	encrypt.Flags().StringVarP(
		&input,
		"input",
		"i",
		"",
		"string to be encrypted",
	)
	encrypt.Flags().StringVarP(
		&cipher,
		"cipher",
		"c",
		"",
		"cipher to be used (transposition, substitution, vigenere)",
	)
	encrypt.Flags().StringVarP(
		&key,
		"key",
		"k",
		"",
		"key to be used",
	)
	encrypt.Flags().StringVarP(
		&cipType,
		"type",
		"t",
		"",
		"sub type of cipher",
	)

	decrypt.Flags().StringVarP(
		&input,
		"input",
		"i",
		"",
		"string to be encrypted",
	)
	decrypt.Flags().StringVarP(
		&cipher,
		"cipher",
		"c",
		"",
		"cipher to be used (transposition, substitution, vigenere)",
	)
	decrypt.Flags().StringVarP(
		&key,
		"key",
		"k",
		"",
		"key to be used",
	)
	decrypt.Flags().StringVarP(
		&cipType,
		"type",
		"t",
		"",
		"sub type of cipher",
	)

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(encrypt, decrypt, crack)
	rootCmd.Execute()
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("helloworld")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("1")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("12")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("123")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("1234")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("12345")))
	//fmt.Println(psyfer.DeTransposeRailFence(psyfer.TransposeRailFence("1234567")))
}
