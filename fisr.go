package main

import (
	"fmt"
	"math"
	"time"
)

func fastInverseSquareRoot(x float64) float64 {
	var x2, y float64
	const threehalfs float64 = 1.5
	x2 = x * 0.5
	y = math.Float64frombits(0x5fe6eb50c7b537a9 - (math.Float64bits(x) >> 1))
	y = y * (threehalfs - (x2 * y * y))
	return y
}

func normalInverseSquareRoot(x float64) float64 {
	return 1 / math.Sqrt(x)
}

func main() {
	var normalInverseSquareRootResult, fastInverseSquareRootResult float64
	t0 := time.Now()
	normalInverseSquareRootResult = normalInverseSquareRoot(42)
	t1 := time.Now()
	fastInverseSquareRootResult = fastInverseSquareRoot(42)
	t2 := time.Now()
	fmt.Printf("1/sqrt(42) = %f in %d ns\n", normalInverseSquareRootResult, t1.Sub(t0).Nanoseconds())
	fmt.Printf("fi sqrt(42) = %f in %d ns\n", fastInverseSquareRootResult, t2.Sub(t1).Nanoseconds())
}
