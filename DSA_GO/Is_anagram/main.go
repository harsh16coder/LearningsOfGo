package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "jar"
	t := "jam"
	fmt.Println(isAnagram(s, t))
}

// bruteforce
//  if len(s) != len(t) {
//     return false
//  }
//  hashmap := make(map[string]int)
//  for i := range(len(s)){
//     hashmap[string(s[i])]++
//  }
//  for i := range(len(s)){
//     hashmap[string(t[i])]--
//  }
//  for i:= range(len(s)) {
//     freq, _ := hashmap[string(s[i])]
//     if freq==0{
//         continue
//     } else {
//         return false
//     }
//  }
//  return true
