package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	res := quickSort([]int{9, 1, 2, 10, 23, 201, 5, 4, 45})
	fmt.Printf("%v", res)
}

func quickSort(arr []int) []int {

	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, left, right int) []int {
	if left < right {
		p := partition(arr, left, right)
		sort(arr, left, p-1)
		sort(arr, p+1, right)
	}
	return arr
}
func partition(arr []int, left, right int) int {
	pivot := left // 使用left作为基准
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] { //和基准比价，保证index前面的都比基准小
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
