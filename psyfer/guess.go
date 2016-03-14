package psyfer

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var charFreq = []float64{
	0.080,
	0.015,
	0.030,
	0.040,
	0.130,
	0.020,
	0.015,
	0.060,
	0.065,
	0.005,
	0.005,
	0.035,
	0.030,
	0.070,
	0.080,
	0.020,
	0.002,
	0.065,
	0.060,
	0.090,
	0.030,
	0.010,
	0.015,
	0.005,
	0.020,
	0.002,
}

func pc(position, i int) int {
	if position-i < 0 {
		return 26 - (int(math.Abs(float64((position - i) % 26))))
	} else {
		return int(math.Abs(float64((position - i) % 26)))
	}
}

func (p phi) String() string {
	return fmt.Sprintf("%d %f\n", p.position, p.value)
}

type phi struct {
	position int
	value    float64
}

type allPhi []phi

func (p allPhi) Len() int           { return len(p) }
func (p allPhi) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p allPhi) Less(i, j int) bool { return p[i].value < p[j].value }

func CharacterFrequency(input string) map[rune]float64 {
	cf := make(map[rune]int)
	freq := make(map[rune]float64)
	for _, c := range input {
		cf[c]++
	}
	l := len(input)
	for c, n := range cf {
		freq[c] = float64(n) / float64(l)
	}
	return freq
}

func Guess(input string) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	input = strings.ToUpper(strings.Replace(input, " ", "", -1))
	freq := CharacterFrequency(input)
	total := []phi{}
	for i := 0; i < 26; i++ {
		a := 0.0
		for c, n := range freq {
			index := int(c - 'A')
			a += n * charFreq[pc(index, i)]
		}
		p := phi{i, a}
		total = append(total, p)
	}
	sort.Sort(allPhi(total))
	for i := len(total) - 1; i > len(total)-6; i-- {
		fmt.Println(VigenereCipher(input, string(alphabet[total[i].position]), true))
	}
}
