package main

import (
	"fmt"
	"strconv"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}
	return result
}

func (s *Solution) Decode(str string) []string {
	res := []string{}
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(str[i:j])
		i = j + 1
		res = append(res, str[i:i+length])
		i += length
	}
	return res
}
func main() {
	sol := Solution{}
	encode := sol.Encode([]string{""})
	fmt.Println(encode)
	fmt.Println(sol.Decode(encode))
}
