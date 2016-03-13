package psyfer

import (
	"fmt"
	"math"
	"sort"
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

func Guess() {
	fmt.Println((6 - 7) % 26)
	total := []phi{}
	for i := 0; i < 26; i++ {
		a := 0.0
		a += 0.1 * charFreq[pc(6, i)]
		a += 0.1 * charFreq[pc(7, i)]
		a += 0.1 * charFreq[pc(10, i)]
		a += 0.3 * charFreq[pc(14, i)]
		a += 0.2 * charFreq[pc(17, i)]
		a += 0.1 * charFreq[pc(20, i)]
		a += 0.1 * charFreq[pc(25, i)]
		p := phi{i, a}
		total = append(total, p)
	}
	fmt.Println(total)
	sort.Sort(allPhi(total))
	fmt.Println(total)
}
