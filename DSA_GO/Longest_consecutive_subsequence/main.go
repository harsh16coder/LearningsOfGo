package main

import "fmt"

func longestConsecutive(nums []int) int {
	hashset := make(map[int]struct{})
	for _, num := range nums {
		hashset[num] = struct{}{}
	}
	result := 0
	for _, num := range nums {
		if _, ok := hashset[num-1]; !ok {
			currLength := 1
			for {
				if _, ok := hashset[num+currLength]; ok {
					currLength++
				} else {
					break
				}
			}
			if result < currLength {
				result = currLength
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println(longestConsecutive(nums))
}

/*

🧠 Longest Consecutive Sequence — Revision Notes (Go)
1️⃣ Goal

Find the length of the longest consecutive elements sequence in an unsorted array.
Example:
[100, 4, 200, 1, 3, 2] → 4 (sequence is 1,2,3,4)

2️⃣ Core Idea

Use a hash set (map) for O(1) lookups and start sequences only from numbers that don’t have a previous element (num-1).
➡ This avoids redundant traversals and ensures linear time.

3️⃣ Algorithm Steps

Store all numbers in a map[int]bool for constant-time lookup.

Iterate over each number:

If (num-1) exists → skip (not start of a new sequence).

If (num-1) doesn’t exist → it's the start of a new sequence.

From that start, keep checking for (num+1), (num+2), etc., until missing.

Track the maximum sequence length.

4️⃣ Example Walkthrough

Input: [100, 4, 200, 1, 3, 2]

num=100 → start (no 99) → len=1

num=4 → skip (3 exists)

num=200 → start (no 199) → len=1

num=1 → start (no 0) → check 2,3,4 → len=4
✅ Max = 4 → Answer = 4

5️⃣ Complexity
Metric	Value	Reason
Time	O(n)	Each number checked once; forward scan only from sequence starts
Space	O(n)	Hash set to store numbers
6️⃣ Key Insights

Using a hash set gives O(1) membership checks.

Start only when (num-1) missing — avoids re-counting subsequences.

Works for negative numbers and duplicates.

Linear-time solution — no sorting needed.
*/
