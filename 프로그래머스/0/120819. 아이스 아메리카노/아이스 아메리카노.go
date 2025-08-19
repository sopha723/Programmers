func solution(money int) []int {

	cup := money/5500
	change := money%5500
	return []int{cup, change}
}