func solution(array []int, n int) int {
    
    number := 0
    
    for _, v := range array {
        if v == n {
        number++
        }
    }
    return number
}