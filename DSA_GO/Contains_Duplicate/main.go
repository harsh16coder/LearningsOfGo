package main

import "fmt"

func hasDuplicate(nums []int) bool {
	hashmap := make(map[int]bool)
	for i := range nums {
		_, ok := hashmap[nums[i]]
		if ok == true {
			return true
		}
		hashmap[nums[i]] = true
	}
	return false
}

func main() {
	var nums = []int{1, 2, 3, 3, 4, 1}
	fmt.Println(hasDuplicate(nums))
}

// Problem

// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

// Example 1:

// Input: nums = [1, 2, 3, 3]

// Output: true

// Example 2:

// Input: nums = [1, 2, 3, 4]

// Output: false
