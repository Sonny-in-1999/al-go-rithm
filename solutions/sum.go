package solutions

func Sum(list []int) int {
	result := 0
	for _, i := range list {
		result += i
	}
	return result
}
