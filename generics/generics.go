package generics

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//1st method
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//2nd method
// func printSlice[T string | int | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//3rd method

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Top() T {
	return s.values[len(s.values)-1]
}
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {

	}
	index := len(s.values) - 1
	poppedEle := s.values[index]
	s.values = s.values[:index]
	return poppedEle, true
}
