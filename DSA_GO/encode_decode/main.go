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
	encode := sol.Encode([]string{"neet", "heat", "sir"})
	fmt.Println(encode)
	fmt.Println(sol.Decode(encode))
}

/*
🧠 Encode and Decode Strings (LeetCode #271) — Revision Notes
1️⃣ Goal
Design algorithms to convert a list of strings → one string (encode) → back to the same list (decode)
➡ Must handle empty strings and special characters like # safely.

2️⃣ Core Idea
Use a length-prefix encoding:
Each string is stored as
[length]#[string]

Example: ["neet","heat","sir"] → "4#neet4#heat3#sir"

3️⃣ Encoding Logic


Iterate through each string.


Append its length (strconv.Itoa(len(str))) + "#" + actual string to the result.


Join everything into one long string.
✅ Works even if the string contains # or is empty ("0#").



4️⃣ Decoding Logic


Traverse the encoded string with two pointers i and j.


Move j until you hit '#' → extract substring s[i:j] = length.


Convert to integer → extract next length characters as actual word.


Append to result and move i ahead.



5️⃣ Complexity


Time: O(N) → Linear in total characters.


Space: O(N) → Result and temporary storage.
(N = total characters across all strings.)



6️⃣ Why It’s Robust
✅ Handles empty strings ([""] → "0#").
✅ Allows # inside words (because decoding uses length, not delimiters).
✅ No ambiguity in splitting — fully deterministic.

🧩 Quick Summary:

Encode → “store length + # + word”
Decode → “read length → slice next word → repeat”


Would you like a short one-liner mnemonic (like “Count–Mark–Slice–Repeat”) to memorize this logic faster?
*/
