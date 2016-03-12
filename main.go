package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"s.mcquay.me/dm/psyfer/psyfer"

	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//psyfer.ReadConfig()
	//psyfer.Substitution("hello")
	var input string
	var cipher string
	var cipType string
	//var key string
	var decrypt bool

	var sub = &cobra.Command{
		Use:   "sub mode -c [cipher] -k [key] -i [input]",
		Short: "substitution cipher",
		Long:  `perform substitution cipher`,
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

	var trans = &cobra.Command{
		Use:   "trans mode",
		Short: "transposition cipher",
		Long:  `perform transposition cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("missing input, see -h (--help) for more info")
		},
	}

	var aes = &cobra.Command{
		Use:   "aes mode -c [cipher] -k [key] -i [input]",
		Short: "aes cipher",
		Long:  `perform aes cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("missing input, see -h (--help) for more info")
		},
	}

	var vig = &cobra.Command{
		Use:   "vig mode -c [cipher] -k [key] -i [input]",
		Short: "vignenere cipher",
		Long:  `perform vigenere cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("missing input, see -h (--help) for more info")
		},
	}

	var random = &cobra.Command{
		Use:   "random",
		Short: "randomly transpose",
		Long:  `randomly transposes input`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide an input string")
				os.Exit(1)
			}
			for _, arg := range args {
				fmt.Println(psyfer.TransposeRandom(arg))
			}
		},
	}

	var railfence = &cobra.Command{
		Use:   "railfence",
		Short: "railfence transpose",
		Long:  `performs railfence transposition on input`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide an input string")
				os.Exit(1)
			}
			for _, arg := range args {
				if decrypt {
					fmt.Println(psyfer.DeTransposeRailFence(arg))
				} else {
					fmt.Println(psyfer.TransposeRailFence(arg))
				}
			}
		},
	}
	var split = &cobra.Command{
		Use:   "split",
		Short: "split transpose",
		Long:  `performs split transposition on input`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide an input string")
				os.Exit(1)
			}
			for _, arg := range args {
				if decrypt {
					fmt.Println(psyfer.DeTransposeSplit(arg))
				} else {
					fmt.Println(psyfer.TransposeSplit(arg))
				}
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

	split.Flags().BoolVarP(
		&decrypt,
		"decrypt",
		"d",
		false,
		"decrypt",
	)
	railfence.Flags().BoolVarP(
		&decrypt,
		"decrypt",
		"d",
		false,
		"decrypt",
	)

	sub.Flags().StringVarP(
		&input,
		"input",
		"i",
		"",
		"string to be encrypted",
	)

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(sub, aes, trans, vig)
	sub.AddCommand(crack)
	aes.AddCommand(crack)
	trans.AddCommand(random, railfence, split)
	vig.AddCommand(crack)
	rootCmd.Execute()
}
