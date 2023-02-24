package main

import (
	"fmt"
	"math/rand"
)

var (
	arr [100]int
)

func less(i, j int) bool {
	return arr[i] <= arr[j]
}
func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func partition(swap func(i, j int), less func(i, j int) bool, low int, high int) int {
	i := low
	for j := low; j < high; j++ {
		if less(j, high) {
			swap(j, i)
			i++
		}
	}
	swap(high, i)
	return i
}
func quickSortRec(swap func(i, j int), less func(i, j int) bool, low int, high int) {
	var q int
	if low < high {
		q = partition(swap, less, low, high)
		quickSortRec(swap, less, low, q-1)
		quickSortRec(swap, less, q, high)
	}

}
func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quickSortRec(swap, less, 0, n-1)
}
func main() {
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	qsort(100, less, swap)
	fmt.Println(arr)
}
