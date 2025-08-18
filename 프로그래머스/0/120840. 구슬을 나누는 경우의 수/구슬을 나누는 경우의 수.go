func solution(balls int, share int) int {
    result := 1

    for i := 1; i <= share; i++ {
        result = result * (balls - i + 1) / i
    }

    return result
}
