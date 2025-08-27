func solution(start_num int, end_num int) []int {
    length := start_num - end_num + 1
    answer := make([]int, length)
    
    for i := 0; i< length; i++ {
        answer[i] = start_num - i
    }
    return answer
}