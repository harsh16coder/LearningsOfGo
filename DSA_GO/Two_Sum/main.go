package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var targetArray []int
	hashmap := make(map[int]int)
	for i := range len(nums) {
		val, ok := hashmap[target-nums[i]]
		if ok != false {
			targetArray = append(targetArray, val)
			targetArray = append(targetArray, i)
		}
		hashmap[nums[i]] = i
	}
	return targetArray
}

func main() {
	s := []int{3, 4, 5, 6}
	t := 7
	fmt.Println(twoSum(s, t))
}
