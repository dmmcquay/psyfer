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
	var key string
	var decrypt bool

	var trans = &cobra.Command{
		Use:   "trans mode",
		Short: "transposition cipher",
		Long:  `perform transposition cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("missing input, see -h (--help) for more info")
		},
	}

	var sub = &cobra.Command{
		Use:   "sub",
		Short: "substitution cipher",
		Long:  `perform substitution cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			if key == "" || len(args) < 1 {
				fmt.Println("missing input, see -h (--help) for more info")
				os.Exit(1)
			}
			psyfer.ReadConfig(key)
			for _, arg := range args {
				fmt.Println(psyfer.Substitution(arg))
			}
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
		&key,
		"key",
		"k",
		"",
		"file containing key",
	)

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(sub, aes, trans, vig)
	aes.AddCommand(crack)
	trans.AddCommand(random, railfence, split)
	vig.AddCommand(crack)
	rootCmd.Execute()
}
