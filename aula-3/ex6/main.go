package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	BUBBLE_SORT    = 0
	INSERTION_SORT = 1
	SELECTION_SORT = 2
)

type BenchmarkInfo struct {
	Routine     func([]int, chan int)
	RoutineName string
}

func Swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func BubbleSort(nums []int, c chan int) {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if sorted[j] > sorted[j+1] {
				Swap(&sorted[j+1], &sorted[j])
			}
		}
	}

	c <- BUBBLE_SORT
}

func InsertionSort(nums []int, c chan int) {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)

	for i := 0; i < n; i++ {
		for j := i; j > 0 && sorted[j] < sorted[j-1]; j-- {
			Swap(&sorted[j], &sorted[j-1])
		}
	}

	c <- INSERTION_SORT
}

func SelectionSort(nums []int, c chan int) {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)

	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i + 1; j < n; j++ {
			if sorted[j] < sorted[minIndex] {
				minIndex = j
			}
		}

		Swap(&sorted[i], &sorted[minIndex])
	}

	c <- SELECTION_SORT
}

func RunBenchmarks(size int) (fastestName string, fastestTime time.Duration) {
	benchmarkInfos := map[int]BenchmarkInfo{
		BUBBLE_SORT:    {BubbleSort, "Bubble sort"},
		INSERTION_SORT: {InsertionSort, "Insertion sort"},
		SELECTION_SORT: {SelectionSort, "Selection sort"},
	}

	array := rand.Perm(size)
	c := make(chan int)

	startTimes := make(map[int]time.Time)
	for id, benchmarkInfo := range benchmarkInfos {
		startTimes[id] = time.Now()
		go benchmarkInfo.Routine(array, c)
	}

	fastestId := -1
	for range benchmarkInfos {
		completedId := <-c
		completedTime := time.Now().Sub(startTimes[completedId])

		if fastestId == -1 || completedTime < fastestTime {
			fastestId = completedId
			fastestTime = completedTime
			fastestName = benchmarkInfos[fastestId].RoutineName
		}
	}

	return
}

func main() {
	arraySizes := []int{100, 1000, 10000}

	for _, size := range arraySizes {
		routineName, routineTime := RunBenchmarks(size)
		fmt.Printf("Fastest algorithm for size %d was %s taking %s\n", size, routineName, routineTime)
	}
}
