package problem

// Find n-th fibonacci's number
func fib(n int) int {
	if n ==0 {
		return 0
	}
	if n ==1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}

// Find the sum of the even-valued fibonacci's number where the sum is below limit
func Problem2(limit int) int {
	n := 0
	result := 0
	sum := 0
	for result <= limit{
		if result % 2 == 0 {
			sum += result
		}
		n++
		result = fib(n)
	}
	return sum
}