func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, maxLength := 0, 1
	for i := 0; i < len(s)-maxLength/2; i++ {
		newStart, newLength := extendPalindrome(s, i, i)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
		newStart, newLength = extendPalindrome(s, i, i+1)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
	}
	return s[start : start+maxLength]
}

func extendPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - i - 1
}
