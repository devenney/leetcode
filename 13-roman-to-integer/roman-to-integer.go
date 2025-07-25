// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

var values = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    total := 0
    prev := 0

    
    for i := len(s) - 1; i >= 0; i-- {
        curr := values[rune(s[i])]
        if curr < prev {
            total -= curr
        } else {
            total += curr
        }
        prev = curr
    }
    
    return total
}