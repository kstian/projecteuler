package utils

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IsPalindrome(s string) bool {
	if s == Reverse(s) {
		return true
	}
	return false
}

func Max(slice []int) int {
	var max int
	for i, value := range slice {
		if i == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}
	return max
}