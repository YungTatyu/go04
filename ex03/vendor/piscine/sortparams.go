package piscine

import (
	"ft"
	"os"
)

func QuickSort(arr []string, leftmostIndex, righmostIndex int) {
	if leftmostIndex < righmostIndex {
		pivotIndex := partition(arr, leftmostIndex, righmostIndex)
		// partitionで分けられた左右のpartをそれぞれsortする
		QuickSort(arr, leftmostIndex, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, righmostIndex)
	}
}

func partition(arr []string, leftmostIndex, rightmostIndex int) int {
	pivot := arr[rightmostIndex] // ピボットを右端とする
	// pivotより値が大きい要素のindex
	firstGreterIndex := leftmostIndex - 1

	for i := leftmostIndex; i < rightmostIndex; i++ {
		if arr[i] <= pivot {
			firstGreterIndex++
			arr[firstGreterIndex], arr[i] = arr[i], arr[firstGreterIndex]
		}
	}
	firstGreterIndex++ // pivotと一番左にあるpivotより大きい値をswap
	arr[firstGreterIndex], arr[rightmostIndex] = arr[rightmostIndex], arr[firstGreterIndex]
	return firstGreterIndex
}

func Print(s string) {
	for _, v := range s {
		ft.PrintRune(v)
	}
}

func length(arr []string) int {
	var i int = 0
	for range arr {
		i++
	}
	return i
}

func SortParams() {
	var argv []string = os.Args[1:]
	len := length(argv)
	if len == 0 {
		return
	}
	QuickSort(argv, 0, length(argv)-1)
	for _, v := range argv {
		Print(v + "\n")
	}
}
