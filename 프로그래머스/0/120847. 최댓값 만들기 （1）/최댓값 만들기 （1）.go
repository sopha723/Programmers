func solution(numbers []int) int {
    max1, max2 := 0, 0

    for _, v := range numbers {
        if v > max1 {          
            max2 = max1       
            max1 = v           
        } else if v > max2 {  
            max2 = v           
        }
    }

    return max1 * max2
}
