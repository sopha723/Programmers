func solution(dot []int) int {

	x, y := 1, 0

	if dot[0]*dot[1] >0 {
	y = 1
	}

	if dot[1]<0 {
	x = 2
	}
	
	return 2*x-y
}