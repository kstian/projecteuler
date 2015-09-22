package problem

//Find the sum of all the multiples of 3 or 5 below limit.
func Problem1(limit int) int {
	const THREE = 3
	const FIVE = 5
	result := 0
	for i := 0; i < limit; i++ {
		if i % THREE == 0 || i % FIVE == 0{
			result += i
		}
	}
	return result
}