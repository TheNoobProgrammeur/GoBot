package commands

func Fibonacci(number int) int {
	if number <= 2 {
		return 1
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
}
