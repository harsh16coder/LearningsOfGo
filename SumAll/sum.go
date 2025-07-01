package sumall

func SumAll(numbersToSum ...[]int) []int {
	var resultSlice []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			resultSlice = append(resultSlice, 0)
		} else {
			resultSlice = append(resultSlice, Sum(number))
		}
	}
	return resultSlice
}

func Sum(num []int) int {
	var result int
	for _, val := range num {
		result += val
	}
	return result
}
