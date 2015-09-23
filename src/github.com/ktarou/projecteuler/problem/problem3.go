package problem

func Problem3(num int) []int {
	var factors []int
	for i := 2; i <= num; i++ {
		for num % i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}
	return factors
}