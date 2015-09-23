package problem
import (
	"github.com/ktarou/projecteuler/utils"
	"strconv"
)

func Problem4() int {
	var palindrome []int
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			if utils.IsPalindrome(strconv.Itoa(i * j)) {
				palindrome = append(palindrome, i * j)
			}
		}
	}
	return utils.Max(palindrome)
}
