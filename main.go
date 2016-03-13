package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"s.mcquay.me/dm/psyfer/psyfer"

	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var key128 = psyfer.Block{
	0x2b, 0x28, 0xab, 0x09,
	0x7e, 0xae, 0xf7, 0xcf,
	0x15, 0xd2, 0x15, 0x4f,
	0x16, 0xa6, 0x88, 0x3c,
}
var key192 = psyfer.Block{
	0x00, 0x04, 0x08, 0x0c, 0x10, 0x14,
	0x01, 0x05, 0x09, 0x0d, 0x11, 0x15,
	0x02, 0x06, 0x0a, 0x0e, 0x12, 0x16,
	0x03, 0x07, 0x0b, 0x0f, 0x13, 0x17,
}
var key256 = psyfer.Block{
	0x00, 0x04, 0x08, 0x0c, 0x10, 0x14, 0x18, 0x1c,
	0x01, 0x05, 0x09, 0x0d, 0x11, 0x15, 0x19, 0x1d,
	0x02, 0x06, 0x0a, 0x0e, 0x12, 0x16, 0x1a, 0x1e,
	0x03, 0x07, 0x0b, 0x0f, 0x13, 0x17, 0x1b, 0x1f,
}

func main() {
	var key string
	var keysize int
	var decrypt bool
	var ascii bool

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
		Use:   "aes",
		Short: "aes cipher",
		Long:  `perform aes cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("missing input, see -h (--help) for more info")
				os.Exit(1)
			}
			for _, arg := range args {
				all := psyfer.BlockGen(arg)
				if ascii {
					switch keysize {
					case 128:
						fmt.Println(psyfer.ToString(all, 128, key128, decrypt))
					case 192:
						fmt.Println(psyfer.ToString(all, 192, key192, decrypt))
					case 256:
						fmt.Println(psyfer.ToString(all, 256, key256, decrypt))
					default:
						log.Fatal("keysize not valid")
					}
				} else {
					switch keysize {
					case 128:
						fmt.Println(psyfer.ToHex(all, 128, key128, decrypt))
					case 192:
						fmt.Println(psyfer.ToHex(all, 192, key192, decrypt))
					case 256:
						fmt.Println(psyfer.ToHex(all, 256, key256, decrypt))
					default:
						log.Fatal("keysize not valid")
					}
				}
			}
		},
	}

	var vig = &cobra.Command{
		Use:   "vig mode -c [cipher] -k [key] -i [input]",
		Short: "vignenere cipher",
		Long:  `perform vigenere cipher`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 || key == "" {
				fmt.Println("missing input, see -h (--help) for more info")
				os.Exit(1)
			}
			for _, arg := range args {
				arg = strings.ToUpper(strings.Replace(arg, " ", "", -1))
				key = strings.ToUpper(strings.Replace(key, " ", "", -1))
				fmt.Printf("%v\n", psyfer.VigenereCipher(arg, key, decrypt))
			}
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

	//transpose flags
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
	//sub flags
	sub.Flags().StringVarP(
		&key,
		"key",
		"k",
		"",
		"file containing key",
	)
	//aes flags
	aes.Flags().IntVarP(
		&keysize,
		"keysize",
		"k",
		0,
		"keysize",
	)
	aes.Flags().BoolVarP(
		&ascii,
		"ascii",
		"a",
		false,
		"display in ascii (may mess up terminal :) )",
	)
	aes.Flags().BoolVarP(
		&decrypt,
		"decrypt",
		"d",
		false,
		"decrypt",
	)
	//virenere flags
	vig.Flags().StringVarP(
		&key,
		"key",
		"k",
		"",
		"encryption key string",
	)
	vig.Flags().BoolVarP(
		&decrypt,
		"decrypt",
		"d",
		false,
		"decrypt",
	)

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(sub, aes, trans, vig)
	trans.AddCommand(random, railfence, split)
	rootCmd.Execute()
}
