package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	//bruteforce -- O(N^2)
	// length := len(numbers)
	result := []int{}
	// for i:=0;i<length;i++ {
	//     for j:=i+1;j<length;j++ {
	//         if (numbers[i]+numbers[j]==target && i!=j) {
	//             result = append(result,i+1)
	//             result = append(result,j+1)
	//             return result
	//         }
	//     }
	// }
	// return result
	// optimized one - O(N)
	// hashset := make(map[int]bool)
	// indexStore := make(map[int]int)
	// for index,val := range numbers {
	//     if (hashset[target-val]) {
	//         result = append(result,indexStore[target-val]+1)
	//         result = append(result, index+1)
	//         return result
	//     }
	//     hashset[val] = true
	//     indexStore[val] = index
	// }
	// return result
	//
	//binary search
	low := 0
	high := len(numbers) - 1
	for low < high {
		if numbers[low]+numbers[high] == target {
			result = append(result, low+1)
			result = append(result, high+1)
			return result
		}
		if numbers[low]+numbers[high] > target {
			high--
		} else if numbers[low]+numbers[high] < target {
			low++
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3))
}
