package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var r *rand.Rand

func x_squared(x float64) float64 {
	return x * x
}

func point_is_below(x float64, y float64) bool {
	return y < x_squared(x)
}

func random_point(x_min float64, width float64, height float64) (float64, float64) {
	x := x_min + rand.Float64()*width
	y := r.Float64() * height
	return x, y
}

func main() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	x1 := 10.0
	x2 := 20.0
	width := x2 - x1
	y2 := x_squared(x2)
	fmt.Print("(x1, x2) = (%f, %f), y2 = %f\n", x1, x2, y2)
	height := y2
	trials := 100
	if len(os.Args) > 1 {
		trialsTmp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("input is not a number: %v, defaulting to 100\n", os.Args[1])
			trials = 100
		} else {
			trials = trialsTmp
		}
	}

	count_below := 0
	for i := 0; i < trials; i++ {
		x, y := random_point(x1, width, height)
		if point_is_below(x, y) {
			count_below = count_below + 1
		}
	}
	//fmt.Print("count_below %v, trials = %v\n" count_below, trials)
	percent_below := float64(count_below) / float64(trials)

	fmt.Printf("percent_below = %\n", percent_below)
	integral_estimate := percent_below * (x2 - x1) * y2
	fmt.Printf("integral estimate is %f\n", integral_estimate)
	// fmt.Printf("f(%f) = %f\n, x, x_squared(x))
	// fmt.Printf("is the point below the line? %v", point_is_below(x, y))

}
