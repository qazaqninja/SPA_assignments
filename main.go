package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	num := 0
	Permut(a, 0, &num)
	fmt.Println(100 * (float64(num) / float64(Factorial(len(a)))))
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func Permut(a []int, cid int, num *int) {
	count := 1
	if cid == len(a) {

		for i := 0; i < len(a)-1; i++ {
			count *= a[i] - a[i+1]
		}
		if float64(count) < math.Sqrt(float64(Factorial(len(a)))) {
			*num++
		}

		return
	}
	for i := cid; i < len(a); i++ {
		Swap(a, i, cid)
		Permut(a, cid+1, num)
		Swap(a, i, cid)
	}
}
