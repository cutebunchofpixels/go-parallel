package main

import (
	"math/rand"
	"sort"
	"sync"
)

const MinSubarraySize = 5000

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) - 500
	}
	return slice
}

func merge(source, left, right []int) {
	i, k, j := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			source[k] = left[i]
			i++
		} else {
			source[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		source[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		source[k] = right[j]
		j++
		k++
	}
}

func serialMergeSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	middle := len(slice) / 2
	left := make([]int, middle)
	right := make([]int, len(slice)-middle)
	copy(left, slice[:middle])
	copy(right, slice[middle:])

	serialMergeSort(left)
	serialMergeSort(right)

	merge(slice, left, right)
}

func parallelMergeSort(slice []int) {
	if len(slice) <= MinSubarraySize {
		sort.Slice(slice, func(i int, j int) bool { return slice[i] <= slice[j] })
		return
	}

	middle := len(slice) / 2
	left := make([]int, middle)
	right := make([]int, len(slice)-middle)
	copy(left, slice[:middle])
	copy(right, slice[middle:])

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		parallelMergeSort(left)
		wg.Done()
	}()

	go func() {
		parallelMergeSort(right)
		wg.Done()
	}()

	wg.Wait()

	merge(slice, left, right)
}
