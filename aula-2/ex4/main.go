package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Min(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("At least one number required")
	}

	min := nums[0]
	for _, val := range nums {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func Max(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("At least one number required")
	}

	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func Average(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("At least one number required")
	}

	total := 0.0
	for _, grade := range nums {
		total += grade
	}

	return total / float64(len(nums)), nil
}

func operation(op string) (func(...float64) (float64, error), error) {
	switch op {
	case minimum:
		return Min, nil
	case average:
		return Average, nil
	case maximum:
		return Max, nil
	}
	return nil, errors.New("invalid operation")
}

func main() {

	minhaFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minValue, _ := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue, _ := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)
}
