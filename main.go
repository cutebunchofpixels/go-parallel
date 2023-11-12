package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	sliceSize := 10000000
	measurementIterations := int64(10)
	totalSerialTime, totalParallelTime := int64(0), int64(0)

	itemsSerial := generateRandomSlice(sliceSize)
	itemsParallel := generateRandomSlice(sliceSize)

	for i := int64(0); i < measurementIterations; i++ {
		start := time.Now()
		serialMergeSort(itemsSerial)
		totalSerialTime += time.Since(start).Milliseconds()

		start = time.Now()
		parallelMergeSort(itemsParallel)
		totalParallelTime += time.Since(start).Milliseconds()

		if !isSliceSorted(itemsSerial) || !isSliceSorted(itemsParallel) {
			fmt.Println("Error during sorting arrays")
			os.Exit(-1)
		}
	}

	fmt.Printf("Average serial execution time: %d ms\n", totalSerialTime/measurementIterations)
	fmt.Printf("Average parallel execution time: %d ms\n", totalParallelTime/measurementIterations)
}

func isSliceSorted(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}

	return true
}
