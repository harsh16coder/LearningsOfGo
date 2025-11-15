package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	minVal := val
	if len(this.minStack) != 0 {
		if this.minStack[len(this.minStack)-1] < val {
			minVal = this.minStack[len(this.minStack)-1]
		}
	}
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	minStack := MinStack{}
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(0)
	fmt.Println(minStack.GetMin()) // return 0
	minStack.Pop()
	fmt.Println(minStack.Top()) //) return 2
	fmt.Println(minStack.GetMin())
}

/*
Design a stack class that supports the push, pop, top, and getMin operations.

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
Each function should run in
O
(
1
)
O(1) time.
*/
