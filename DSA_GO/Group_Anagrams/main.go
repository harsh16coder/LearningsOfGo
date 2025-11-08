package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	for _, s := range strs {
		countMap := [26]int{}
		for _, c := range s {
			countMap[c-'a']++
		}
		res[countMap] = append(res[countMap], s)
		fmt.Println(countMap)
		fmt.Println(res)
	}
	resultStringArray := [][]string{}
	for _, group := range res {
		resultStringArray = append(resultStringArray, group)
	}
	return resultStringArray
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))
}

// 🧠 1. Function Purpose

// groupAnagrams(strs []string) groups all words that are anagrams of each other into separate lists.

// Example: ["act", "cat"] → same group because both have the same character counts.

// 🧩 2. Data Structure Used
// res := make(map[[26]int][]string)

// Creates a map where:

// Key: a [26]int array representing letter counts (a → index 0, b → index 1, …, z → index 25).

// Value: a slice of strings — all words matching that character count pattern (anagram group).

// 🔤 3. Iterating Through Each Word
// for _, s := range strs {

// Loops through every string in the input list strs.

// Each word will be analyzed to build its character frequency signature.

// 🧮 4. Creating a Character Frequency Signature
// countMap := [26]int{}
// for _, c := range s {
//     countMap[c-'a']++
// }

// Initializes a 26-element array (all zeros).

// For every character in the word, increments its corresponding index.

// Example for "act" →
// [1 0 1 0 0 0 0 0 0 0 0 0 1 ...] (1 ‘a’, 1 ‘c’, 1 ‘t’).

// 🪄 5. Using the Frequency Array as a Key
// res[countMap] = append(res[countMap], s)

// Uses the letter count array as the map key.

// Appends the word to the corresponding group.

// Words with identical counts (like "act" and "cat") share the same key → grouped together.

// 🧾 6. Debug Print Statements
// fmt.Println(countMap)
// fmt.Println(res)

// Prints each word’s frequency map and the growing result map — useful for understanding intermediate grouping while debugging.

// 🏗️ 7. Preparing the Final 2D Slice
// resultStringArray := [][]string{}
// for _, group := range res {
//     resultStringArray = append(resultStringArray, group)
// }

// Converts the map (unordered) into a slice of slices.

// Each inner slice is one group of anagrams.

// 📦 8. Return the Result
// return resultStringArray

// Returns the grouped anagrams as a 2D slice.

// Example output:
// [["act" "cat"] ["pots" "tops" "stop"] ["hat"]]

// 🧰 9. Main Function (Driver Code)
// func main() {
//     strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
//     fmt.Println(groupAnagrams(strs))
// }

// Defines a test input list.

// Calls groupAnagrams and prints the final grouped result.

// 🧩 10. Key Takeaway

// This approach doesn’t sort strings (which would be O(n log n) per word).

// Instead, it uses a fixed-size frequency array (26 elements) as a hash key → making it O(n * m) overall, where

// n = number of words,

// m = average word length.

// ✅ Fast, deterministic, and avoids conversions or sorting.
