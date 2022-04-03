package main

import (
	"math"
)

func main() {
}

//1
func Problem1(n int) float64 {

	res := 1.0
	for i := 1; i <= n; i++ {
		res = res + float64(1)/(float64(Factorial(i)))
	}
	return res
}

//2
func Problem2(x float64, n int) float64 {

	res := 1.0

	for i := 1; i <= n; i++ {
		res = res + math.Pow(x, float64(i))/float64(Factorial(i))
	}
	return res
}

func Problem3(x float64) float64 {
	return math.Sin(x)
}
func Problem4(x float64) float64 {
	return math.Cos(x)
}
func Problem5(x float64) float64 {
	return math.Sin(x)
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Problem11(index int) int {
	if index == 0 {
		return 1
	} else {
		return (2 + 1) / Problem11(index-1)
	}
}

func Problem12(index int) int {
	if index == 0 {
		return 1
	} else {
		return ((Problem12(index-1) - 1) / index)
	}
}

func Problem13(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Problem13(index-1) + Problem13(index-2)
	}
}

func Problem16(stdbase, stdpower int) int {
	count := 0
	for i := float64(1); i <= float64(stdbase); i++ {
		count = count + int(math.Pow(i, float64(stdpower)))
	}
	return count
}

func Problem17(stdin int) int {
	count := 0
	for i := float64(1); i <= float64(stdin); i++ {
		count = count + int(math.Pow(i, i))
	}
	return count
}
