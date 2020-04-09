package main

import "fmt"

func main() {
	//Array must be sorted for binary search
	var randomInt = []int{1, 2, 3, 4, 5, 6, 7, 23}
	search := binarySearch(randomInt, 23)
	fmt.Println("Search found: ", search)
}

func binarySearch(arr []int, num int) int {
	var found = 0
	low := 0
	high := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		mid := (low + high) / 2
		if arr[mid] == num {
			return arr[mid]
		} else if arr[mid] > num {
			high = mid - 1
		} else if arr[mid] < num {
			low = mid + 1
		}
	}
	return found
}
