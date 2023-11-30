package solutions

import "strings"

func Reserve(word string) string {
	/* // Solution 1.
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}
	return res
	*/

	/* // Solution 2
	var res string
	for _, w := range word {
		res = string(w) + res
	}
	return res
	*/

	// solution 3
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
