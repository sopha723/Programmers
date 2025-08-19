func solution(numbers []int, direction string) []int {

	if direction == "right" {
		return append([]int{numbers[len(numbers)-1]}, numbers[:len(numbers)-1]...)
		
	} else {
return append(numbers[1:], numbers[0])
	}
	return []int{}
}