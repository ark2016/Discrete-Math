package main

import (
	"fmt"
)

func add(a, b []int32, p int) []int32 {
	lenA := len(a)
	lenB := len(b)
	radix := int32(p)
	maxLen := lenA
	if lenB > lenA {
		maxLen = lenB
	}
	result := make([]int32, maxLen+1)
	division := int32(0)
	for i := 0; i < maxLen; i++ {
		digitA := int32(0)
		digitB := int32(0)
		if i < lenA {
			digitA = a[i]
		}
		if i < lenB {
			digitB = b[i]
		}
		sum := digitA + digitB + division
		result[i] = sum % radix
		division = sum / radix
	}
	if division > 0 {
		result[maxLen] = division
	} else {
		result = result[:maxLen]
	}
	return result
}

func main() {
	a := []int32{3, 5, 7}
	b := []int32{1, 2, 4}
	p := 10
	result := add(a, b, p)
	fmt.Println(result)
	x := []int32{1, 1, 1, 1}
	y := []int32{1}
	z := 2
	result = add(x, y, z)
	fmt.Println(result)
}
