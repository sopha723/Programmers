func solution(start_num int, end_num int) []int {
    res := []int{}
    
    for i := start_num; i>=end_num; i-- {
        res = append(res, i)
    }
    
    return res
}