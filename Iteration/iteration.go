package iteration

import "strings"

const repeatation = 5

func iteration(char string) string {
	var appendstring string
	for i := 0; i < repeatation; i++ {
		appendstring += char
	}
	return appendstring
}

func iterationloop(char string) string {
	var appendstring strings.Builder
	for i := 0; i < repeatation; i++ {
		appendstring.WriteString(char)
	}
	return appendstring.String()
}
