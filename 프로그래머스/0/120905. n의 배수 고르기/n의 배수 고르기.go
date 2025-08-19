func solution(n int, numlist []int) []int {
    result := []int{} 

    for _, v := range numlist {
        if v % n == 0 {       
            result = append(result, v)
        }
    }

    return result
}
