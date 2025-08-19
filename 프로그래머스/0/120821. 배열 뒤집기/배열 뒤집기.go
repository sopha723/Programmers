func solution(num_list []int) []int {
    n := len(num_list)
    result := make([]int, n)

    for i := 0; i < n; i++ {
        result[i] = num_list[n-1-i]
    }

    return result
}
