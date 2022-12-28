package main

import "fmt"

func Sort(arr []int, low, high int) []int {
	if low < high {
		var i int
		arr, i = SortP(arr, low, high)
		arr = Sort(arr, low, i-1)
		arr = Sort(arr, i+1, high)
	}
	// fmt.Println(arr)
	return arr
}

func SortP(arr []int, low, high int) ([]int, int) {
	pi := arr[high]
	i := low
	// fmt.Println(pi)
	// fmt.Println(i)
	for j := low; j < high; j++ {
		if arr[j] < pi {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	// fmt.Println(arr)
	return arr, i
}

func main() {
	array := []int{5, 6, 7, 2, 1, 0}
	result := Sort(array, 0, len(array)-1)
	fmt.Println(result)
}
