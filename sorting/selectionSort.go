package main

import "fmt"

func main() {
	var randomInt = []int{10, 9, 8, 22, 11, 23, 9}

	fmt.Println("Unsorted : ")
	fmt.Println(randomInt)

	fmt.Println("Selection Sorted : ")

	selectionSort(randomInt)
	fmt.Println(randomInt)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := findMinIndex(arr, i)
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func findMinIndex(arr []int, index int) int {
	min := index
	for i := index; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return min
}
