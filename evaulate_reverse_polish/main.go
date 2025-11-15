package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		if t == "+" || t == "*" || t == "/" || t == "-" {
			switch t {
			case "+":
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				prev := stack[len(stack)-1]
				sum := curr + prev
				stack = stack[:len(stack)-1]
				stack = append(stack, sum)
			case "*":
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				prev := stack[len(stack)-1]
				multi := curr * prev
				stack = stack[:len(stack)-1]
				stack = append(stack, multi)
			case "-":
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				prev := stack[len(stack)-1]
				sub := prev - curr
				stack = stack[:len(stack)-1]
				stack = append(stack, sub)
			case "/":
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				prev := stack[len(stack)-1]
				div := prev / curr
				stack = stack[:len(stack)-1]
				stack = append(stack, div)
			}
		} else {
			num, err := strconv.Atoi(t)
			if err != nil {
				return 0
			}
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	tokens := []string{"1", "2", "+", "3", "*", "4", "-"}
	fmt.Println(evalRPN(tokens))
}
