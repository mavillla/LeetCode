func lengthOfLongestSubstring(s string) int {
    start, max := 0, 0
    if len(s) != 0 {
        charMap := make(map[rune]int)
        for i, c := range s {
            if v, ok := charMap[c]; ok && start <= v {
                start = v + 1
            }
            charMap[c] = i
            if max < i - start + 1 {
                max = i - start + 1
            }
        }
    }
    return max
}
