func solution(i int, j int, k int) int {
    count := 0

    for n := i; n <= j; n++ {
        num := n
        for num > 0 {
            digit := num % 10      
            if digit == k {
                count++
            }
            num /= 10              
        }
    }

    return count
}
