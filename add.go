package main

import (
	"math"
	"math/rand"
)

func isOwerflowSum(first int32, second int32, base int) bool {
	if first > 0 && second > 0 && (first+second <= 0) || (first+second >= int32(base)) {
		return true
	} else {
		return false
	}
}

func add(a, b []int32, p int) []int32 {
	if len(a) < len(b) {
		a, b = b, a
	}
	base := int32(p)
	a = append(a, 0)
	for i := 0; i < len(b); i++ {
		ai := a[i]
		bi := b[i]
		if isOwerflowSum(ai, bi, p) {
			switch {
			case len(a) > i && (a[i+1] < math.MaxInt32 || a[i+1] < base):
				a[i+1]++
				a[i] = a[i] - base + b[i]
			//case len(a) > i:
			//	a[i+1] = 0
			//	a = append(a, 1)
			default:
				a = append(a, 1)
				a[i] = a[i] - base + b[i]
			}
		} else {
			a[i] += b[i]
		}
	}
	if a[len(a)-1] == 0 {
		return a[:len(a)-1]
	}
	return a
}
func main() {
	var (
		base = 10
		a    []int32
		b    []int32
	)
	for i := 0; i < 3; i++ {
		a = append(a, int32(rand.Intn(base)))
		b = append(b, int32(rand.Intn(base)))
		//a[i] = rand.Int31(100)
		//b[i] = rand.Int31(100)
	}
	println("a")
	for _, x := range a {
		print(x, " ")
	}
	print("\n")
	println("b")
	for _, x := range b {
		print(x, " ")
	}
	print("\n")
	c := add(a, b, base)
	println("c")
	for _, x := range c {
		print(x, " ")
	}
	print("\n")
}
