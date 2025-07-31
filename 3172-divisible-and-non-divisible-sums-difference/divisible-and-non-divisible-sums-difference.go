// num1: The sum of all integers in the range [1, n] (both inclusive) that are not divisible by m.
// num2: The sum of all integers in the range [1, n] (both inclusive) that are divisible by m.

func differenceOfSums(n int, m int) int {
    // TODO: Can be O(1) time with pure arithmetic.
    
    num1 := 0
    num2 := 0

    for i := 1; i <= n; i++ {
        if i % m == 0 {
            num2 += i
        } else {
            num1 += i
        }
    }

    return num1 - num2
}