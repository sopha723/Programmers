func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func solution(numer1 int, denom1 int, numer2 int, denom2 int) []int {
    numer := numer1*denom2 + numer2*denom1
    denom := denom1 * denom2

    g := gcd(numer, denom) 
    numer /= g
    denom /= g

    return []int{numer, denom}
}
