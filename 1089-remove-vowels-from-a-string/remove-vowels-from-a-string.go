func removeVowels(s string) string {
    // Seems bad, but apparently gives us O(1) - rather than slices.Contains()
    vowels := map[rune]struct{}{
        'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
    }

   var b strings.Builder

    for _, c := range s {
        if _, isVowel := vowels[c]; !isVowel {
            b.WriteRune(c)
        }
    }

    return b.String()
}