package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	if k == 1 && len(nums) == 1 {
		return nums
	}
	targetarray := []int{}
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	frequencyArray := make([][]int, len(nums)+1)
	for num, cnt := range countMap {
		frequencyArray[cnt] = append(frequencyArray[cnt], num)
	}
	for i := len(frequencyArray) - 1; i > 0; i-- {
		targetarray = append(targetarray, frequencyArray[i]...)
		if len(targetarray) == k {
			return targetarray
		}
	}
	return targetarray
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3}
	target := 2
	fmt.Println(topKFrequent(arr, target))
}

/*
🧠 Top K Frequent Elements – Flash Notes (Go)
🎯 Goal

Find the k most frequent elements from an integer array.

⚙️ Core Idea

Use Bucket Sort — frequency as an index, collect from highest bucket down.

🪣 Algorithm Flashcards

1️⃣ Count Frequencies
→ Use map[int]int to count each number.

countMap[num]++


2️⃣ Create Buckets
→ Use slice of slices, size = len(nums)+1.
→ Each index i stores numbers with frequency i.

frequencyArray[cnt] = append(frequencyArray[cnt], num)


3️⃣ Collect Top K
→ Traverse from highest frequency → lowest.
→ Append until you have k elements.

for i := len(frequencyArray)-1; i > 0; i-- { ... }


4️⃣ Return Result
→ Once len(result) == k, return.

⏱ Complexity

Time: O(n)

Space: O(n)

⚡ Concept Keywords

bucket sort • hashmap counting • top-k pattern • O(n) selection

🧩 Edge Case

If single element and k=1 → return directly

if k == 1 && len(nums) == 1 { return nums }

💡 Intuition

Instead of sorting frequencies →
Place numbers in buckets by their frequency count,
then read back from highest to lowest.

🧠 Compare

Heap approach: O(n log k)

Bucket approach: O(n) (faster, simpler)

🔁 When to Use

Use bucket sort when:

You need Top K Frequent elements

You can afford O(n) space

The range of possible frequencies ≤ n

🚀 Key Takeaway

Group by frequency → traverse high to low → grab K → done.
*/
