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
    lastUnit := rune(s[len(s)-1])
    
    for i := len(s) - 1; i >= 0; i-- {
        currentRune := rune(s[i])

        fmt.Printf("lastRune: %#U, currentRune: %#U\n", lastUnit, currentRune)
        switch currentRune {
        case 'I':
            if lastUnit == 'V' || lastUnit == 'X' {
                total -= values[currentRune]
            } else {
                total += values[currentRune]
            }
        case 'X':
            if lastUnit == 'L' || lastUnit == 'C' {
                total -= values[currentRune]
            } else {
                total += values[currentRune]
            }
        case 'C':
            if lastUnit == 'D' || lastUnit == 'M' {
                total -= values[currentRune]
                fmt.Printf("subtracting %v from total (now %v)", values[currentRune], total)
            } else {
                total += values[currentRune]
            }
        default:
            total += values[currentRune]
        }

        if i>0 && (rune(s[i-1]) != currentRune) {
            lastUnit = currentRune
        }
    }
    
    return total
}