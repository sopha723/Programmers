func solution(n int) int {
    fac := 1
    i := 0

        for{
            if fac <= n {
                fac *= i + 1
                i++
            }else {
                break
        }
        }
        return i-1
}