package main

func partition(low int, high int, less func(i int, j int) bool, swap func(i int, j int)) int {
	i := low
	for j := low; j < high; j++ {
		if less(j, high) {
			swap(i, j)
			i++
		}
	}
	swap(high, i)
	return i
}

func quickSortRec(low int, high int, less func(i int, j int) bool, swap func(i int, j int)) {
	if low < high {
		q := partition(low, high, less, swap)
		if low < q-1 {
			quickSortRec(low, q-1, less, swap)
		}
		if high > q+1 {
			quickSortRec(q+1, high, less, swap)
		}
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quickSortRec(0, n-1, less, swap)
}

func main() {

}
