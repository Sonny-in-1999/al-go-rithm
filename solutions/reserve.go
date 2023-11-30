package solutions

import "strings"

func Reserve(word string) string {
	//var res string
	//for i := len(word) - 1; i >= 0; i-- {
	//	res += string(word[i])
	//}
	//return res
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
