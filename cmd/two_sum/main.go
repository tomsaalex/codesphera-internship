package main

import (
	"fmt"
)

/*
Two Sum Problem

Write a function that takes a slice of integers and a target integer,
and returns the indices of the two numbers that add up to the target.

Constraints:
- Do not use the same element twice
- Exactly one solution is guaranteed
- You can return the result in any order

Example:
nums := []int{2, 7, 11, 15}
target := 9
Output: [0, 1] because nums[0] + nums[1] == 9
*/

// TwoSum finds two indices in nums whose values add up to target.
// It returns the indices as a slice of two integers.
// Assumes exactly one solution exists.
func TwoSum(nums []int, target int) []int {
	numberMap := make(map[int][]int)

	for idx, num := range nums {
		positions, exists := numberMap[num]
		if !exists {
			numberMap[num] = []int{idx}
			continue
		}
		positions = append(positions, idx)
		numberMap[num] = positions
	}

	for idx, num := range nums {
		diff := target - num

		diffList := numberMap[diff]
		for _, diffIdx := range diffList {
			if diffIdx != idx {
				return []int{idx, diffIdx}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println("two sum")
}
