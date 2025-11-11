package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	productPrefix := 1
	productSuffix := 1
	prefix[0] = 1
	for i := 1; i <= len(nums); i++ {
		prefix[i-1] = productPrefix
		productPrefix *= nums[i-1]
	}
	suffix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= -1; i-- {
		suffix[i+1] = productSuffix
		productSuffix *= nums[i+1]
		fmt.Println(productSuffix)
		fmt.Println(suffix)
	}
	for i := 0; i < len(nums); i++ {
		result[i] = prefix[i] * suffix[i]
	}
	return result
}

func main() {
	nums := []int{1, 2, 4, 6}
	fmt.Println(productExceptSelf(nums))
}
