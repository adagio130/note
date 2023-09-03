package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	BuildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		MaxHeapify(arr[:i], 0)
	}
}

func BuildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		MaxHeapify(arr, i)
	}
}

func MaxHeapify(arr []int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < len(arr) && arr[l] > arr[largest] {
		largest = l
	}
	if r < len(arr) && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		MaxHeapify(arr, largest)
	}
}

