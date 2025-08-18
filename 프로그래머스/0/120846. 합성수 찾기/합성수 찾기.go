func solution(n int) int {
    answer := 0

    for i := 1; i <= n; i++ {
             cnt := 0
        for j := 1; j <= i; j++ {
            if i % j == 0 {
                cnt++
            }
        }
            if cnt >= 3 {
                answer++
            }
        }

        return answer
    }
