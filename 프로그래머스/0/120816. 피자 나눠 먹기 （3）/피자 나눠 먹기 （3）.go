func solution(slice int, n int) int {
     if n%slice == 0 {
        return n / slice
    }
    return n/slice + 1
}