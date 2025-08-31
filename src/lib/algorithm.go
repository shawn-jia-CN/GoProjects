package lib

import (
	"fmt"
)

func BinarySearch(input []int, target int) (int, bool) {
	var left int = 0
	var right int = len(input) - 1
	var mid int = 0

	if len(input) == 0 {
		fmt.Println("Input length is 0")
		return 0, false
	}

	for left != right {
		mid = (left + right) / 2
		if input[mid] == target {
			return mid, true
		} else if input[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	if input[left] == target {
		return left, true
	}
	return 0, false

}

func SelectSort(input []int) bool {
	if len(input) == 0 {
		fmt.Println("Input is nil")
		return false
	}

	var tail = len(input) - 1
	for tail != 0 {
		index := 1
		max := 0
		for index <= tail {
			if input[index] > input[max] {
				max = index
			}
			index += 1
		}
		tmp := input[max]
		input[max] = input[tail]
		input[tail] = tmp
		tail -= 1
	}
	return true
}

func InsertSort(input []int) bool {
	if len(input) == 0 {
		fmt.Println("Input is nil")
		return false
	}

	var tail int = 1
	for tail <= len(input)-1 {
		index := 0
		val := input[tail]
		last := tail - 1
		insertIndex := tail
		for index <= last {
			if input[index] > input[tail] {
				insertIndex = index
				break
			}
			index = index + 1
		}
		for last >= insertIndex {
			input[last+1] = input[last]
			last -= 1
		}
		input[insertIndex] = val
		tail = tail + 1
	}
	return true
}
